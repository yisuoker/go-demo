package main

import (
	"fmt"
	"os"

	"github.com/yisuoker/go-demo/conf"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing config file. eg: go run main config.yaml")
		return
	}

	if err := conf.Init(os.Args[1]); err != nil {
		fmt.Printf("Failed to load the configuration file, err: %v\n", err)
		return
	}
	fmt.Printf("%#v\n", conf.Cfg)
	fmt.Println("running...")
}
