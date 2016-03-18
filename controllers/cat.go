package cat

import (
    "github.com/astaxie/beego"
    "gopkg.in/redis.v3"
    "math/rand"
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
        panic(err)
    }
    return target
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
    target := s.GetString("target")
    // TODO: Target validate
    RedisConnect()
    code := CodeGenerator()
    StoreURL(code, target)
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
    ctx.Output.Body([]byte(target))
    w.Redirect(target, 302)
}