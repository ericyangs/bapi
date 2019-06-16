package test

import (
	"bapi/models"
	_ "bapi/routers"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/object", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing TestGet Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			var objs map[string]*models.Object
			json.Unmarshal(w.Body.Bytes(), &objs)
			byteobj, err := json.Marshal(objs)
			if err != nil {
				errInfo := fmt.Sprintf("Mashall objs error : %v! \n", err)
				panic(errInfo)
			}

			var bBuffer bytes.Buffer
			_ = json.Indent(&bBuffer, byteobj, "", "   ")
			logs.Info("\ncurrent objects list:\n  %v", bBuffer.String())

			logs.Info("Object hjkhsbnmn123's player name:%s", objs["hjkhsbnmn123"].PlayerName)
			So(len(objs), ShouldBeGreaterThan, 0)
			// So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
