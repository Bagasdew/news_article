package main

import (
	"errors"
	"fmt"
	"net/http"
	"news/entity"
	"news/usecase"

	"github.com/gin-gonic/gin"
)

var albums = []entity.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func createArticle(c *gin.Context) {
	var article entity.Article
	if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  errors.New("bad request"),
			Type: 0,
		})
	}
	err := usecase.CreateArticle(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.Error{
			Err:  err,
			Type: 0,
		})
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func getArticles(c *gin.Context) {
	var query entity.ArticleParam

	query.Query = c.Query("query")
	query.Author = c.Query("author")

	result, err := usecase.GetArticle(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.Error{
			Err:  err,
			Type: 0,
		})
		return
	}

	//jsonResult, err := json.Marshal(result)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.Error{
	//		Err:  err,
	//		Type: 0,
	//	})
	//}
	c.IndentedJSON(http.StatusOK, result)
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/albums", getAlbums)
	r.POST("/article", createArticle)
	r.GET("/article", getArticles)
	r.Run()
}
