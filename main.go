package main

import (
	"fmt"
	"os"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/static"
	"github.com/patrickmn/go-cache"

	"github.com/TheLazarusNetwork/LogTracker/api"
	"github.com/TheLazarusNetwork/LogTracker/utility"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

	viper.AddConfigPath(".")    // Look for config in the working directory
	viper.SetConfigFile(".env") //Load .env file
	viper.SetConfigName(".env") //Load .env file
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	utility.CheckError("Error in reading config file:", err)
	log.Infof("Reading Config File: %s", viper.ConfigFileUsed())
}

func main() {
	log.WithFields(utility.StandardFields).Infof("Lazarus Network LogTracker Version: %s", utility.Version)

	err := viper.ReadInConfig()
	utility.CheckError("Error while reading config file:", err)

	if viper.Get("APP_MODE").(string) == "debug" {
		// set gin release debug
		gin.SetMode(gin.DebugMode)
	} else {
		// set gin release mode
		gin.SetMode(gin.ReleaseMode)
		// disable console color
		gin.DisableConsoleColor()
		// log level info
		log.SetLevel(log.InfoLevel)
	}

	// creates a gin router with default middleware: logger and recovery (crash-free) middleware
	ginApp := gin.Default()

	// cors middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	ginApp.Use(cors.New(config))

	// protection middleware
	ginApp.Use(helmet.Default())

	// add cache storage to gin app
	ginApp.Use(func(ctx *gin.Context) {
		ctx.Set("cache", cache.New(60*time.Minute, 10*time.Minute))
		ctx.Next()
	})

	// serve static files
	ginApp.Use(static.Serve("/", static.LocalFile("./ui", false)))

	// no route redirect to frontend app
	ginApp.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"result": "failure", "message": "Invalid Endpoint Request"})
	})

	// apply public api routes
	api.ApplyRoutes(ginApp, false)

	// apply private api routes
	api.ApplyRoutes(ginApp, true)

	err = ginApp.Run(fmt.Sprintf("%s:%s", viper.Get("SERVER").(string), viper.Get("PORT").(string)))
	if err != nil {
		log.WithFields(utility.StandardFields).WithFields(log.Fields{
			"Error": err,
		}).Fatal("Failed to Start Server")
	}
}
