package mynet

import(
	"fmt"
	"encoding/json"
	"net"
	"encoding/binary"
	"bytes"
	"chatroom/common/message"
)

func Serialize(v interface{})(data []byte){
	data,err := json.Marshal(v)
	
	if err != nil{
		fmt.Println("serialize.go Serialize() error: ", err)
		return
	}

	return
}



func ReadPkgLen(conn net.Conn)(pkglen int, islogin bool){
	buf := make([]byte, 4096)
	islogin = true
	
	var pkglen32 uint32
	
	_,err := conn.Read(buf)
	if err!= nil{
		fmt.Println("controller tcontroller.go Func:Tcontroller() ", err)
		islogin = false
		return
	}

	//2.1 get size of message
	bufreader := bytes.NewReader(buf[:4])
	err = binary.Read(bufreader, binary.BigEndian, &pkglen32)
	if err != nil{
		fmt.Println("binary.Read() error: ", err)
		islogin = false
		return
	}
	fmt.Println("	len of data: ", pkglen32)

	pkglen = int(pkglen32)
	return
}

func ReadPkg(conn net.Conn, pkglen int, mes *message.Message)(islogin bool){
	islogin = true
	buf := make([]byte, 4096)
	_,err := conn.Read(buf) 
	if err!= nil{
		fmt.Println("controller tcontroller.go readPkg(): ", err)
		islogin = false
		return
	}
	//3. unserialize data to message.Message
	err = json.Unmarshal(buf[:pkglen], &mes)
	if err != nil{
		fmt.Println("json.Unmarshal() error: ", err)
		islogin = false
		return
	}
	//fmt.Println("	message.Message: ", mes)
	return
}

func MesDataUnmarshal(loginmes *message.LoginMes, mes *message.Message)(islogin bool){
	islogin = true
	buf1 := make([]byte, 4096)
	buf1 = []byte(mes.Data)

	err := json.Unmarshal(buf1, &loginmes)
	if err != nil{
		fmt.Println("Message.Data unmarshal error: ", err)
		islogin = false
	}

	return
}

