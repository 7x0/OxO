package main
import (
    "fmt"
    "time"
    "strings"
)

const codeLength = 5
const BASE62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var power = [codeLength]int {14776336, 238328, 3844, 62, 1}

func IDToBase62(id int) [codeLength]int {
    result := [codeLength]int{0}
    for counter := codeLength - 1; id > 0; counter-- {
        remainder := id % 62
        result[counter] = remainder
        id /= 62
    }
    return result
}

func Base62ToCode(result [codeLength]int) string {
    source := ""
    for _, value := range result {
        source += string(BASE62[value])
    }
    return source
}

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

func URLTransfer(id int) {

    start := time.Now().UnixNano()
    //code := Base62ToCode(IDToBase62(id))
    code := IDToCode(id)
    _id := CodeToID(code)
    stop := time.Now().UnixNano()
    runtime := stop - start
    fmt.Printf("Runtime: %dns ≈ %dms ID: %d\n", runtime, runtime / 1000, _id)
    //fmt.Printf("Origin ID: %d, Translated code: %s, Translated ID: %d\n", id,  code, _id)
}


func main() {
    URLTransfer(158123123)
}
