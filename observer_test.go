package patterns

func ExampleSubscriber() {
	shirtItem := NewProduct("Go React")
	firstCustomer := &Customer{id: "vue@yahoo.com"}
	secondCustomer := &Customer{id: "angular@yahoo.com"}
	shirtItem.Register(firstCustomer)
	shirtItem.Register(secondCustomer)
	shirtItem.UpdateAvailability()
	// Output:
	// Tell customer vue@yahoo.com for new product Go React
	// Tell customer angular@yahoo.com for new product Go React
}
