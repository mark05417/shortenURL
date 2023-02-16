package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"urlServer/utils"
)

type handler struct {
	store utils.ShortURLStore
}

func (h *handler) ShortenURL(c *gin.Context) {
	var url utils.URL
	if err := c.ShouldBindJSON(&url); err != nil || url.Original == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL", "url": url})
		return
	}
	short := h.store.Save(url)
	shortURL := c.Request.Host + "/" + short
	c.JSON(http.StatusOK, utils.URL{Original: url.Original, Short: shortURL})
}

func (h *handler) ListURLs(c *gin.Context) {

}

func (h *handler) RetrieveURL(c *gin.Context) {
	short := c.Param("short")
	url, ok := h.store.Retrieve(short)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	fmt.Println(short)
	fmt.Println(url)
	c.Redirect(http.StatusFound, url.Original)
}

func main() {
	router := gin.Default()
	store := &utils.ShortURLStore{Store: make(map[string]utils.URL)}
	h := handler{store: *store}

	router.POST("/api/shorten", h.ShortenURL)
	router.GET("/api/list", h.ListURLs)
	router.GET("/:short", h.RetrieveURL)

	if err := router.Run(":8888"); err != nil {
		log.Fatal(err)
	}
}
