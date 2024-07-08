package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CfmReq(c *gin.Context) {
	fmt.Println("Request confirmed!!!!!!!!!!!!!!!!!!!!")

	fmt.Print("method: ")
	fmt.Println(c.Request.Method)
	fmt.Print("url: ")
	fmt.Println(c.Request.URL)
	// fmt.Print("tls ver: ")
	// fmt.Println(c.Request.TLS.Version)
	fmt.Print("header: ")
	fmt.Println(c.Request.Header)
	fmt.Print("body: ")
	fmt.Println(c.Request.Body)
	fmt.Print("url query: ")
	fmt.Println(c.Request.URL.Query())
	fmt.Println()
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "つながってます！"})
} // /v1/test/cfmreq
