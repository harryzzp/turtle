package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

//用json配置测试
type Config struct {
	Test1 string `json:"Test1:`
	Test2 int    `json:"Test2:`
}

var (
	config     *Config
	configLock = new(sync.RWMutex)
)

func loadConfig() bool {

	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("load config error: ", err)
		return false
	} else {
		fmt.Println(string(f))
	}

	//不同的配置规则，解析复杂度不同
	temp := new(Config)
	err = json.Unmarshal(f, &config)
	if err != nil {
		fmt.Println("Para config failed: ", err)
		return false
	}

	configLock.Lock()
	config = temp
	configLock.Unlock()
	return true
}

func GetConfig() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func init() {
	if !loadConfig() {
		os.Exit(1)
	}

	//热更新配置可能有多种触发方式，这里使用系统信号量sigusr1实现
	//s := make(chan os.Signal, 1)
	//signal.Notify(s, syscall.SIGINT)
	s := make(chan string, 1)
	go func() {
		for {
			time.Sleep(5*time.Second)
			s <- time.Now().String()
		}
	}()
	go func() {
		for {
			//<-s
			log.Printf("Reloaded config: %v, time: %s\n", loadConfig(), <-s)
		}
	}()

}

func main() {
	select {}
}