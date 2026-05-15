package handler

import (
	"gateway/genproto/contract"
	"gateway/genproto/user"
	"gateway/genproto/services"
	"log/slog"
	"gateway/redis"

	"github.com/casbin/casbin/v2"
)

type Handler struct {
	UserService user.UserServiceClient
	Log         *slog.Logger
	ProductionService contract.ContractServiceClient
	EmployeeService   services.ServicesServiceClient
	Enforcer    *casbin.Enforcer
	Redis       *redis.RedisRepo
}

func NewHandler(user user.UserServiceClient, production contract.ContractServiceClient, EmployeeService   services.ServicesServiceClient, logger *slog.Logger, Enforcer *casbin.Enforcer, Redis *redis.RedisRepo) *Handler {
	return &Handler{
		UserService: user,
		Log:         logger,
		ProductionService: production,
		EmployeeService: EmployeeService,
		Enforcer:    Enforcer,
		Redis: Redis,
	}
}

// AttendanceSwagger godoc
type AttendanceSwagger struct {
	WorkerId string `json:"worker_id" example:"uuid"`
	WorkDate string `json:"work_date" example:"2024-01-15T00:00:00Z"`
	Status   string `json:"status"    example:"present"`
}

// AttendanceUpdateSwagger godoc
type AttendanceUpdateSwagger struct {
	Id       string `json:"id"        example:"uuid"`
	WorkDate string `json:"work_date" example:"2024-01-15T00:00:00Z"`
	Status   string `json:"status"    example:"present"`
}

// GetDailyAttendanceResSwagger godoc
type GetDailyAttendanceResSwagger struct {
	Attendance []AttendanceSwagger `json:"attendance"`
}

// GetAllAttendanceResSwagger godoc
type GetAllAttendanceResSwagger struct {
	Attendance []AttendanceSwagger `json:"attendance"`
}

// CreateAttendanceResSwagger godoc
type CreateAttendanceResSwagger struct {
	Message string `json:"message" example:"success"`
}

// UpdateAttendanceResSwagger godoc
type UpdateAttendanceResSwagger struct {
	Message string `json:"message" example:"success"`
}

// DeleteAttendanceResSwagger godoc
type DeleteAttendanceResSwagger struct {
	Message string `json:"message" example:"success"`
}