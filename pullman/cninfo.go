package pullman

import (
	"encoding/json"
	"net/http"
	"bytes"
	"strings"
	"fmt"
	"io/ioutil"
	"net/url"
)

const (
	CNINFO_URL = "http://www.cninfo.com.cn/cninfo-new/announcement/query"
)

type cninfo struct {
	ClassifiedAnnouncements string
	TotalSecurities         int
	TotalAnnouncement       int
	TotalRecordNum          int
	Announcements           []announcements
	CategoryList            string
	HasMore                 bool
	Totalpages              int
}

type announcements struct {
	Id                    string
	SecCode               string
	SecName               string
	OrgId                 string
	AnnouncementId        string
	AnnouncementTitle     string
	AnnouncementTime      int
	AdjunctUrl            string
	AdjunctSize           int
	AdjunctType           string
	StorageTime           string
	ColumnId              string
	PageColumn            string
	AnnouncementType      string
	AssociateAnnouncement string
	Important             string
	BatchNum              string
	AnnouncementContent   string
	AnnouncementTypeName  string
}

func ParseCnInfo(pageNum string, pageSize string, market string) cninfo {
	var ci cninfo
	body := PollAnnouncement("1", "30", "sse")
	err := json.Unmarshal(body, &ci)
	if err != nil {
		panic(err.Error())
	}
	return ci
}

func PollAnnouncement(pageNum string, pageSize string, market string) []byte {
	client := &http.Client{}

	var tail bytes.Buffer
	tail.WriteString(CNINFO_URL)

	formData := url.Values{
		"columnTitle": {"历史公告查询"},
		"pageNum":     {pageNum},
		"pageSize":    {pageSize},
		"tabName":     {"fulltext"},
		"column":      {market}, //sse：沪；szse：深
	}

	req, err := http.NewRequest("POST", tail.String(), strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Printf("poll error: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read cninfo error!")
	}
	return body
}
