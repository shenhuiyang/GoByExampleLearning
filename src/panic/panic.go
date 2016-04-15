package main
import "os"

func main() {
    panic("a problem accured")

    _,err := os.Create("/tmp/file")
    if err != nil{
        panic(err)
    }
}
