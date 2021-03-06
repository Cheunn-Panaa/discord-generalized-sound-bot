package server

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lon9/discord-generalized-sound-bot/backend/config"
	"github.com/lon9/discord-generalized-sound-bot/backend/controllers"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// NewRouter returns gin router
func NewRouter() (*gin.Engine, error) {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	cfg := config.GetConfig()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.GetStringSlice("server.cors"),
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Admin zone",
		Key:         []byte(config.GetConfig().GetString("auth.secret")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(string); ok {
				return jwt.MapClaims{
					"id": v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return claims["id"].(string)
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {

			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			isSuccess := func(password string) bool {
				if err := bcrypt.CompareHashAndPassword([]byte(cfg.GetString("auth.password")), []byte(password)); err != nil {
					return false
				}
				return true
			}
			if loginVals.Username == cfg.GetString("auth.username") && isSuccess(loginVals.Password) {
				return loginVals.Username, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(string); ok && v == cfg.GetString("auth.username") {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"status":  code,
				"message": message,
			})
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		return nil, err
	}

	health := new(controllers.HealthController)
	categories := new(controllers.CategoriesController)
	sounds := new(controllers.SoundsController)

	router.GET("/health", health.Status)

	categoriesRoutes := router.Group("/categories")
	{
		categoriesRoutes.GET("/", categories.Index)
		categoriesRoutes.GET("/:id", categories.Show)
	}

	soundRoutes := router.Group("/sounds")
	{
		soundRoutes.GET("/", sounds.Index)
	}

	router.POST("/login", authMiddleware.LoginHandler)
	admin := router.Group("/admin")
	admin.Use(authMiddleware.MiddlewareFunc())
	{
		admin.POST("/sounds", sounds.Create)
		admin.GET("refresh_token", authMiddleware.RefreshHandler)
	}

	return router, nil
}
