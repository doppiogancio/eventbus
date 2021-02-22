```go
b := bus.EventBus()

cart := subscriber.NewCart()

b.Subscribe(event.ItemAddedToCartEvent, cart)
b.Publish(event.AddItemToCart(model.NewItem("Iphone X", 1000.00)))
b.Publish(event.AddItemToCart(model.NewItem("USB Plug", 12.06)))

fmt.Println(cart.Total())
```