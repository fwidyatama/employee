package server

import (
	"database/sql"
	"employee/internal/config"
	"github.com/labstack/echo/v4"
)

type Rest struct {
	Echo   *echo.Echo
	Config *config.Config
	SQL    *sql.DB
}

func NewServer(cfg *config.Config, sql *sql.DB) *Rest {
	return &Rest{
		Echo:   echo.New(),
		Config: cfg,
		SQL:    sql,
	}
}

func (r *Rest) Start(addr string) error {
	return r.Echo.Start(":" + addr)
}
