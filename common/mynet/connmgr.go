package mynet

import(
	"net"
	"encoding/binary"
	"fmt"
	"chatroom/common/message"
)

type ConnMgr struct{
	Conn net.Conn
}

func (this *ConnMgr) SendPkg(mes message.Message) (err error){
	//send message to server
	
	//1. serialize message.Message
	pkg := Serialize(mes)
	

	//3. send size of serialized message
	var pkglen uint32
	pkglen = uint32(len(pkg))
	buf := make([]byte, 4096)
	//3.1 writing size in buf
	binary.BigEndian.PutUint32(buf[:4], pkglen)

	//3.2 send size
	_,err = this.Conn.Write(buf)
	if err!=nil{
		fmt.Println("conn.Write(buf) size error: ", err)
		return
	}

	//4. send pkg(serialized message.Message)
	buf = pkg
	_,err = this.Conn.Write(buf)
	if err!=nil{
		fmt.Println("conn.Write(buf) data error: ", err)
		return
	}
	return
}

func (this *ConnMgr) RevPkg() (mes message.Message){
	
	//1. get pkg len
	pkglen,islogin := ReadPkgLen(this.Conn)
	if !islogin{
		fmt.Println("mynet.ReadPkgLen(conn) error")
		return
	}

	//2. get data(serialized message.Message)
	//var mes message.Message
	islogin = ReadPkg(this.Conn, pkglen, &mes)
	if !islogin{
		fmt.Println("mynet.ReadPkg(conn) error")
		return
	}
	fmt.Println("	message.Message: ", mes)

	return
}