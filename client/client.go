package main

import (
    "log"
    "tcpserver/goserial"
    "io"
    "fmt"
    "time"
    "os"
)
func Write(conn io.ReadWriteCloser, bs []byte){
    bs = append(bs, 0x0D, 0x0A)
    _, err := conn.Write(bs)
    if err != nil {
        log.Fatal(err)
    }
}
func Read(s io.ReadWriteCloser, buf []byte)int{
    if n, err := s.Read(buf);err != nil{
        log.Fatal(err)
        return 0
    }else{
        fmt.Print("[",string(buf[:n]), "]")
        return n
    }
}
func Send(conn io.ReadWriteCloser, bs []byte){
    buf := make([]byte, 128)
    Write(conn, bs)
    Read(conn, buf)
    Read(conn, buf)
}
func Send2Ip(s io.ReadWriteCloser){
    time.Sleep(20 *time.Second)
    for {
        time.Sleep(1 * time.Second)
        Write(s, []byte("555555555555555555555555555555555555555555555"))
    }
}
func main() {
    fileName := "ll.log"
    logFile,err  := os.Create(fileName)
    defer logFile.Close()
    if err != nil {
        fmt.Println("open file error !")
    }
    // 创建一个日志对象
    debugLog := log.New(logFile,"[Debug]",log.LstdFlags)
    //配置一个日志格式的前缀
    debugLog.SetPrefix("[Total]")

    c := &goserial.Config{Name: "COM3", Baud: 115200}
    s, err := goserial.OpenPort(c)
    if err != nil {
        log.Fatal(err)
    }
    for i := 0 ; i< 1; i++{
        Send(s, []byte("AT"))
    }
    Send(s, []byte("AT+CIPMODE=1"))
    Send(s, []byte("AT+CIPSTART=\"TCP\",\"hualiyun.cc\",9000"))
    go Send2Ip(s)
    buf := make([]byte, 128)
    total := 0
    Read(s, buf)
    Read(s, buf)
    Read(s, buf)
    Read(s, buf)
    for{
        total += Read(s, buf)
        fmt.Println("\n", time.Now().Format("20060102150405"), " total:", total)
        log.Println(time.Now().Format("20060102150405"), " total:", total)
    }
}