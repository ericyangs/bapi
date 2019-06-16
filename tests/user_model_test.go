package test

import (
	"bapi/models"
	_ "bapi/routers"
	"path/filepath"
	"runtime"
	"testing"

	"bytes"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/logs"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {

	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)

}

// TestDB is a sample to run an Database models test
func TestDB(t *testing.T) {

	Convey("User model test0 \n", t, func() {
		Convey("All users data should be right", func() {

			users := models.GetAllUsers()

			byteobj, err := json.Marshal(users)
			if err != nil {
				errInfo := fmt.Sprintf("Mashall users error : %v! \n", err)
				panic(errInfo)
			}

			var bBuffer bytes.Buffer
			_ = json.Indent(&bBuffer, byteobj, "", "   ")
			logs.Info("\ncurrent objects list:\n  %v", bBuffer.String())
			So(len(users), ShouldBeGreaterThan, 0)
		})

	})

	Convey("User model test1 \n", t, func() {
		Convey("One user data should be right", func() {
			user, _ := models.GetUser("user_11111")
			logs.Info(user.Username)
			So(user.Username, ShouldEqual, "astaxie")
		})

	})

	var uid string

	Convey("User model test2 \n", t, func() {
		Convey("Add user data should be right", func() {
			user := models.User{
				Username: "Bobo",
				Password: "123456",
				Email:    "Bobo@oracle.com",
				Gender:   "M",
				Age:      32,
				Address:  "日本",
			}
			uid = models.AddUser(user)
			logs.Info("UID:[%v]\n", uid)

		})

	})

	Convey("User model test3 \n", t, func() {
		Convey("Update user data should be right", func() {
			user := models.User{
				Username: "Fredric",
				Password: "123456",
				Email:    "Fredric2010@outlook.com",
				Gender:   "M",
				Age:      30,
				Address:  "朝阳",
			}

			user1, _ := models.UpdateUser("user_1560671379405397800", &user)
			logs.Info("Password:[%v]\n", user1.Password)

		})

	})

	Convey("User model test4 \n", t, func() {
		Convey("Login Logic  should be right", func() {

			isLogin := models.Login("Bobo", "123456")
			logs.Info("isLogin:[%v]", isLogin)
			So(isLogin, ShouldEqual, true)

		})

	})

	Convey("User model test5 \n", t, func() {
		Convey("Delete user should be right", func() {

			models.DeleteUser(uid)

		})

	})
}
