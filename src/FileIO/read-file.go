package main
import (
    "io/ioutil"
    "fmt"
    "os"
    "bufio"
)

func main() {
    filePath := "E:\\test\\a.txt"
    dat, err := ioutil.ReadFile(filePath)
    check(err)
    fmt.Println(string(dat))

    f, err := os.Open(filePath)
    check(err)

    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    //从指定位置开始读取，比如从第六个字符开始读取
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    //将读取位置重新设置为从头开始，如果没有f.Seek(0, 0)操作，接下来的读取位置会默认为之前设置的从第六个字符开始
    _, err = f.Seek(0, 0)
    check(err)

    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n",string(b4))

    f.Close()
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
