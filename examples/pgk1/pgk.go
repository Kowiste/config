package pgk1

import (
	"fmt"

	conf "example/config"

	"github.com/kowiste/config"
)

func Test() {
	cfg, err := config.Get[conf.MyConfig]()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	fmt.Println("Package 1", cfg)
}
