package server

import "github.com/gin-gonic/gin"

/*
POST BODY

name
alias
description
owner
scm{

}

*/


func CreateProject(c *gin.Context){


	c.JSON(200, gin.H{
		"api": "create project",
	})
}