package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/joeleonardo/golang_class_bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
	default:
		t.Error(fmt.Sprintf("type %s is not *chi.Mux", v))
	}
}
