package main

import "fmt"


func rang(){
  nums := []int{2,3,4}
  sum := 0
  for _,num := range nums {
    sum += num 
  }
  fmt.Println("sum: ",sum)


  kvp := map[string]string{"a": "apple","b":"banana"}
  for k,v := range kvp { // USING RANGE TO ITERATE THROUGH THE MAP
    fmt.Printf("%s -> %s\n",k,v)
  }
}
