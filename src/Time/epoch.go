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

    //����unixʱ�������ھ�����������������ת��Ϊʱ��
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}
