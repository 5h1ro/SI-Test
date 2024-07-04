package app

import (
	"customer/config"
	"customer/internal/entity"
	"customer/internal/routes"
	"customer/internal/types"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Config(path string) *config.Config {
	// Configuration
	cfg, err := config.NewConfig(path)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	return cfg
}

func GenerateAuthToken() *jwtauth.JWTAuth {
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SIGN_KEY")), nil)
	return tokenAuth
}

func Database(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Database, cfg.DB.Port)
	var log logger.Interface

	if cfg.App.LogLevel == "production" {
		log = nil
	} else {
		log = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		fmt.Printf("app - Run - postgres.New: %v", err)
	}
	db.AutoMigrate(&entity.Customer{})

	return db
}

func Server(cfg *config.Config) *types.Server {
	db := Database(cfg)
	app := *chi.NewRouter()
	AuthToken := GenerateAuthToken()
	routes.ApiRouter(&app, db, cfg, AuthToken)
	return &types.Server{
		Router:    &app,
		DB:        db,
		AuthToken: AuthToken,
	}
}

func Run() *chi.Mux {
	cfg := Config(".env")
	log.Println("Starting with port " + cfg.HTTP.Port)

	server := Server(cfg)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.HTTP.Port), server.Router)

	return server.Router
}
