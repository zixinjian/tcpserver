package main

import (
    "net"
    "time"
    "fmt"
    "os"
)
func EchoFunc(conn net.Conn){
    fmt.Println("CONNECT FROM: ", conn.RemoteAddr())
    defer conn.Close()
    total := 0
    go SendLoop(conn)
    for {
        time.Sleep(5 * time.Second)
        buf := make([]byte, 1024)
        if i, err:= conn.Read(buf); err== nil{
            fmt.Println("R:", buf[:i])
            total += i
            fmt.Println("RTotal:", total)
        }else{
            fmt.Println("Error:", err.Error())
        }
    }
}
func SendLoop(conn net.Conn){
    total := 0
    for {
        time.Sleep(5 * time.Second)
        bs := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
        conn.Write(bs)
        total += len(bs)
        fmt.Println("STotal:", total)
    }
}
func main() {
    listener, err := net.Listen("tcp", "0.0.0.0:" + "9000")
    if err != nil {
        fmt.Println("error listening:", err.Error())
        os.Exit(1)
    }
    defer listener.Close()
    fmt.Println("TcpServer Running on :", "9000")
    for {
        conn, err := listener.Accept()
        if err != nil {
            println("Error accept:", err.Error())
            return
        }
        go EchoFunc(conn)
    }
}
