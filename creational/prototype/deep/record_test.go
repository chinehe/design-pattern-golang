package serialize

import (
	"testing"
	"time"
)

func TestSerializePrototype(t *testing.T) {
	now := time.Now()
	userRecord:= &UserRecord{
		Name:         "userRecord",
		VisitRecords: []*VisitRecord{
			{
				Website:    "www.baidu.com",
				IP:         "12.13.14.15",
				Type:       "search",
				UserName:   "user001",
				LoginTime:  &now,
				LogoutTime: &now,
			},
			{
				Website:    "www.baidu.com",
				IP:         "12.13.14.16",
				Type:       "search",
				UserName:   "user001",
				LoginTime:  &now,
				LogoutTime: &now,
			},
			{
				Website:    "www.baidu.com",
				IP:         "12.13.14.17",
				Type:       "search",
				UserName:   "user001",
				LoginTime:  &now,
				LogoutTime: &now,
			},
			{
				Website:    "www.baidu.com",
				IP:         "12.13.14.18",
				Type:       "search",
				UserName:   "user001",
				LoginTime:  &now,
				LogoutTime: &now,
			},
		},
	}
	t.Logf("origin user record\n:%+v", userRecord)
	record := userRecord.Clone()
	t.Logf("new user record\n:%+v", record)
}