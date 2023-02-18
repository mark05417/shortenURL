package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"urlServer/utils"
)

type handler struct {
	store utils.ShortURLStore
}

func (h *handler) ShortenURL(c *gin.Context) {
	var url utils.URL
	if err := c.ShouldBindJSON(&url); err != nil || url.Original == "" {
		errMsg := gin.H{"error": "Invalid URL", "url": url}
		fmt.Println(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}
	short := h.store.Save(url)
	shortURL := c.Request.Host + "/" + short
	c.JSON(http.StatusOK, utils.URL{Original: url.Original, Short: shortURL})
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

func (h *handler) ListURLs(c *gin.Context) {
	urls := h.store.ListURLs()
	for i, _ := range urls {
		urls[i].Short = c.Request.Host + "/" + urls[i].Short
	}
	fmt.Println(urls)
	c.JSON(http.StatusOK, urls)
}

func (h *handler) DeleteURLs(c *gin.Context) {
	h.store.DeleteURLs()
	c.JSON(http.StatusOK, gin.H{})
}

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	router := gin.Default()
	router.Use(cors.New(corsConfig))

	store := &utils.ShortURLStore{Store: make(map[string]utils.URL)}
	h := handler{store: *store}

	router.POST("/api/shorten", h.ShortenURL)
	router.GET("/:short", h.RetrieveURL)
	router.GET("/api/urls", h.ListURLs)
	router.DELETE("/api/urls", h.DeleteURLs)

	if err := router.Run(":8888"); err != nil {
		log.Fatal(err)
	}
}
