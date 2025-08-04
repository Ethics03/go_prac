package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
  ID int
  Status string
  
}

func main(){

  orders := generateOrders(10)

  go  processOrders(orders)

  go updateOrderStatuses(orders)

  go reportOrders(orders)
  fmt.Printf("All operations done successfully!")
}

func updateOrderStatuses(orders []*Order){
    for _, order := range orders {
      time.Sleep(
        time.Duration(rand.Intn(500))*time.Millisecond,)
        status := []string{"Processing","Shipped","Delivered",}[rand.Intn(3)]
        order.Status = status
        fmt.Printf("Updated Order %d status : %s\n",order.ID,status)
    }

}

func reportOrders(orders []*Order){
   for i := 0 ; i < 5 ; i++ {
     time.Sleep(1*time.Second)
     fmt.Println("\n---Order Status Report---")
     for _, order := range orders {
       fmt.Printf("Order %d: %s\n",order.ID,order.Status)
     }
     fmt.Println("------------------\n")
   }
}
func processOrders(orders []*Order){
  for _, order := range orders {
      time.Sleep(
        time.Duration(rand.Intn(500))*time.Millisecond,)
      fmt.Printf("Procesing order %d\n",order.ID)
  }
}

func generateOrders(count int) []*Order {
  orders := make([]*Order, count) //order slice with empty elements
    
  for i := 0 ; i < count; i++ {
    orders[i]  = &Order{ID: i+1, Status: "Pending"}
  }
  return orders 


}
