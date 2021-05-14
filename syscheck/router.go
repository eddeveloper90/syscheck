package syscheck

import "xcheck/gin"

func AddRoutes(r *gin.Engine) {
	// api call to all routes and show the results. mvc
	r.GET("syscheck", syscheck)

	//shows uptime per service. lsat 30 days
	//r.GET("api/syscheck/sysinfo", getUserProfileByToken)

	//
	//r.GET("api/syscheck/sysinfo", getUserProfileByToken)
}