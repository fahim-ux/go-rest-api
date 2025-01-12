package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

type profile struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Dept string `json:"dept"`
	CGPA float64 `json:"cgpa"`
}

var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var profiles = []profile{
	{ID:"22U10360", Name: "Fahim",Dept: "CSE",CGPA: 8.64},
	{ID:"22U10268", Name: "Subhodipa",Dept: "CSE",CGPA: 8.9},
	{ID:"22U10376", Name: "Sahitya",Dept: "BT",CGPA: 8.64},
}

func getAlbums(c * gin.Context){
	c.IndentedJSON(http.StatusOK,albums)
}

func getAlbumByID(c *gin.Context){
	id := c.Param("id")
	for _,a := range albums{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK,a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"album not found"})
}


func getProfiles(c * gin.Context){
	c.IndentedJSON(http.StatusOK,profiles)
}

func addProfile(c * gin.Context){
	var newprofile profile	
	if err:= c.BindJSON(&newprofile); err!=nil{
		return
	}

	profiles = append(profiles, newprofile)

	fmt.Println("New profile added.......")
	c.IndentedJSON(http.StatusCreated,newprofile)
}


func main(){
	router := gin.Default()
	router.GET("/albums",getAlbums)
	router.GET("/albums/:id",getAlbumByID)
	router.GET("/profile",getProfiles)
	router.POST("/profile",addProfile)

	router.Run("localhost:8800")
}