package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"github.com/tonyalaribe/monitor-server/config"
	"github.com/tonyalaribe/monitor-server/constants"
	"github.com/tonyalaribe/monitor-server/logger"
	"github.com/tonyalaribe/monitor-server/models"
	"github.com/tonyalaribe/monitor-server/resources"
)

func GetMainEngine() *gin.Engine {
	// set Production mode if necessary
	if config.Get().IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	//router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
	router.Use(cors.New(cors.Config{
		AllowedOrigins:   config.Get().Cors.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "HEAD", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{constants.CONTENT_TYPE, "X-AUTH-TOKEN"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.Use(static.Serve("/", static.LocalFile("./client", false)))

	//router.StaticFile("/", "./client/index.html")
	//r.Use(static.Serve("/", static.LocalFile("./build", false)))

	router.GET("/ping", func(c *gin.Context) {
		c.Status(200)
	})

	router.GET("/gen4yesterday", func(c *gin.Context) {
		models.DoArchiveForYesterdayWrapper()
	})

	router.Static("/static", "./static")
	api := router.Group("/api")

	resources.Register("screenshots", resources.Screenshots{}, api)
	resources.Register("users", resources.Users{}, api)
	resources.Register("timelog", resources.Timelog{}, api)
	resources.Register("archive", resources.Archive{}, api)

	//resources.Register("/v0.1/login", resources.Users{}, api)

	router.POST("/api/v0.1/add", resources.Screenshots{}.Post)
	router.GET("/api/v0.1/login", resources.Users{}.Post)
	//resources.Register("/advert", resources.Advert{}, apiAuth)

	router.NoRoute(func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./client/index.html")
		c.Abort()
	})

	return router
}

func main() {
	configFile := flag.String("config", "./config.toml", "path to config file")
	flag.Parse()

	config.Init(*configFile)
	logger.Init()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("No Global port has been defined, using default")

		PORT = config.Get().Port

	}

	gocron.Every(1).Day().At("10:30").Do(models.DoArchiveForYesterdayWrapper)
	gocron.Every(1).Day().At("16:30").Do(models.DoArchiveForYesterdayWrapper)

	gocron.Start()

	GetMainEngine().Run(":" + PORT)
}
