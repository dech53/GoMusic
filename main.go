package main

import (
	"GoMusic/handler"
	"GoMusic/misc/log"
	_ "GoMusic/repo/db"
	"fmt"
)

func main() {
	r := handler.NewRouter()
	if err := r.Run(fmt.Sprintf(":%d", 8081)); err != nil {
		log.Errorf("fail to run server: %v", err)
		panic(err)
	}
}
