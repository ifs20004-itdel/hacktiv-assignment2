package main

import (
	router "assignment2/routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router.StartServer().Run(":8080")
}
