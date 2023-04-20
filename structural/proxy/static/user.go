package static

import "fmt"

// User 用户
type User struct {
	UserID string
	UserName string
	Password string
}

func (u *User) Login() error {
	fmt.Printf("user login:userID:%s,userName:%s,password:%s\n",u.UserID,u.UserName,u.Password)
	return nil
}


