package handler

import (
	"backend/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := NewsfeedPostRequest{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
