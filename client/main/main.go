package main

import(
	"fmt"
	"chatroom/client/controller"
)

func main(){
	var order int
	loop := true
	for loop{
		viewMainManu()
		fmt.Scanf("%d\n", &order)
		switch order{
			case 1:
				id,pwd := selection1()
				conn,err := controller.Login(id,pwd)
				if err != nil{
					continue
				}

				loop := true
				for loop{
					loop = controller.Logined(conn)
				}
			case 2:
				fmt.Println("input your user id")
			case 3:
				loop = false
				continue
			default:
				fmt.Println("please input valid order")
				continue
		}
	}
}

func viewMainManu(){
	fmt.Println("--------------------------------")
	fmt.Println("1. user login.")
	fmt.Println("2. user register.")
	fmt.Println("3. exit manu...")
	fmt.Print("please input your order: ")
}

func selection1()(id int, pwd string){
	fmt.Print("input your user id: ")
	fmt.Scanf("%d\n", &id)
	fmt.Print("input your user password: ")
	fmt.Scanf("%s\n", &pwd)
	return
}