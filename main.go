package main

import (
	"fmt"
	//	"os"

	"github.com/yoe7501/blog/internal/config"
)

func main() {
	//path, err := os.UserHomeDir()
	//	path = path + "/.gatorconfig.json"
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
	c.SetUser("Yoel")
	fmt.Println(c)
}
