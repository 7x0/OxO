package cat

import (
    "github.com/astaxie/beego"
)

type SFO struct {    // Shovel Feces Officer
    beego.Controller
}

type URLShortener struct {
    beego.Controller
}

type URLWizard struct {
    beego.Controller
}

func (x *SFO) Get() {
    x.TplName = "index.tpl"
}

func (s *URLShortener) Post() {
    s.Data["Website"] = "beego.me"
    s.Data["Email"] = "astaxie@gmail.com"
    s.TplName = "index.tpl"
}


func (w *URLWizard) Get() {

}