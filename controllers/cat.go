package cat

import (
    "github.com/astaxie/beego"
    "strings"
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
    x.TplName = "index.tpl"
}

func (s *URLShortener) Post() {
    code := "aksjq"
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