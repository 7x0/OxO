package cat

import (
    "github.com/astaxie/beego"
    "strings"
    "gopkg.in/redis.v3"
    "fmt"
    "strconv"
)

const codeLength = 5
const BASE62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var power = [codeLength]int {14776336, 238328, 3844, 62, 1}
var redisClient *redis.Client

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

func RedisConnect() {
    ip := beego.AppConfig.String("redisip")
    port := beego.AppConfig.String("redisport")
    pass := beego.AppConfig.String("redispass")
    db, err := beego.AppConfig.Int64("redisdb")
    // Get config from conf

    client := redis.NewClient(&redis.Options{
        Addr:     ip+":"+port,
        Password: pass, // no password set
        DB:       db,  // use default DB
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    redisClient = client
    // Output: PONG <nil>
}

func AssignID() int {
    val, err := redisClient.Get("id:pointer").Result()
    if err != nil {
        panic(err)
    }
    id, err := strconv.Atoi(val)
    return id + 1
}

func StoreURL(id int, target string) {
    err := redisClient.Set("id:"+strconv.Itoa(id), "value", 0).Err()
    if err != nil {
        panic(err)
    }
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
    id := AssignID()
    code := IDToCode(id)
    StoreURL(id, target)
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