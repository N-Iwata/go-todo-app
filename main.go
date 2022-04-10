package main

import (
	"fmt"
	"todo-app/app/controllers"
	"todo-app/app/models"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
