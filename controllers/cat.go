package cat

import (
    "github.com/astaxie/beego"
    "strings"
    "gopkg.in/redis.v3"
    "fmt"
)

const codeLength = 5
const BASE62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var power = [codeLength]int {14776336, 238328, 3844, 62, 1}

func IDToCode(id int) string{
    code := ""
    counter := codeLength - 1
    for ; id > 0; counter-- {
        remainder := id % 62
        code = string(BASE62[remainder]) + code
        id /= 62
    }
    for i := counter; i > -1; i-- {
        code = "a" + code
    }
    return code
}

func CodeToID(code string) int {
    id := 0
    for i := 0; i < codeLength; i++ {
        id += strings.Index(BASE62, string(code[i])) * power[i]
    }
    return id
}

func ExampleNewClient() {
    client := redis.NewClient(&redis.Options{
        Addr:     "120.27.122.121:5577",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    // Output: PONG <nil>
}

type SFO struct {    // Shovel Feces Officer
    beego.Controller
}

type URLShortener struct {
    beego.Controller
}

type URLChecker struct {
    beego.Controller
}

type URLWizard struct {
    beego.Controller
}

func (x *SFO) Get() {
    //ExampleNewClient()
    x.Data["appName"] = beego.AppConfig.String("appname")
    x.Data["appDescription"] = beego.AppConfig.String("description")
    x.Data["appSite"] = beego.AppConfig.String("appsite")
    x.TplName = "index.tpl"
}

func (s *URLShortener) Post() {
    code := "xxx"
    s.Data["json"] = &code
    s.ServeJSON()
}

func (c *URLChecker) Post() {
    //tag := c.GetString("custom")
    available := false
    c.Data["json"] = &available
    c.ServeJSON()
}


func (w *URLWizard) Get() {

}