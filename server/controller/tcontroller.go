package controller

import(
	"fmt"
	"net"
	"chatroom/common/message"
	"chatroom/common/mynet"
	"chatroom/server/model"
	_ "encoding/json"
)

func Tcontroller(conn net.Conn){
	defer conn.Close()
	fmt.Println("get a connection: ", conn)

	islogin := UserLogin(conn)
	if !islogin{
		fmt.Println("user login failed...")
		return
	}
	
}

func UserLogin(conn net.Conn)(islogin bool){
	var mes message.Message
	var connmgr mynet.ConnMgr

	connmgr.Conn = conn
	mes = connmgr.RevPkg()

	//3. pkg response
	if mes.Type != message.LoginMesType{
		islogin = false
		fmt.Println("tcontrooler.go UserLogin() mes.type != message.LoginMesType")
		return
	}

	//3 get loginmes
	var loginmes message.LoginMes
	islogin = mynet.MesDataUnmarshal(&loginmes, &mes)
	if !islogin{
		return
	}
	fmt.Println("		loginmes: ", loginmes.UserId, loginmes.UserPwd)

	//4.userlogin vertification
	isexit := true
	isexit = model.UserLoginVerti(loginmes.UserId, loginmes.UserPwd)
	

	//5. response
	//5.1 init a message.LoginMesRes
	var loginmesres message.LoginMesRes

	//5.2 if vertification successfully, failed
	if isexit{
		loginmesres.Code = 200
	}else{
		loginmesres.Code = 500
		loginmesres.Err = "user not exits"
	}

	//5.3 send message.LoginMesRes
	data := mynet.Serialize(loginmesres)

	var mesres message.Message
	mesres.Type = message.LoginMesResType
	mesres.Data = string(data)

	err := connmgr.SendPkg(mesres)
	if err != nil{
		fmt.Println("connmgr.SendPkg(mesres) error: ", err)
	}

	


	//6. if user not exits: exit func
	if !isexit{
		fmt.Println("user login failed, user not exit...")
		return
	}

	return 
}



