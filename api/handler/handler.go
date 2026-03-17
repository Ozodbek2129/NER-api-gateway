package handler

import (
	"gateway/genproto/ishlab_chiqarish"
	"gateway/genproto/user"
	"log/slog"
	"gateway/redis"

	"github.com/casbin/casbin/v2"
)

type Handler struct {
	UserService user.UserServiceClient
	Log         *slog.Logger
	ProductionService ishlab_chiqarish.IshlabChiqarishServiceClient
	Enforcer    *casbin.Enforcer
	Redis       *redis.RedisRepo
}

func NewHandler(user user.UserServiceClient, production ishlab_chiqarish.IshlabChiqarishServiceClient, logger *slog.Logger, Enforcer *casbin.Enforcer, Redis *redis.RedisRepo) *Handler {
	return &Handler{
		UserService: user,
		Log:         logger,
		ProductionService: production,
		Enforcer:    Enforcer,
		Redis: Redis,
	}
}
