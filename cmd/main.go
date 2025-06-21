package main

import (
	"github.com/hiepphatle1104/matcha_latte_store/config"
	. "github.com/hiepphatle1104/matcha_latte_store/modules/account/routes/api"
	"github.com/hiepphatle1104/matcha_latte_store/pkg/logger"
	"golang.org/x/sync/errgroup"
	"net/http"
)

var g errgroup.Group

func main() {
	logger.InitLogger()
	logger.Info("Matcha Latte Store is starting...")

	cfg, err := config.LoadServerConfig()
	if err != nil {
		logger.Fatal("Failed to load server config: " + err.Error())
	}
	logger.Info("Server configuration loaded.")

	accountDB, err := config.LoadAccountStorage(cfg)
	if err != nil {
		logger.Fatal("Failed to connect to account storage: " + err.Error())
	}
	logger.Info("Connected to account storage.")

	accountModule := &http.Server{
		Addr:    ":3000",
		Handler: SetupAccountModule(accountDB),
	}

	g.Go(func() error {
		return accountModule.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		logger.Fatal(err.Error())
	}
}
