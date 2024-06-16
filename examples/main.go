package main

import (
	"fmt"

	"github.com/kowiste/config"
	conf "example/config"
	pkg1 "example/pgk1"
	pkg2 "example/pgk2"

)



func main() {
	err := config.New[conf.MyConfig](config.GetPathEnv("dev"))
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	cfg, err := config.Get[conf.MyConfig]()
	if err != nil {
		fmt.Println("Error getting config:", err)
		return
	}
	fmt.Println("main",cfg)

	pkg1.Test()
	pkg2.Test()
}
