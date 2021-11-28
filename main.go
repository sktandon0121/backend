package main

import (
	"fmt"

	"github.com/sktandon0121/backend/controllers"
	"github.com/sktandon0121/backend/repo"
)

func main() {
	repo.InitDBConnection()
	controllers.StartServer()
	fmt.Println("THis is main")
}
