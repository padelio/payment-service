package cmd

// import (
// 	"encoding/base64"
// 	"fmt"
// 	"net/http"
// 	"payment-service/clients"
// 	"payment-service/common/gcs"
// 	"payment-service/common/response"
// 	"payment-service/config"
// 	"payment-service/constants"
// 	"payment-service/controllers"
// 	"payment-service/domain/models"
// 	"payment-service/middlewares"
// 	"payment-service/repositories"
// 	"payment-service/routes"
// 	"payment-service/services"
// 	"time"

// 	"github.com/didip/tollbooth"
// 	"github.com/didip/tollbooth/limiter"
// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// 	"github.com/spf13/cobra"
// )

// var command = &cobra.Command{
// 	Use:   "serve",
// 	Short: "Start the server",
// 	Run: func(c *cobra.Command, args []string) {
// 		_ = godotenv.Load()
// 		config.Init()
// 		db, err := config.InitDatabase()
// 		if err != nil {
// 			panic(err)
// 		}

// 		loc, err := time.LoadLocation("Asia/Jakarta")
// 		if err != nil {
// 			panic(err)
// 		}
// 		time.Local = loc

// 		err = db.AutoMigrate(
// 			&models.Field{},
// 			&models.FieldSchedule{},
// 			&models.Time{},
// 		)
// 		if err != nil {
// 			panic(err)
// 		}

// 		gcs := initGCS()
// 		client := clients.NewClientRegistry()
// 		repository := repositories.NewRepositoryRegistry(db)
// 		service := services.NewServiceRegistry(repository, gcs)
// 		controller := controllers.NewControllerRegistry(service)

// 		router := gin.Default()
// 		router.Use(middlewares.HandlePanic())
// 		router.NoRoute(func(c *gin.Context) {
// 			c.JSON(http.StatusNotFound, response.Response{
// 				Status:  constants.Error,
// 				Message: fmt.Sprintf("Path %s", http.StatusText(http.StatusNotFound)),
// 			})
// 		})
// 		router.GET("/", func(c *gin.Context) {
// 			c.JSON(http.StatusOK, response.Response{
// 				Status:  constants.Success,
// 				Message: "Welcome to Field Service",
// 			})
// 		})
// 		router.Use(func(c *gin.Context) {
// 			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
// 			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-service-name, x-request-at, x-api-key")
// 			if c.Request.Method == "OPTIONS" {
// 				c.AbortWithStatus(204)
// 				return
// 			}
// 			c.Next()
// 		})

// 		lmt := tollbooth.NewLimiter(
// 			config.Config.RateLimiterMaxRequest,
// 			&limiter.ExpirableOptions{
// 				DefaultExpirationTTL: time.Duration(config.Config.RateLimiterTimeSecond) * time.Second,
// 			})
// 		router.Use(middlewares.RateLimiter(lmt))

// 		group := router.Group("/api/v1")
// 		route := routes.NewRouteRegistry(controller, group, client)
// 		route.Serve()

// 		port := fmt.Sprintf(":%d", config.Config.Port)
// 		router.Run(port)
// 	},
// }

// func Run() {
// 	err := command.Execute()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func initGCS() gcs.IGCSClient {
// 	decode, err := base64.StdEncoding.DecodeString(config.Config.GCSPrivateKey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	stringPrivateKey := string(decode)
// 	gcsServiceAccount := gcs.ServiceAccountKeyJSON{
// 		Type:                    config.Config.GCSType,
// 		ProjectID:               config.Config.GCSProjectID,
// 		PrivateKeyID:            config.Config.GCSPrivateKeyID,
// 		PrivateKey:              stringPrivateKey,
// 		ClientEmail:             config.Config.GCSClientEmail,
// 		ClientID:                config.Config.GCSClientID,
// 		AuthURI:                 config.Config.GCSAuthURI,
// 		TokenURI:                config.Config.GCSTokenURI,
// 		AuthProviderX509CertURL: config.Config.GCSAuthProviderX509CertURL,
// 		ClientX509CertURL:       config.Config.GCSClientX509CertURL,
// 		UniverseDomain:          config.Config.GCSUniverseDomain,
// 	}
// 	gcsClient := gcs.NewGCSClient(
// 		gcsServiceAccount,
// 		config.Config.GCSBucketName,
// 	)
// 	return gcsClient
// }
