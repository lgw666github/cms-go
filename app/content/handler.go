package content

import (
	"cmsGo/app/common"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GET(ctx *gin.Context) {

	id := ctx.Param("id")
	row := &Content{}

	// 基于id，查询content内容，将查询的记录map到row中，Content结构体中
	if result := common.DB.Debug().First(row, id); result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "未查到id为" + id + " 的记录",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get请求",
		"code":    "200",
		"data":    row,
	})
}

func PUT(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "put 请求",
		"id":      ctx.Param("id"),
	})
}

func POST(ctx *gin.Context) {

	row := &Content{}
	row.Status = "new"

	//插入
	if result := common.DB.Create(row); result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "添加内容失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "添加内容成功",
	})

}

func DELETE(ctx *gin.Context) {
	id := ctx.Param("id")

	if result := common.DB.Delete(&Content{}, id); result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "删除失败",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "删除id：" + id + " 记录成功",
	})
}

func init() {
	fmt.Println(common.DB)
}
