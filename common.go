// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package goutils

import "github.com/astaxie/beego/logs"

func init() {

}

var Logger *logs.BeeLogger

func InitLogs() {
	Logger = logs.NewLogger()
	Logger.SetLogger(logs.AdapterFile, `{"filename":"common.log","maxsize":33554432}`) //32M
	Logger.Debug("this is a debug message")
}

//要讓程式停下來才用的
func CheckErr(err error) {
	if err != nil {
		panic(err)
		//log.Println("DB Error:", err.Error())
	}
}
