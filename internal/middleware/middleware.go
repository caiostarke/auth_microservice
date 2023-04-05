package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/casbin/casbin/persist"
	"github.com/gin-gonic/gin"
)

func Authorize(obj string, act string, adapter persist.Adapter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ok, err := enforce("joao", obj, act, adapter)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(500, "error occurred when authorizing user")
			return
		}
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, "forbidden")
			return
		}
		ctx.Next()
	}
}

func enforce(sub, obj, act string, adapter persist.Adapter) (bool, error) {
	enforcer := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	err := enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok := enforcer.Enforce(sub, obj, act)
	return ok, nil
}
