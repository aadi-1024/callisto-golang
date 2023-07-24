package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/aadi-1024/callisto-golang/database"
)

func Clips(engine *gin.Engine) {
	clips := engine.Group("/clips")
	clips.GET("all", func(ctx *gin.Context) {
		data := database.Redis.LRange(ctx, "clips", 0, -1)
		ctx.JSON(200, gin.H {
			"Content": data.Val(),
		})
	})
	clips.POST("push", func(ctx *gin.Context) {
		// var b []byte
		_, err := ctx.GetRawData()
		if err != nil {
			fmt.Println(err)
			ctx.JSON(400, nil)
		} else {
			database.Redis.LPush(ctx, "clips", "hehe")
		}
	})
}