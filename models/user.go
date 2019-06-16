package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	O orm.Ormer
)

func init() {

	DBInitNew().InitDatabase()
	orm.RegisterModel(new(User))
	O = orm.NewOrm()
	O.Using("default")
}

type User struct {
	Id       string `orm:"column(id);pk" description:"id"`
	Username string `orm:"column(username)" description:"username"`
	Password string `orm:"column(password)" description:"password"`
	Gender   string `orm:"column(gender)" description:"gender"`
	Age      int    `orm:"column(age)" description:"age"`
	Address  string `orm:"column(address)" description:"address"`
	Email    string `orm:"column(email)" description:"email"`
}

func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	O.Insert(&u)
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	var user User
	err = O.QueryTable("user").Filter("Id", uid).One(&user)
	if err != nil {
		return nil, errors.New("User not exists")
	}

	return &user, nil

}

func GetAllUsers() []*User {
	var users []*User
	O.QueryTable("user").All(&users)
	return users
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	user := User{Id: uid}
	if O.Read(&user) != nil {
		return nil, errors.New("User Not Exist")
	}
	user.Username = uu.Username
	user.Password = uu.Password
	user.Gender = uu.Gender
	user.Age = uu.Age
	user.Email = uu.Email
	user.Address = uu.Address

	_, err = O.Update(&user)
	if err != nil {
		return nil, errors.New("Update user failure")
	} else {
		return &user, nil
	}
}

func Login(username, password string) bool {

	users := GetAllUsers()

	for _, u := range users {
		if u.Username == username && u.Password == password {
			return true
		}
	}

	return false
}

func DeleteUser(uid string) {
	user := User{Id: uid}
	if _, err := O.Delete(&user); err == nil {
		logs.Info("Delete user [%v] success!", uid)
	} else {
		logs.Info("Delete user [%v] failure, Error: [%v]!", uid, err.Error())
	}
}
