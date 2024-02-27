package main

import "github.com/gin-gonic/gin"

type V1Router struct{}

func v1Router(parent *gin.RouterGroup) {
	router := parent.Group("")
	r := V1Router{}
	router.GET("/ping", r.helloWorld)
	router.GET("/hls/:hls_id", r.hlsHandler)
	router.GET("/hls/m3u8", r.m3u8Handler)
}

func (v *V1Router) helloWorld(c *gin.Context) {
	responseSuccessWithMessage(c, "Hello world")
}
