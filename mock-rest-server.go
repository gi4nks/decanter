package main

import (
	"github.com/gin-gonic/gin"
)

func serve() {
	r := gin.Default()
	r.Use(Cors())
	v1 := r.Group("api/v1")
	{
		v1.POST("/votes", PostVote)
	}
	r.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func PostVote(c *gin.Context) {
	var vote VoteResource
	c.Bind(&vote)

	parrot.Info(vote.String())

	c.JSON(201, vote)
}
