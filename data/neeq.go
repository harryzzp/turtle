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
)

const (
	//?callback=jQuery183039972617035050884_1487818194973
	URL = "http://www.neeq.com.cn/disclosureInfoController/infoResult.do"
)

func PollCompanyBulletin() {
	client := &http.Client{}

	var tail bytes.Buffer
	tail.WriteString(URL)
	tail.WriteString("?callback=jQuery")
	tail.WriteString(strconv.Itoa(time.Now().Nanosecond()))

	formData := url.Values{
		"disclosureType": {"5"},
		"page": {"0"},
		"companyCd": {""},
		"isNewThree": {"1"},
		"startTime": {"2017-01-24"},
		"endTime": {"2017-02-23"},
		"keyword": {"0"},
		"xxfcbj": {"0"},
	}

	req, err := http.NewRequest("POST", tail.String(), strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Printf("poll error: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}
