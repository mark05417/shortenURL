package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"

	"urlServer/utils"
)

type handler struct {
	Store utils.Store
}

func (h *handler) ShortenURL(c *gin.Context) {
	var url utils.URL
	if err := c.ShouldBindJSON(&url); err != nil || url.Original == "" {
		errMsg := gin.H{"error": "Invalid URL", "url": url}
		fmt.Println(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}
	short := h.Store.Save(url)
	shortURL := c.Request.Host + "/" + short
	c.JSON(http.StatusOK, utils.URL{Original: url.Original, Short: shortURL})
}

func (h *handler) RetrieveURL(c *gin.Context) {
	short := c.Param("short")
	url, ok := h.Store.Retrieve(short)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	h.Store.IncreaseCount(short)
	// fmt.Println(short)
	// fmt.Println(url)
	c.Redirect(http.StatusFound, url.Original)
}

func (h *handler) ListURLs(c *gin.Context) {
	urls := h.Store.ListURLs()
	for i := range urls {
		urls[i].Short = c.Request.Host + "/" + urls[i].Short
	}
	// fmt.Println(urls)
	c.JSON(http.StatusOK, urls)
}

func (h *handler) DeleteURL(c *gin.Context) {
	short := c.Param("short")
	h.Store.DeleteURL(short)
	// fmt.Println(short)
	c.JSON(http.StatusOK, gin.H{})
}

func (h *handler) DeleteURLs(c *gin.Context) {
	h.Store.DeleteURLs()
	c.JSON(http.StatusOK, gin.H{})
}

var dbType = "mysql"

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router := gin.Default()
	router.Use(cors.New(corsConfig))

	ctx := context.Background()
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	cfg := mysql.NewConfig()
	cfg.User = "myuser"
	cfg.Passwd = "mypassword"
	cfg.DBName = "mydatabase"
	sqlDB, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	storeMemory := utils.ShortURLStore{Store: make(map[string]utils.URL)}
	storeRedis := utils.ShortURLStoreRedis{DB: redisDB, Ctx: ctx}
	storeMysql := utils.ShortURLStoreMysql{DB: sqlDB}

	var h handler
	if dbType == "memory" {
		h = handler{Store: &storeMemory}
	} else if dbType == "redis" {
		h = handler{Store: &storeRedis}
	} else if dbType == "mysql" {
		h = handler{Store: &storeMysql}
	} else {
		log.Fatal("Unknown db type")
	}

	router.POST("/api/shorten", h.ShortenURL)
	router.GET("/:short", h.RetrieveURL)
	router.GET("/api/urls", h.ListURLs)
	router.DELETE("/api/:short", h.DeleteURL)
	router.DELETE("/api/urls", h.DeleteURLs)

	if err := router.Run(":8888"); err != nil {
		log.Fatal(err)
	}
}
