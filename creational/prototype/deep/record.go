package serialize

import (
	"fmt"
	"time"
)

// UserRecord 用户记录（原型对象）
type UserRecord struct {
	Name string
	VisitRecords []*VisitRecord
}

// VisitRecord 访问记录
type VisitRecord struct {
	Website string
	IP string
	Type string
	UserName string
	LoginTime *time.Time
	LogoutTime *time.Time
}

func (visitRecord *VisitRecord)Clone() *VisitRecord {
	return &VisitRecord{
		Website:    visitRecord.Website,
		IP:         visitRecord.IP,
		Type:       visitRecord.Type,
		UserName:   visitRecord.UserName,
		LoginTime:  visitRecord.LoginTime,
		LogoutTime: visitRecord.LogoutTime,
	}
}

func (userRecord *UserRecord)Clone() *UserRecord {
	newRecord := new(UserRecord)
	newRecord.Name = userRecord.Name
	newRecord.VisitRecords = make([]*VisitRecord,0)
	for _, visitRecord := range userRecord.VisitRecords {
		newRecord.VisitRecords = append(newRecord.VisitRecords,visitRecord.Clone())
	}
	return newRecord
}

func (userRecord *UserRecord)String() string {
	record := fmt.Sprintf("userName:%v\n",userRecord.Name)
	for _, visitRecord := range userRecord.VisitRecords {
		record+=fmt.Sprintf("%+v\n",visitRecord)
	}
	return record
}
