package service

import (
	"github.com/dinorain/checkoutaja/config"
	"github.com/dinorain/checkoutaja/internal/session"
	"github.com/dinorain/checkoutaja/internal/user"
	"github.com/dinorain/checkoutaja/pkg/logger"
)

type usersServiceGRPC struct {
	logger logger.Logger
	cfg    *config.Config
	userUC user.UserUseCase
	sessUC session.SessUseCase
}

// Auth service constructor
func NewAuthServerGRPC(logger logger.Logger, cfg *config.Config, userUC user.UserUseCase, sessUC session.SessUseCase) *usersServiceGRPC {
	return &usersServiceGRPC{logger: logger, cfg: cfg, userUC: userUC, sessUC: sessUC}
}
