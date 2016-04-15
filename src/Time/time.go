package main
import (
    "time"
    "fmt"
)

func main() {
    now := time.Now()
    fmt.Println(now)

    then := time.Date(2016, 4, 15, 15, 49, 0, 6587272,time.UTC)
    fmt.Println(then)

    fmt.Println(then.Weekday())

    fmt.Println(then.Before(now))


    diff := now.Sub(then)
    fmt.Println(diff)
    fmt.Println(diff.Hours())

}
