package cmd

import (
	"taro/config"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)



func registerCloudAPIs(r *gin.Engine){

}

func startWeb() {

	dbURL := config.GetString("mongodb_url")
	//DatabaseName := config.GetString("database_name")

	r := gin.Default()
	//r.Static("img", "./html/img")
	//r.Static("js", "./html/js")
	//r.Static("css", "./html/css")
	//r.Static("fonts", "./html/fonts")
	//r.Static("bower_components", "./html/bower_components")
	//r.Static("layui", "./html/layui")
	//r.Static("layer", "./html/layer")
	//
	//r.LoadHTMLGlob("html/*.html")
	r.Use(func(c *gin.Context) {
		session, err := mgo.Dial(dbURL)
		session.SetMode(mgo.Monotonic, true)

		if err != nil {
			panic(err)
		}

		c.Set("session", session)

		c.Next()

		session.Close()

	})

	registerCloudAPIs(r)

	r.Run(":8080")
}
