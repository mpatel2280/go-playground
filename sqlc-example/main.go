package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"sqlc-example/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/sqlc_db?parseTime=true&multiStatements=true"
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queries := db.New(conn) // sqlc-generated queries

	r := gin.Default()

	// ---- Classes ----
	r.POST("/classes", func(c *gin.Context) {
		var req struct {
			Name string `json:"name"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := queries.CreateClass(context.Background(), req.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Class created successfully"})
	})

	r.GET("/classes", func(c *gin.Context) {
		classes, err := queries.ListClasses(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, classes)
	})

	r.Run(":8080")
}
