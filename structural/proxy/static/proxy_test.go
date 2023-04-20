package static

import (
	"fmt"
	"testing"
)

func TestUserProxy_Login(t *testing.T) {
	proxy := &UserProxy{
		User: &User{
			UserID:   "",
			UserName: "ChineHe",
			Password: "123456",
		},
	}
	err := proxy.Login()
	if err != nil {
		fmt.Printf("error:%v\n",err)
	}
}
