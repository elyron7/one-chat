package main

import (
	"database/sql"
	user_controller "go-chat/internal/user/controller"
	"go-chat/pkg/zaplog"

	_ "github.com/lib/pq" // Postgres driver

	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	zapLogger := &zaplog.Zaplog{}
	zapLogger.Init()
	defer zapLogger.Logger.Sync() // 确保日志缓冲区在程序退出前刷新

	dsn := "host=localhost port=5432 user=your_username password=your_password dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		// some log
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		// some log
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userCtrl := user_controller.NewUserController(db, zapLogger)
	userCtrl.SetUpRoutes(r)

	r.Run(":8080")
}
