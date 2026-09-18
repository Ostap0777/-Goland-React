package main

import (
	"backend/internal/auth"
	"backend/internal/car"
	"backend/internal/database"
	"backend/internal/httputil"
	"backend/internal/makes"
	"backend/internal/users"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			httputil.WriteError(w, http.StatusServiceUnavailable, "database unavailable")
			return
		}
		httputil.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

// 1. Existing modules
	makesStore := makes.NewStore(db)
	makes.NewHandler(makesStore).Register(mux)

	carRepo := car.NewPostgresRepository(db)
	car.NewHandler(car.NewService(carRepo)).Register(mux)

	// 2. Users module
	userRepo := users.NewPostgresRepository(db)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)
	userHandler.Register(mux)

	// 3. Auth module (використовує userRepo для створення/пошуку користувачів)
	authService := auth.NewService(userRepo)
	authHandler := auth.NewHandler(authService)
	authHandler.RegisterRoutes(mux)
	addr := ":8080"
	log.Printf("car marketplace backend listening on %s", addr)
	if err := http.ListenAndServe(addr, httputil.WithCORS(mux)); err != nil {
		log.Fatal(err)
	}
}
