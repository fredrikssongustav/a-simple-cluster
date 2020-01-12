package main

import (
	"fmt"

	"github.com/fredrikssongustav/a-simple-cluster/a-simple-go-app/service"
)

var appName = "a-simple-go-app"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("1487")
}
