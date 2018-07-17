package server

import "github.com/gin-gonic/gin"


/*
POST BODY

name: 运行环境名
type: 运行环境类型  Container or Binary
docker:
	host: 运行环境 ip 地址



*/
func CreateEnvType(c *gin.Context){

	c.JSON(200, gin.H{
		"api": "create envtype",
	})
}