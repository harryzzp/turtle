package pullman

import ()
import (
	"github.com/harryzzp/turtle/utils/db"
	"log"
	"strconv"
	"fmt"
)

func PullNeeq() {
	neeq := ParseCompanyBulletin("2017-01-01", "2017-03-04")
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("Can not connect to database! %s", err.Error()) //don't use panic here with your production code base.
	}
	defer db.Close()
	content := neeq.ListInfo.Content
	for _, c := range content {
		fmt.Println(c)
		stmt, err := db.Prepare("INSERT INTO `neeq`(`CompanyCd`,`CompanyName`,`DKey`,`DestFilePath`,`DisclosureCode`,`DisclosurePostTitle`,`DisclosureSubType`,`DisclosureTitle`,`DisclosureType`,`DisclosureYear`,`FileExt`,`FilePath`,`IsEmergency`,`IsNewThree`,`PublishDate`,`PublishOrg`,`State`,`UpDate`,`Xxfcbj`) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		update := strconv.Itoa(c.UpDate.Hours) + ":" + strconv.Itoa(c.UpDate.Minutes) + ":" + strconv.Itoa(c.UpDate.Seconds)
		_, err = stmt.Exec(c.CompanyCd, c.CompanyName, c.DKey, c.DestFilePath, c.DisclosureCode, c.DisclosurePostTitle, c.DisclosureSubType, c.DisclosureTitle, c.DisclosureType, c.DisclosureYear, c.FileExt, c.FilePath, c.IsEmergency, c.IsNewThree, c.PublishDate, c.PublishOrg, c.State, update, c.Xxfcbj)
		if err != nil {
			log.Fatalf("insert data[%+v] error. %s\n", err.Error()) // handle error here instead of using panic directly

		}

	}
}

func PullCninfo() {
	cnInfo := ParseCnInfo("1", "30", "sse")
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("Can not connect to database! %s", err.Error()) //don't use panic here with your production code base.
	}
	defer db.Close()
	annoucements := cnInfo.Announcements
	for _, a := range annoucements {
		fmt.Println(a)
		stmt, err := db.Prepare("INSERT INTO `cninfo`(`secCode`, `secName`, `orgId`, `announcementId`, `announcementTitle`, `announcementTime`, `adjunctUrl`, `adjunctSize`, `adjunctType`, `storageTime`, `columnId`, `pageColumn`, `announcementType`, `associateAnnouncement`, `important`, `batchNum`, `announcementContent`, `announcementTypeName`) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		_, err = stmt.Exec(a.SecCode, a.SecName, a.OrgId, a.AnnouncementId, a.AnnouncementTitle, strconv.Itoa(a.AnnouncementTime), a.AdjunctUrl, a.AdjunctSize, a.AdjunctType, a.StorageTime, a.ColumnId, a.PageColumn, a.AnnouncementType, a.AssociateAnnouncement, a.Important, a.BatchNum, a.AnnouncementContent, a.AnnouncementTypeName)
		if err != nil {
			log.Fatalf("insert data[%+v] error. %s\n", err.Error()) // handle error here instead of using panic directly

		}

	}
}
