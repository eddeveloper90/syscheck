package main

import (
	"fmt"
	"net/http"
	"os"
	"xcheck/auth"
	"xcheck/config"
	database "xcheck/db"
	"xcheck/gin"
	"xcheck/health"
	"xcheck/syscheck"
	"xcheck/utils"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// check for enabling debug mode or not
	if config.CONFIG.App.StageMode == config.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")
		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	health.AddRoutes(r)
	auth.AddRoutes(r)
	syscheck.AddRoutes(r)
	return r
}

func cmdCheck() {
	args := os.Args
	if len(args) >= 2 {
		arg1 := args[1]
		if arg1 == "version" {
			fmt.Println(config.CONFIG.App.Version)
			os.Exit(0)
		}

		if arg1 == "pidfile" {
			fmt.Println(config.CONFIG.App.PidFile)
			os.Exit(0)
		}

		if arg1 == "config" {
			utils.JsonPrettyPrint(config.CONFIG)
			os.Exit(0)
		}

		if arg1 == "db" {
			fmt.Println(database.Version())
			os.Exit(0)
		}
	}
}

func main() {
	config.LoadConfig()
	cmdCheck()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + config.CONFIG.HttpServer.Port)
}
