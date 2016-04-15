package main
import (
    "time"
    "fmt"
)

func main() {
    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()

    fmt.Println(secs)
    fmt.Println(nanos)

    //将自Unix时间起到现在经过的秒数和纳秒数转化为时间
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}
