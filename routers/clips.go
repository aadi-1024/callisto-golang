package routers

import (
	"encoding/json"
	"fmt"
	"github.com/aadi-1024/callisto-golang/database"
	"github.com/gin-gonic/gin"
)

func Clips(engine *gin.Engine) {
	clips := engine.Group("/clips")

	clips.GET("all", func(ctx *gin.Context) {
		data := database.Redis.LRange(ctx, "clips", 0, -1)
		ctx.JSON(200, gin.H{
			"Content": data.Val(),
		})
	})

	clips.POST("push", func(ctx *gin.Context) {
		var b map[string]interface{}
		body, err := ctx.GetRawData()

		json.Unmarshal(body, &b)

		if err != nil {
			fmt.Println(err)
			ctx.JSON(400, nil)
		} else {
			err := database.Redis.LPush(ctx, "clips", "hehe").Err()
			if err != nil {
				fmt.Println(err)
				ctx.JSON(400, nil)
			} else {
				ctx.JSON(200, b)
			}
		}
	})
}
