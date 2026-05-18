package api

import (
	"gateway/api/handler"
	"gateway/api/middleware"
	"gateway/config"
	"gateway/redis"
	"log/slog"

	_ "gateway/api/docs"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gateway API
// @version 1.0
// @description Gateway service with Casbin & JWT
// @host 192.168.0.44:4030
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

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
	c.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		users.PUT("/image_update/:email", h.ProfileImage)
		users.GET("/all_users", h.GetAllUsers)
		users.POST("/logout", h.Logout)
	}

	contract := router.Group("/contract")
	{
		contract.POST("/newcontract", h.NewContract)
		contract.PUT("/contract_update", h.NewContractUpdate)
		contract.DELETE("/contract_delete/:id", h.NewContractDelete)
		contract.GET("/get_name/:name", h.NewContractGetName)
		contract.GET("/all_contract", h.NewContractGetAll)

		contract.POST("/inside_contract", h.NewInsideTheContract)
		contract.PUT("/insidecontract_update", h.NewInsideTheContractUpdate)
		contract.DELETE("/insidecontract_delete/:id", h.NewInsideTheContractDelete)
		contract.GET("/all_insidecontract", h.NewInsideTheContractGetAll)
	}

	employee := router.Group("/employee")
	{
		employee.POST("/creategroup", h.CreateGroup)
		employee.PUT("/updategroup", h.UpdateGroup)
		employee.DELETE("/deletegroup/:id", h.DeleteGroup)
		employee.GET("/getallgroup", h.GetAllGroup)

		employee.POST("/createworker", h.CreateWorker)
		employee.PUT("/updateworker", h.UpdateWorker)
		employee.DELETE("/deleteworker/:id", h.DeleteWorker)
		employee.GET("/getallworker", h.GetAllWorker)

		employee.POST("/createattendace", h.CreateAttendance)
		employee.PUT("/updateattendace", h.UpdateAttendance)
		employee.DELETE("/deleteattendance", h.DeleteAttendance)
		employee.GET("/getdailyattendace", h.GetDailyAttendance)
		employee.GET("/getallattendace", h.GetAllAttendance)

		employee.POST("/createtask", h.CreateTask)
		employee.PUT("/updatetask", h.UpdateTask)
		employee.DELETE("/deletetask/:id", h.DeleteTask)
		employee.GET("/getalltask", h.GetAllTask)
	}
}