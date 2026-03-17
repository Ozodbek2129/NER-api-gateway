package handler

import (
	"gateway/genproto/ishlab_chiqarish"
	"gateway/genproto/user"
	"log/slog"

	"github.com/casbin/casbin/v2"
)

type Handler struct {
	UserService user.UserServiceClient
	Log         *slog.Logger
	DocsService ishlab_chiqarish.IshlabChiqarishServiceClient
	Enforcer    *casbin.Enforcer
}

func NewHandler(user user.UserServiceClient, production ishlab_chiqarish.IshlabChiqarishServiceClient, logger *slog.Logger, Enforcer *casbin.Enforcer) *Handler {
	return &Handler{
		UserService: user,
		Log:         logger,
		DocsService: production,
		Enforcer:    Enforcer,
	}
}
