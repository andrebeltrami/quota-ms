package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Quota struct {
	ID string `json:"id"`
	Flavor Flavor `json:"flavor"`
	MaxMem int `json:"maxMem"`
	MaxCores int `json:"maxCores"`
}

type Flavor string

const (
	Small Flavor = "small"
	Medium Flavor = "medium"
	Big Flavor = "big"
)

var db = []Quota{
	{
		ID: "1",
		Flavor: Small,
		MaxMem: 4,
		MaxCores: 1,
	},
	{
		ID: "2",
		Flavor: Medium,
		MaxMem: 8,
		MaxCores: 2,
	},
	{
		ID: "3",
		Flavor: Big,
		MaxMem: 16,
		MaxCores: 4,
	},
}

func postQuotas(c *gin.Context) {
    var newQuota Quota

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
    	c.AbortWithError(400, err)
   		return
	}

	err = json.Unmarshal(body, &newQuota)
	if err != nil {
    	c.AbortWithError(400, err)
    	return
	}

    db = append(db, newQuota)
    c.IndentedJSON(http.StatusCreated, newQuota)
}

func getQuotaByID(c *gin.Context) {
	quotaID := c.Param("quotaID")

	for i := range db {
		if db[i].ID == quotaID {
			c.IndentedJSON(http.StatusOK, db[i])
			return
		}
	}
	c.String(404, "Quota not find.")
}


func main(){

	router := gin.Default()

	router.GET("/quotas", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, db)
	})

	router.GET("/quota/:quotaID", getQuotaByID)

	router.POST("/quota", postQuotas)


	router.Run("localhost:8081")

}