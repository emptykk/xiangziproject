package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {

	authList := make(map[int]string)
	var arr1 [1]int
	var userList []int = arr1[:]

	log.WithFields(log.Fields{
		"App": "Console",
	}).Info("Start App...")

	r := gin.Default()
	// Basic Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Wellcome": "Hi There.",
		})
	})
	r.GET("/ping", pingpong)
	r.GET("/regUser/:passToken", func(c *gin.Context) {
		//, authList map[int]string, userList []int
		token := c.Param("passToken")
		fmt.Println(token)
		var num int
		if authList != nil && len(userList) != 0 {
			num = len(userList)
			authList[num] = token
		} else {
			authList[0] = token
		}
		finalCode := "acHad56-ed3e0208d2ac" + token
		log.Infof("User %d Reg Token as %s;", num, finalCode)

		c.JSON(200, gin.H{
			"status": "ok,Your Code is " + finalCode,
		})
	})
	r.Run()
}

func pingpong(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "pong",
	})
}
