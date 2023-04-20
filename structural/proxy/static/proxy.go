package static

import "fmt"

// UserProxy 用户代理类
type UserProxy struct {
	User *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{User: user}
}

func (u *UserProxy) Login() error {
	if u.User.UserID=="" {
		return fmt.Errorf("userID is empty\n")
	}
	if u.User.UserName=="" {
		return fmt.Errorf("userName is empty\n")
	}
	if u.User.Password=="" {
		return fmt.Errorf("password is empty\n")
	}
	u.User.Login()
	fmt.Printf("user login success:userID:%s,userName:%s,password:%s\n",u.User.UserID,u.User.UserName,u.User.Password)
	return nil
}

