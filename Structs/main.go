package main

import (
	"encoding/json"
	"fmt"

	"github.com/pnaskardev/go-mastery/structs/models"
)

func main() {
	fmt.Println("HELLO FROM STRUCTS")

	b := []byte(`{
		"id": "ORD_123",
		"status": "CONFIRMED",
		"total_amount": 2499.50,
		"currency": "INR",
		"created_at": "2025-12-29T14:32:00Z",
		"customer": {
			"id": "CUST_42",
			"name": "Priyanshu",
			"email": "priyanshu@example.com"
		},
		"items": [
			{
				"item_id": "ITEM_1",
				"name": "Fast Charger",
				"quantity": 1,
				"price": 1999.50
			},
			{
				"item_id": "ITEM_2",
				"name": "Cable",
				"quantity": 1,
				"price": 500.00
			}
		],
		"payment": {
			"method": "UPI",
			"status": "SUCCESS",
			"transaction_id": "TXN_999"
		}
		}`)

	var order_instance models.Order

	if err := json.Unmarshal(b, &order_instance); err != nil {
		panic(err)
	}

	fmt.Println(order_instance.ID)
	fmt.Println(order_instance.Customer.Name)
	fmt.Println(order_instance.Items[0].Name)

	server_instance := models.NewServer(3000)
	port := server_instance.Port()
	server_instance.Start()
	fmt.Println(port)

	// Child And Base Struct
	// Java Method Overriding does not works in Golang
	// In here Base has a different Describe and Child has a different Describe
	child_instance := models.NewChildInstance(23, 65)
	child_instance.Describe()
	child_instance.Base.Describe()

	// Method Promotion

	flying_car_instance := models.FlyingCar{
		Flyable:  models.Flyable{},
		Drivable: models.Drivable{},
	}

	flying_car_instance.Drive()
	flying_car_instance.Fly()

}
