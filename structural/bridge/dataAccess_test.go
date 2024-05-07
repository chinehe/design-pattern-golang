package bridge

import (
	"fmt"
	"testing"
)

func TestDataAccess(t *testing.T) {
	mysql := &MySql{}
	oracle := &Oracle{}
	access := DataAccess{}

	access.SetDatabase(mysql)
	fmt.Println(access.Query("select * from user"))

	access.SetDatabase(oracle)
	fmt.Println(access.Query("select * from user"))
}
