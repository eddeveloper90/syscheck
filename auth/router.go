package auth

import (
	"xcheck/gin"
)

func AddRoutes(r *gin.Engine) {
	r.POST("api/security/v1/profile/get", getUserProfileByToken)
}
