/*
Задание №2
Для работы со строками часто используется стандартная библиотека strings. В данном
практическом задании вам понадобятся следующие функции:
// обрезает символы, переданные вторым аргументом, с обеих сторон строки
Trim(s, cutset string) string
// пример
strings.Trim(" hello ", " ") // "hello"
// преобразует все буквы в строке в нижний регистр
strings.ToLower(s string) string
// пример
strings.ToLower("пРиВеТ") // "привет"
// озаглавливает первую букву в каждом слове в строке
strings.Title(s string) string
// пример
strings.Title("привет, джон") // "Привет, Джон"

Реализуйте функцию Greetings(name, string) string, которая вернет строку-приветствие.
Например, при передаче имени Иван, возвращается «Привет, Иван!» Обратите внимание,
что имя может быть написано в разном регистре и содержать пробелы.
*/

package main

import (
  "fmt"
  "strings"
  )

func main(){

  var name string 
  var text string = "Привет, "

  fmt.Printf("Please, enter an name: ")
  fmt.Scanln(&name)
  fmt.Println(Greetings(name, text))
  

}

func Greetings(name string, text string) string{
  name=strings.Trim(name, " ")
  name=strings.ToLower(name)
  name=strings.Title(name)
  return text+" "+name+"!"
}
