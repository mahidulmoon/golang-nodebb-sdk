package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_nodebb_sdk/models"
	"log"
	"net/http"
)


func CheckUser() gin.HandlerFunc{
	return func(c *gin.Context) {

		var values models.User

		err := c.ShouldBindJSON(&values)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind json",
			})
		}
		//c.JSON(200,gin.H{"message":values,})
		url := "http://127.0.0.1:4567/api/user/email/"+values.Email
		resp,err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		// Print the HTTP Status Code and Status Name
		//fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
		if resp.StatusCode == 200 {
			c.JSON(200,gin.H{
				"message" : "userfound",
				"status" : http.StatusText(resp.StatusCode),
			})
		}else{
			c.JSON(http.StatusInternalServerError,gin.H{
				"message" : "user not found",
			})
		}
	}
}