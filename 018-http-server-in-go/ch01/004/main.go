package main

import (
	"database/sql"
	"learn/internal/database"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform       string
	jwtSecret      string
	polkaKey       string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	jwtSecret := os.Getenv("JWT_SECRET")
	polkaKey := os.Getenv("POLKA_KEY")

	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	apiConfig := &apiConfig{fileserverHits: atomic.Int32{}, db: dbQueries, platform: platform, jwtSecret: jwtSecret, polkaKey: polkaKey}

	const FILE_PATH_ROOT = "."
	const PORT = "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", apiConfig.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(FILE_PATH_ROOT)))))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/chirps", apiConfig.handleChiprsCreate)
	mux.HandleFunc("GET /api/chirps", apiConfig.handleChirpsGet)
	mux.HandleFunc("/api/chirps/{chirpID}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			apiConfig.handleChirpsGetOne(w, r)
		case http.MethodDelete:
			apiConfig.handleChirpsDelete(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("POST /api/users", apiConfig.handleUsersCreate)
	mux.HandleFunc("POST /api/login", apiConfig.handlerLogin)
	mux.HandleFunc("POST /api/refresh", apiConfig.handlerRefreshToken)
	mux.HandleFunc("POST /api/revoke", apiConfig.handlerRevoke)
	mux.HandleFunc("PUT /api/users", apiConfig.handleUsersUpdate)

	mux.HandleFunc("GET /admin/metrics", apiConfig.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiConfig.handlerReset)

	mux.HandleFunc("POST /api/polka/webhooks", apiConfig.handleWebhook)

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", FILE_PATH_ROOT, PORT)
	log.Fatal(server.ListenAndServe())
}
