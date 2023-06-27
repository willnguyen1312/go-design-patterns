package patterns

import "fmt"

// Subject interface
type Subject interface {
	Register(Subscriber)
	Deregister(Subscriber)
	NotifyAll()
}

// Product struct
type Product struct {
	SubscriberList []Subscriber
	name           string
	inStock        bool
}

// NewProduct function create a brand-new Product
func NewProduct(name string) *Product {
	return &Product{
		name: name,
	}
}

// UpdateAvailability method on Product struct
func (i *Product) UpdateAvailability() {
	i.inStock = true
	i.NotifyAll()
}

// Register method on Product struct
func (i *Product) Register(o Subscriber) {
	i.SubscriberList = append(i.SubscriberList, o)
}

// Deregister method on Product struct
func (i *Product) Deregister(o Subscriber) {
	i.SubscriberList = removeFromslice(i.SubscriberList, o)
}

// NotifyAll method on Product struct
func (i *Product) NotifyAll() {
	for _, Subscriber := range i.SubscriberList {
		Subscriber.Update(i.name)
	}
}

func removeFromslice(SubscriberList []Subscriber, observerToRemove Subscriber) []Subscriber {
	observerListLength := len(SubscriberList)
	for i, Subscriber := range SubscriberList {
		if observerToRemove.GetID() == Subscriber.GetID() {
			SubscriberList[observerListLength-1], SubscriberList[i] = SubscriberList[i], SubscriberList[observerListLength-1]
			return SubscriberList[:observerListLength-1]
		}
	}
	return SubscriberList
}

// Subscriber interface
type Subscriber interface {
	Update(string)
	GetID() string
}

// Customer struct
type Customer struct {
	id string
}

// Update method on Custome struct
func (c *Customer) Update(itemName string) {
	fmt.Printf("Tell customer %s for new product %s\n", c.id, itemName)
}

// GetID method on Customer struct
func (c *Customer) GetID() string {
	return c.id
}
