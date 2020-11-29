package message


const(
	LoginMesType = "LoginMes"
	LoginMesResType = "LoginMesRes"
)
type Message struct{
	Type string `json:"type"` //other structs type
	Data string `json:"data"` //other structs
}


//user login send
type LoginMes struct{
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
}

//server back
type LoginMesRes struct{
	Code int `json:"code"` 
	Err string `json:"err"` 
}