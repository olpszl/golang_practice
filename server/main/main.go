package main

import(
	"fmt"
    "net"
    "chatroom/server/controller"
)

func main(){
    fmt.Println("waiting connection...")
    listen,err := net.Listen("tcp", "127.0.0.1:4399")
    if err != nil{
        fmt.Println("main: main.go net.Listen() ", err)
        return
    }
    defer listen.Close()

	for{
        conn,err := listen.Accept()
        if err != nil{
            fmt.Println("main: main.go listen.accept() error: ", err)
            continue
        }else{
            go controller.Tcontroller(conn)
        }
	}
}