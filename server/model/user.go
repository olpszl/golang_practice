package model

import(
	"fmt"
)

func UserLoginVerti(id int, pwd string)(isexit bool){
	isexit = true
	if id!= 1{
		isexit = false
		fmt.Println("UserId wrong")
	}
	return
}
