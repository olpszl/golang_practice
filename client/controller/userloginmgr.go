package controller

import(
	"net"
	"fmt"
	"chatroom/common/message"
	"chatroom/common/mynet"
	_ "encoding/binary"

)


func Login(id int, pwd string)(conn net.Conn, err error){
	conn,err = net.Dial("tcp", "127.0.0.1:4399")
	if err != nil{
		fmt.Println("controller: userloginmgr.go Login()")
		fmt.Println("connection with server error: ", err)
		return
	}
	fmt.Println("get connection with ", conn)

	//uesrLogin, global varibal a instence of UserLogin(define in global.go)
	//userLigin.UesrId = id
	//userLogin.UersPwd = pwd

	//1. init a ins of message.LoginMes
	var loginMes message.LoginMes
	loginMes.UserId = id
	loginMes.UserPwd = pwd

	//2. serialize message.LoginMes
	data := mynet.Serialize(loginMes)

	//3. put serialized message.LoginMes in message.Message
	var mes message.Message
	mes.Type = message.LoginMesType
	mes.Data = string(data)

	//4.create a instance of mynet.ConnMgr
	var connmgr  mynet.ConnMgr
	connmgr.Conn = conn

	//5. send message
	err = connmgr.SendPkg(mes)
	if err != nil{
		fmt.Println("connmgr.Conn.SendPkg(mes) error: ", err)
		return
	}

	var mesres message.Message
	mesres = connmgr.RevPkg()
	fmt.Println("server response message: ", mesres)

	//6. judge the code
	//data := mynet. 

	return
}
