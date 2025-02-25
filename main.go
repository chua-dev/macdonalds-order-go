package main

import (
	"fmt"

	"macdonald-order-service/constant"
	"macdonald-order-service/model"
)

func main() {
	kitchen := model.NewKitchen()

	for {
		var input string
		promptMessage := fmt.Sprintf("\n  Enter command:\n"+"> 1. %s \n"+"> 2. %s \n"+"> 3. %s \n"+"> 4. %s \n"+"> 5. exit \n",
			constant.NewNormalOrder, constant.NewVipOrder, constant.AddBot, constant.RemoveBot)
		fmt.Println(promptMessage)
		fmt.Scanln(&input)

		switch input {
		case string(constant.NewNormalOrderId):
			kitchen.AddOrder(false)
		case string(constant.NewVipOrderId):
			kitchen.AddOrder(true)
		case string(constant.AddBotId):
			kitchen.AddBot()
		case string(constant.RemoveBotId):
			kitchen.RemoveBot()
		case string(constant.ExitId):
			fmt.Println("Exiting... \n")
			return
		default:
			fmt.Println("Command not recognized \n")
		}
	}
}
