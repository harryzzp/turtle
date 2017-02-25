package main

import (
	//"github.com/harryzzp/turtle/actor"
	"github.com/harryzzp/turtle/data"
	"fmt"
)
import ()

func main() {
	//actor.TestActor()
	neeq := data.ParseCompanyBulletin("2017-01-25", "2017-02-24")
	list := neeq.List
	for _, l := range list {
		fmt.Printf("%+v\n",l)
	}
}
