package serialize

import (
	"encoding/json"
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

func (userRecord *UserRecord)Clone()(*UserRecord,error)  {
	bytes, err := json.Marshal(userRecord)
	if err != nil {
		return nil, err
	}
	newRecord := new(UserRecord)
	err = json.Unmarshal(bytes, newRecord)
	return newRecord,err
}

func (userRecord *UserRecord)String() string {
	record := fmt.Sprintf("userName:%v\n",userRecord.Name)
	for _, visitRecord := range userRecord.VisitRecords {
		record+=fmt.Sprintf("%+v\n",visitRecord)
	}
	return record
}
