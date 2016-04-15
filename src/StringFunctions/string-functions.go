package main
import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Contains("test", "es"))
    fmt.Println(strings.Count("test", "t"))
    fmt.Println(strings.HasPrefix("test", "te"))
    fmt.Println(strings.HasSuffix("test","st"))
    fmt.Println(strings.Index("test", "t"))
    fmt.Println(strings.ToUpper("test"))
    fmt.Println(strings.Join([]string{"a", "b"},"-"))
    fmt.Println(strings.Split("a/b/c/s", "/"))
    fmt.Println(strings.Repeat("a", 5))
    fmt.Println(strings.Replace("fooo", "o", "x", 0))
    fmt.Println(strings.Replace("fooo", "o", "x", 1))
    fmt.Println(strings.Replace("fooo", "o", "x", 2))
}
