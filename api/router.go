package api

import (
	"gateway/api/handler"
	"gateway/api/middleware"
	"gateway/config"
	"gateway/redis"
	"log/slog"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	SetupRoutes(handler.Handler, *slog.Logger)
	StartServer(config.Config) error
}

type controllerImpl struct {
	Port     string
	Router   *gin.Engine
	Redis    *redis.RedisRepo
	Enforcer *casbin.Enforcer
}

func NewController(router *gin.Engine, rdb *redis.RedisRepo, enf *casbin.Enforcer) Controller {
	return &controllerImpl{
		Router:   router,
		Redis:    rdb,
		Enforcer: enf,
	}
}

func (c *controllerImpl) StartServer(cfg config.Config) error {
	c.Port = cfg.API_GATEWAY
	return c.Router.Run(c.Port)
}

func (c *controllerImpl) SetupRoutes(h handler.Handler, logger *slog.Logger) {
	// Middleware instancelarini yaratish
	authMiddleware := middleware.NewAuthMiddleware(c.Redis)
	casbinMiddleware := middleware.NewCasbinMiddleware(c.Enforcer)

	router := c.Router.Group("/api")
	router.Use(authMiddleware.Check())
	router.Use(casbinMiddleware.Middleware())

	users := router.Group("/user")
	{
		users.GET("/getbyuser/:email", h.GetUSerByEmail)
		users.PUT("/update_password", h.UpdatePassword)
		users.DELETE("/delete_user/:id", h.DeleteUser)
		users.PUT("/update_role", h.UpdateRole)
		users.PUT("/image_update", h.ProfileImage)
		users.GET("/all_users/:limit/:page", h.GetAllUsers)
		users.POST("/logout", h.Logout)
	}

	contract := router.Group("/contract")
	{
		contract.POST("/newcontract", h.NewContract)
		contract.PUT("/contract_update", h.NewContractUpdate)
		contract.DELETE("/contract_delete/:id", h.NewContractDelete)
		contract.GET("/get_name/:name", h.NewContractGetName)
		contract.GET("/all_contract/:limit/:page", h.NewContractGetAll)

		contract.POST("/inside_contract", h.NewInsideTheContract)
		contract.PUT("/insidecontract_update", h.NewInsideTheContractUpdate)
		contract.DELETE("/insidecontract_delete/:id", h.NewInsideTheContractDelete)
		contract.GET("/all_insidecontract/:limit/:page", h.NewInsideTheContractGetAll)
	}
}