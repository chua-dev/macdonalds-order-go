package model

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"macdonald-order-service/constant"
)

type Kitchen struct {
	mu              sync.Mutex
	pendingOrders   []*Order
	completedOrders []*Order
	bots            []*Bot
	orderID         int
	botID           int
	NewOrder        chan bool
}

func NewKitchen() *Kitchen {
	return &Kitchen{NewOrder: make(chan bool, 1)}
}

func (k *Kitchen) AddOrder(vip bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	order := &Order{ID: k.orderID + 1, VIP: vip, ProcessingTime: constant.DefaultBotProcessingTime}
	k.orderID++

	k.pendingOrders = append(k.pendingOrders, order)
	sortOrders(k.pendingOrders)

	fmt.Printf("New Order ID %d (VIP: %v) added to PENDING AREA \n\n", order.ID, order.VIP)
	if len(k.bots) > 0 {
		go func() {
			k.NewOrder <- true
		}()
	}
	k.ShowOrders()

}

func (k *Kitchen) AddBot(faster bool) {
	k.mu.Lock()

	speed := constant.DefaultBotProcessingTime
	if faster {
		speed = constant.FasterBotProcessingTime
	}
	bot := &Bot{ID: k.botID + 1, Idle: make(chan bool, 1), Stop: make(chan bool, 1), ProcessingSpeed: speed}
	k.botID++
	k.bots = append(k.bots, bot)
	k.mu.Unlock()

	go func() {
		fmt.Printf("A new bot (ID %d) is activated.", bot.ID)
		for {
			select {
			case <-bot.Stop:
				fmt.Printf("\nStopping Bot... \n")
				k.ShowOrders()
				return
			default:
				if len(k.pendingOrders) == 0 {
					fmt.Printf("\n Bot (ID %d) is now idle \n", bot.ID)
					<-k.NewOrder
					fmt.Printf("\n Bot (ID %d) waked up \n", bot.ID)
					continue
				}
				frontOrder := k.pendingOrders[0]
				k.pendingOrders = k.pendingOrders[1:]
				fmt.Printf("\n Bot (ID %d) is now processing Order #%d (VIP: %v)\n", bot.ID, frontOrder.ID, frontOrder.VIP)
				select {
				case <-bot.Stop:
					fmt.Printf("\n Bot removed, order remained\n")
					k.pendingOrders = append([]*Order{frontOrder}, k.pendingOrders...)
					sortOrders(k.pendingOrders)
					k.ShowOrders()
					continue
				case <-time.After(time.Duration(bot.ProcessingSpeed) * time.Second):
				}
				fmt.Printf("Order #%d completed\n", frontOrder.ID)
				k.completedOrders = append(k.completedOrders, frontOrder)
				k.ShowOrders()
			}
		}
	}()
}

func (k *Kitchen) RemoveBot() {
	k.mu.Lock()
	defer k.mu.Unlock()
	if len(k.bots) > 0 {
		removedBot := k.bots[len(k.bots)-1]
		close(removedBot.Stop)
		k.bots = k.bots[:len(k.bots)-1]
		fmt.Printf("Bot (ID %d)has been removed.", removedBot.ID)
	}
}

func (k *Kitchen) ShowOrders() {
	fmt.Println("\nPENDING AREA:")
	for _, order := range k.pendingOrders {
		fmt.Printf("Order ID: %d,", order.ID)
		fmt.Printf(" VIP: %t,", order.VIP)
		fmt.Printf(" Processing Time: %d seconds\n", order.ProcessingTime)
	}

	fmt.Println("\nCOMPLETED AREA:")
	for _, order := range k.completedOrders {
		fmt.Printf("Order ID: %d", order.ID)
		fmt.Printf(" VIP: %t\n", order.VIP)
	}
}

func sortOrders(orders []*Order) {
	sort.Slice(orders, func(i, j int) bool {
		if orders[i].VIP != orders[j].VIP {
			return orders[i].VIP
		}
		return orders[i].ID < orders[j].ID
	})
}

/*
func sortOrdersLogicWithVipBool(orders []*Order, order *Order, vip bool) {
	if vip {
		index := 0
		for i, o := range orders {
			if !o.VIP {
				index = i
				break
			}
		}
		orders = append(orders[:index], append([]*Order{order}, orders[index:]...)...)
	} else {
		orders = append(orders, order)
	}
}*/
