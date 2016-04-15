package main
import (
    "sort"
    "fmt"
)

func main() {
    stringList := []string{"a", "c", "b"}
    sort.Strings(stringList)
    fmt.Println(stringList)

    intArray := []int{8,10,9}
    sort.Ints(intArray)
    fmt.Println(intArray)

    s := sort.IntsAreSorted(intArray)
    fmt.Println(s)
}
