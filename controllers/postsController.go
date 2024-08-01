package controller

import (
	"database/sql"
	"net/http"
	"postsapi/initializers"
	"postsapi/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var currentTime = time.Now().Format("2006-01-02 15:04:05")

// ListPosts this function return all posts
func ListPosts(c *gin.Context) {
	query := `SELECT * FROM posts`

	rows, err := initializers.DB.Query(query)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	posts := make([]*models.Post, 0)
	for rows.Next() {
		post := new(models.Post)
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Date)
		if err != nil {
			c.IndentedJSON(500, gin.H{
				"code":  http.StatusInternalServerError,
				"error": err.Error(),
			})
			return
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		c.IndentedJSON(500, gin.H{
			"code":  500,
			"error": err.Error(),
		})
	}

	c.IndentedJSON(200, posts)
}

// FindPost this function find unique post using ID
func FindPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	idInt, _ := strconv.ParseInt(id, 10, 64)

	err := initializers.DB.QueryRow("SELECT * FROM posts WHERE id = ?", idInt).Scan(&post.ID, &post.Title, &post.Content, &post.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(404, gin.H{
				"code":  404,
				"error": err.Error(),
			})
			return
		}

		c.IndentedJSON(500, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, post)
}

// CreatePost this function create a new post
func CreatePost(c *gin.Context) {
	var posts models.RequestPost

	if err := c.BindJSON(&posts); err != nil {
		c.IndentedJSON(500, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}

	query := `INSERT INTO posts(title, content, date) VALUES (?, ?, ?)`
	posts.Date = currentTime

	_, err := initializers.DB.Exec(query, posts.Title, posts.Content, posts.Date)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, gin.H{
		"code":    200,
		"message": "Post created sucessfully!",
	})
}

// DeletePost this function delete a post
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.ParseInt(id, 10, 64)

	query := `DELETE FROM posts WHERE id = ?`
	_, err := initializers.DB.Exec(query, intId)
	if err != nil {
		c.IndentedJSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, gin.H{
		"code":    200,
		"message": "User deleted sucessfully!",
	})
}

// UpdatePost this function update a post
func UpdatePost(c *gin.Context) {
	var post models.Post
	query := `UPDATE posts SET title = ?, content = ?, date = ? WHERE id = ?`

	if err := c.BindJSON(&post); err != nil {
		c.IndentedJSON(404, gin.H{
			"code":  404,
			"error": err.Error(),
		})
		return
	}

	post.Date = currentTime

	_, err := initializers.DB.Exec(query, post.Title, post.Content, post.Date, post.ID)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, gin.H{
		"code":    200,
		"message": "Post updated sucessfully!",
	})
}
