package api

import (
	"github.com/gin-gonic/gin"
	. "github.com/hiepphatle1104/matcha_latte_store/modules/account/service"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/storage"
	"github.com/hiepphatle1104/matcha_latte_store/modules/middleware"
	"gorm.io/gorm"
	"net/http"
)

type accountHandler struct {
	service IAccountService
}

func newAccountHandler(db *gorm.DB) *accountHandler {
	repo := storage.NewAccountStorage(db)
	service := NewAccountService(repo)
	return &accountHandler{service}
}

func SetupAccountModule(db *gorm.DB) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	handler := newAccountHandler(db)
	router := e.Group("/api/v1")

	router.POST("/register", handler.CreateAccount)
	router.POST("/login", handler.GetToken)

	router.Use(middleware.JWTAuth)
	router.GET("/logout", handler.DeleteToken)
	router.GET("/me", handler.GetAccount)

	return e
}
