package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

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

	// ---- Students ----
	r.POST("/students", func(c *gin.Context) {
		var req struct {
			Name    string `json:"name"`
			ClassID int32  `json:"class_id"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := queries.CreateStudent(context.Background(), db.CreateStudentParams{
			Name:    req.Name,
			ClassID: sql.NullInt32{Int32: req.ClassID, Valid: req.ClassID != 0},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "Student created successfully",
			"name":     req.Name,
			"class_id": req.ClassID,
		})
	})

	r.GET("/students/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}
		student, err := queries.GetStudentWithClass(context.Background(), int32(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, student)
	})

	// ---- Assignments ----
	r.POST("/assignments", func(c *gin.Context) {
		var req struct {
			Title     string `json:"title"`
			StudentID int32  `json:"student_id"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := queries.CreateAssignment(context.Background(), db.CreateAssignmentParams{
			Title:     req.Title,
			StudentID: sql.NullInt32{Int32: req.StudentID, Valid: req.StudentID != 0},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":    "Assignment created successfully",
			"title":      req.Title,
			"student_id": req.StudentID,
		})
	})

	r.GET("/students/:id/assignments", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}
		assignments, err := queries.ListAssignmentsByStudent(context.Background(), sql.NullInt32{Int32: int32(id), Valid: true})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, assignments)
	})

	r.Run(":8080")
}
