package main
import (
    "io/ioutil"
    "os"
    "fmt"
    "bufio"
)

func main() {
    filePath := "E:\\test\\a.txt"
    d1 := []byte("hello go")
    err := ioutil.WriteFile(filePath, d1, 0644)
    check(err)

    f, err := os.Create(filePath)
    check(err)
    defer f.Close()

    d2 := []byte{115, 111, 109, 101}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)

    f.Sync()

    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered string write\n")
    fmt.Println("wrote ", n4, " bytes")

    w.Flush()
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
