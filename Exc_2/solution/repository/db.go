package repository

import (
	"ordersystem/model"
	"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// initiliaze db with test data
func NewDatabaseHandler() *DatabaseHandler {

	drinks := []model.Drink{
		{ID: 1, Name: "Iced Caramel Latte", Price: 5.0, Description: "Sweet and refreshing"},
		{ID: 2, Name: "Cola", Price: 2.0, Description: "Classic soft drink"},
		{ID: 3, Name: "Water", Price: 1.0, Description: "Still water"},
	}

	// Init orders slice with some test data

	orders := []model.Order{
		{DrinkID: 1, Amount: 1, CreatedAt: time.Now()},
		{DrinkID: 2, Amount: 3, CreatedAt: time.Now()},
		{DrinkID: 3, Amount: 5, CreatedAt: time.Now()},
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo // key = DrinkID, value = Amount of orders
// totalledOrders map[uint64]uint64
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {

	totalledOrders := make(map[uint64]uint64)
	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += uint64(order.Amount)
	}

	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	db.orders = append(db.orders, *order)
}
