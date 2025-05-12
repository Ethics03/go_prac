package main


//ANONYMOUS FUNCTIONS
//DEFINING FUNCTIONS IN-LINE WITHOUT NAMING THEM


import "fmt"


func intSq() func() int { //THIS FUNCTION BASICALLY RETURNS ANOTHER FUNCTION
  i := 0;
  return func() int {
     i++
     return i
  }
}

func closure(){
  nextInt := intSq()

  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  nextInts := intSq()

  fmt.Println(nextInts())
}

