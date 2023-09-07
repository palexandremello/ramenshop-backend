package utils

import (
	"fmt"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
)

// FormatOrder formats an order to a string
func FormatOrder(order *entities.Order) string {
	var itemsStr string
	for _, item := range order.Items {
		itemsStr += fmt.Sprintf("Prato: %s, Amount: %d\n", item.Dish.Name, item.Amount)

	}

	customerName := ""
	if order.CustomerName != nil {
		customerName = *order.CustomerName
	}

	phoneNumber := ""
	if order.PhoneNumber != nil {
		phoneNumber = *order.PhoneNumber
	}

	orderStr := fmt.Sprintf(`
Order ID: %d
Customer Name: %s
Phone Number: %s
Table ID: %d
Items: %s
Status: %s
Created At: %s
Closed At: %s
`, order.ID, customerName, phoneNumber, order.TableID, itemsStr, order.Status, order.CreatedAt, order.ClosedAt)

	return orderStr
}
