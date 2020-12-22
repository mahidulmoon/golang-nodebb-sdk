package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_nodebb_sdk/models"
	"net/http"
)

func CheckGroupExist() gin.HandlerFunc{
	return func(c *gin.Context){
		var group models.Group
		err := c.ShouldBindJSON(&group)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind json",
			})
		}

		url := models.Domain+"/api/groups/"+group.Name

		resp,_ := http.Get(url)

		if resp.StatusCode != 200{
			//c.JSON(http.StatusNotFound,gin.H{
			//	"error":"workshop not found",
			//})

			data := map[string]string{"name" : group.Name}

			jsonData,_ := json.Marshal(data)

			url = models.Domain+"/api/v2/groups/"
			res,err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
			if err != nil{
				fmt.Println(err)
			}else{
				fmt.Println(res)
			}


		}else{
			fmt.Println("workshop found")

		}



	}
}