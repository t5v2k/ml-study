/*
Задание №1
Напишите программу, которая последовательно выполняет следующие операции с введенным числом:
• Число умножается на 2.
• К числу прибавляется 20.
*/

package main

import "fmt"

func main(){

  var i int //integer nimber
  var m int = 2 //multiplier
  var s int = 20 //summand

  fmt.Printf("Please, enter an integer: ")
  fmt.Scanln(&i)
  
  
  fmt.Printf("%d * %d = %d\n", i, m, multiplication(i, m))
  fmt.Printf("%d + %d = %d\n", i, s, summ(i, s))

}

func multiplication (i int, m int) int{
  return i*m
}

func summ (i int, s int) int{
  return i+s
}
