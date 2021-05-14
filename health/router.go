package health

import (
	"net/http"
	"xcheck/config"
	database "xcheck/db"
	"xcheck/gin"
	util "xcheck/utils"
)

func AddRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"payload": gin.H{
			"uptime":          util.UnixToDuration(config.CONFIG.App.StartUnix),
			"startedDateTime": config.CONFIG.App.StartDateTime,
			"version":         config.CONFIG.App.Version,
			"databaseVersion": database.Version(),
			"active":          gin.Conns.GetConn()},
			"code": 200, 
			"primaryMessage": "OK"})
	})
}
