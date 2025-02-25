package utils

import (
	"sort"

	"macdonald-order-service/model"
)

func SortOrders(orders []*model.Order) {
	sort.Slice(orders, func(i, j int) bool {
		if orders[i].VIP != orders[j].VIP {
			return orders[i].VIP
		}
		return orders[i].ID < orders[j].ID
	})
}
