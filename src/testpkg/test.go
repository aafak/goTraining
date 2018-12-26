package main

import "fmt"

func main(){

   fmt.Println("Enter first number:")
   var a,b int
   fmt.Scan(&a)
   fmt.Println("Enter second number:")
   fmt.Scan(&b)

   if b==0 {
     panic("Divide by zero")
   }
   c:= a/b
   fmt.Println("Result: ", c) 

   
}