package main

import (
	"github.com/oaraujocesar/donates-control-api/configs"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
}
