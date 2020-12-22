package api

import (
	"bytes"
	"encoding/json"
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
		url := models.Domain+"/api/user/email/"+values.Email
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
			//values.Password = "upskillpass" //defaultpass
			json_fmt := map[string]string{"username": values.Username,"email": values.Email,"password": "upskillpass"}
			value,_ := json.Marshal(json_fmt)

			//fmt.Println(bytes.NewBuffer(value)) //{"email":"upskill@gmail.com","username":"testuser","password":"upskillpass"}

			url = models.Domain+"/api/v2/users/"
			res,err := http.Post(url,"application/json",bytes.NewBuffer(value))

			if err != nil{
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError,gin.H{
					"error":"cannot create new user for nodebb",
				})
			}else{
				c.JSON(http.StatusOK,gin.H{
					"message":"New user created for nodebb",
					"nodebb" : res.StatusCode,
				})
			}


		}
	}
}