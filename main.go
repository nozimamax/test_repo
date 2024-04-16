package main

import (
	"fmt"

	"github.com/nozimamax/test_repo/user"
)
func main(){
    fmt.Println("name: ")
    var name string
    fmt.Scan(&name)
    fmt.Println("age: ")
    var age int
    fmt.Scan(&age)
    fmt.Println("email: ")
    var email string
    fmt.Scan(&email)
    a := user.User{Name: name, Age: age, Email: email}
    m, e := a.ValidateUser()
    if m == "Errors: " {
      println("Errorrs: \n")
      for _, err := range e {
        fmt.Println(err.Error())
      }
    } else {
      fmt.Println(a.ValidateUser())
    }
}
