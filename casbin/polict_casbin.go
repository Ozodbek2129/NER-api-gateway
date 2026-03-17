package casbin

import (
	"fmt"
	"log/slog"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

const (
	host     = "localhost"
	port     = "5432"
	dbname   = "casbin"
	username = "postgres"
	password = "salom"
)

func CasbinEnforcer(logger *slog.Logger) (*casbin.Enforcer, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, username, dbname, password)

	// adapter yaratish
	adapter, err := xormadapter.NewAdapter("postgres", connStr)
	if err != nil {
		logger.Error("Error creating Casbin adapter", "error", err.Error())
		return nil, err
	}

	// enforcer yaratish
	enforcer, err := casbin.NewEnforcer("casbin/model.conf", adapter)
	if err != nil {
		logger.Error("Error creating Casbin enforcer", "error", err.Error())
		return nil, err
	}

	if err := enforcer.LoadPolicy(); err != nil {
		logger.Error("Error loading Casbin policy", "error", err.Error())
		return nil, err
	}

	// policies
	policies := [][]string{
		// super
		{"super", "/api/user/getbyuser/:email", "GET"},
		{"super", "/api/user/update_password", "PUT"},
		{"super", "/api/user/delete_user/:id", "DELETE"},
		{"super", "/api/user/update_role", "PUT"},
		{"super", "/api/user/image_update", "PUT"},
		{"super", "/api/user/all_users/:limit/:page", "GET"},
		{"super", "/api/user/logout", "POST"},

		{"super", "/api/contract/newcontract", "POST"},
		{"super", "/api/contract/contract_update", "PUT"},
		{"super", "/api/contract/contract_delete/:id", "DELETE"},
		{"super", "/api/contract/get_name/:name", "GET"},
		{"super", "/api/contract/all_contract/:limit/:page", "GET"},

		{"super", "/api/contract/inside_contract", "POST"},
		{"super", "/api/contract/insidecontract_update", "PUT"},
		{"super", "/api/contract/insidecontract_delete/:id", "DELETE"},
		{"super", "/api/contract/all_insidecontract/:limit/:page", "GET"},

		// admin
		{"admin", "/api/user/update_password", "PUT"},
		{"admin", "/api/user/image_update", "PUT"},
		{"admin", "/api/user/logout", "POST"},

		// user
		{"user", "/api/user/update_password", "PUT"},
		{"user", "/api/user/logout", "POST"},
	}

	// duplicate qo‘shilmasligi uchun
	for _, p := range policies {
		exists, _ := enforcer.HasPolicy(p)
		if !exists {
			if _, err := enforcer.AddPolicy(p); err != nil {
				logger.Error("Error adding policy", "policy", p, "error", err.Error())
			}
		}
	}

	// policy saqlash
	if err := enforcer.SavePolicy(); err != nil {
		logger.Error("Error saving Casbin policy", "error", err.Error())
		return nil, err
	}

	return enforcer, nil
}