package cat

import (
    "github.com/astaxie/beego"
    "gopkg.in/redis.v3"
    "math/rand"
    "regexp"
)

const codeLength = 5
const BASE62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var power = [codeLength]int {14776336, 238328, 3844, 62, 1}
var redisClient *redis.Client

func CodeGenerator() string {
    code := ""
    for i := 0; i < codeLength; i++ {
        code += string(BASE62[rand.Intn(61)])
    }
    return code
}

func RedisConnect() {
    ip := beego.AppConfig.String("redisip")
    port := beego.AppConfig.String("redisport")
    pass := beego.AppConfig.String("redispass")
    db, _ := beego.AppConfig.Int64("redisdb")
    // Get config from conf

    client := redis.NewClient(&redis.Options{
        Addr:     ip+":"+port,
        Password: pass, // no password set
        DB:       db,  // use default DB
    })

    redisClient = client
}

func StoreURL(tag string, target string) {
    err := redisClient.Set("tag:"+tag, target, 0).Err()
    if err != nil {
        panic(err)
    }
}

func GetURL(tag string) string {
    target, err := redisClient.Get("tag:"+tag).Result()
    if err != nil {
        //panic(err)
        return "/"
    }
    return target
}

func URLValidator(url string) bool {
    re := "^(https:[/][/]|http:[/][/]|www.)[a-zA-Z0-9\\-\\.]+\\.[a-zA-Z]{2,3}(:[a-zA-Z0-9]*)?/?([a-zA-Z0-9\\-\\._\\?\\,\\'/\\\\+&amp;%\\$#\\=~])*$"
    match, _ := regexp.MatchString(re, "http://www.google.com")
    return match
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
    //RedisConnect()
    x.Data["appName"] = beego.AppConfig.String("appname")
    x.Data["appDescription"] = beego.AppConfig.String("description")
    x.Data["appSite"] = beego.AppConfig.String("appsite")
    x.TplName = "index.tpl"
}

func (s *URLShortener) Post() {
    target := s.GetString("url")
    code := ""
    if URLValidator(target) {
        RedisConnect()
        code = CodeGenerator()
        for true {
            _, err := redisClient.Get("tag:"+code).Result()
            if err == redis.Nil {
                break
            }
            code = CodeGenerator()
        }
        StoreURL(code, target)
    }
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
    RedisConnect()
    tag := w.Ctx.Input.Param(":shorten")
    target := GetURL(tag)
    w.Redirect(target, 302)
}