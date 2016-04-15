package main
import (
    "fmt"
)

func main() {
    messages := make(chan string)     //同步通道
//    messages := make(chan string, 2)    //异步通道
    signals := make(chan bool)

    select {
    case msg := <-messages:
    fmt.Println("received message:", msg)
    default:
    fmt.Println("no messages received")
    }

    msg := "hi"
    select {
    case messages <- msg:
    fmt.Println("sent message", msg)
    default:
    fmt.Println("no message sent")
    }

    select {
    case msg := <-messages:
    fmt.Println("received message", msg)
    case sig := <-signals:
    fmt.Println("received signals:", sig)
    default:
    fmt.Println("no activity")
    }
}
