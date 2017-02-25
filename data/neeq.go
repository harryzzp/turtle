package data

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"time"
	"strconv"
	"bytes"
	"strings"
	"encoding/json"
	"github.com/harryzzp/turtle/freejson"
)

const (
	//?callback=jQuery183039972617035050884_1487818194973
	NEEQ_URL = "http://www.neeq.com.cn/disclosureInfoController/infoResult.do"
)

type update struct {
	Date           int
	Day            int
	Hours          int
	Minutes        int
	Month          int
	Seconds        int
	Time           int64
	TimezoneOffset int
	Year           int
}

type content struct {
	CompanyCd           string
	CompanyName         string
	DKey                string
	DestFilePath        string
	DisclosureCode      string
	DisclosurePostTitle string
	DisclosureSubType   string
	DisclosureTitle     string
	DisclosureType      string
	DisclosureYear      int
	FileExt             string
	FilePath            string
	IsEmergency         int
	IsNewThree          int
	PublishDate         string
	PublishOrg          string
	State               int
	UpDate              update
	Xxfcbj              string
}

type listInfo struct {
	Content          []content
	FirstPage        bool
	LastPage         bool
	Number           int
	NumberOfElements int
	Size             int
	Sort             string
	TotalElements    int
	TotalPages       int
}

type list struct {
	Code   string
	Dkey   string
	Dvalue string
	Id     int
	Name   string
}

type Neeq struct {
	ListInfo listInfo
	List     []list
}

func ParseCompanyBulletin(startTime string, endTime string) Neeq {
	var s Neeq
	str := PollCompanyBulletin(startTime, endTime)
	freejson.FreeType(string(str))
	json.Unmarshal([]byte(str), &s)
	return s
}

func PollCompanyBulletin(startTime string, endTime string) string {
	client := &http.Client{}

	var tail bytes.Buffer
	tail.WriteString(NEEQ_URL)
	tail.WriteString("?callback=jQuery")
	tail.WriteString(strconv.Itoa(time.Now().Nanosecond()))

	formData := url.Values{
		"disclosureType": {"5"},
		"page":           {"0"},
		"companyCd":      {""},
		"isNewThree":     {"1"},
		"startTime":      {startTime},
		"endTime":        {endTime},
		"keyword":        {"0"},
		"xxfcbj":         {"0"},
	}

	req, err := http.NewRequest("POST", tail.String(), strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Printf("poll error: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read neeq error!")
	}
	raw := string(body)
	return raw[strings.Index(raw, "[")+1:strings.LastIndex(raw, "]")]
}
