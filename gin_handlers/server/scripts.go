package server
import (
	"github.com/gin-gonic/gin"
	"taro/util"
	//"strings"

)
/*
GET /scripts/pinyin/:hanzi
Response

*/
func GetPinyin(c *gin.Context){
	hanzi := c.Param("hanzi")

	py := util.Hanzi2pinyin(hanzi)
	c.JSON(200, gin.H{
		"pinyin": py,
	})
}