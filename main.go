package main

import (
	"fmt"
	"kk/config"
)

func main() {
	conf := config.NewConfig()

	fmt.Println(conf)
	return
}
