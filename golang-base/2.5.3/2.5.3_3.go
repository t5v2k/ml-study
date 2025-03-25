/*
Задание №3
Реализуйте функцию IsValid(id int, text string) bool, которая проверяет, что
переданный идентификатор id является положительным числом и текст text не пустой.
Пример:
IsValid(0, "hello world") // false
IsValid(-22, "hello world") // false
IsValid(22, "") // false
IsValid(225, "hello world") // true
*/

package main

import (
  "fmt"
//  "strings"
  )

func main(){

  var id int 
  var text string

  fmt.Printf("Please, enter an id: ")
  fmt.Scanln(&id)
  fmt.Printf("Please, enter a text: ")
  fmt.Scanln(&text)
  fmt.Printf("IsValid response: %#v\n", IsValid(id, text))  

}

func IsValid(id int, text string) bool{
  return id>0 && text != ""
}
