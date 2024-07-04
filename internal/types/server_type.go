package types

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
)

type Server struct {
	Router    *chi.Mux
	DB        *gorm.DB
	AuthToken *jwtauth.JWTAuth
}
