package main

import (
	"fmt"
	"github.com/aadi-1024/callisto-golang/routers"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func main() {
	var PORT int

	if len(os.Args) == 1 {
		PORT = 8080
	} else {
		var err error
		PORT, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	
	engine := gin.Default()
	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Get on /")
	})
	routers.Clips(engine)
	engine.Run(fmt.Sprintf("0.0.0.0:%v", PORT))
}
