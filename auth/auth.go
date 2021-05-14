package auth

import (
	"net/http"
	"xcheck/gin"
)

// get profile
func getUserProfileByToken(c *gin.Context) {
	// Parse JSON
	var body struct {
		Token string `json:"token" binding:"required"`
	}

	if c.Bind(&body) == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"token":  body.Token})
	}
}
