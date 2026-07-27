package main

import (
	"backend/internal/database"
	"backend/internal/httputil"
	"backend/internal/makes"
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

	makesStore := makes.NewStore(db)
	makes.NewHandler(makesStore).Register(mux)

	addr := ":8080"
	log.Printf("car marketplace backend listening on %s", addr)
	if err := http.ListenAndServe(addr, httputil.WithCORS(mux)); err != nil {
		log.Fatal(err)
	}
}
