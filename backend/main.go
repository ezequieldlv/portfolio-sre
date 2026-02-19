package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"time"
)

// 1. Estructura de Datos (El Molde del JSON)
type SystemStats struct {
	Status       string `json:"status"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
	GoVersion    string `json:"go_version"`
	ServerTime   string `json:"server_time"`
	Message      string `json:"message"`
}

// 2. Middleware de Seguridad (CORS)
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Lista blanca de orígenes permitidos
		origin := r.Header.Get("Origin")
		allowedOrigins := map[string]bool{
			"https://ez-lab.site":           true,
			"https://www.ez-lab.site":       true,
			"https://ezequieldlv.github.io": true,
			"http://localhost:1313":         true, // Para desarrollo local
		}

		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// 3. El Controlador de Telemetría (El Endpoint)
func healthCheck(w http.ResponseWriter, r *http.Request) {
	// Recopilamos datos reales del contenedor/Raspberry
	stats := SystemStats{
		Status:       "Prime 🟢",
		OS:           runtime.GOOS,                    // ej: linux
		Architecture: runtime.GOARCH,                  // ej: arm64
		GoVersion:    runtime.Version(),               // ej: go1.22
		ServerTime:   time.Now().Format(time.Kitchen), // Hora actual del server
		Message:      "Live from Raspberry Pi 5 Cluster",
	}

	// Le decimos al cliente que le enviamos un JSON
	w.Header().Set("Content-Type", "application/json")

	// Convertimos la estructura a JSON y la enviamos
	json.NewEncoder(w).Encode(stats)
}

// 4. El Motor Principal
func main() {
	// Creamos el enrutador
	mux := http.NewServeMux()

	// Asignamos la ruta "/api/health" a nuestra función, pasándola por el filtro CORS
	mux.HandleFunc("/api/health", enableCORS(healthCheck))

	log.Println("🚀 Backend SRE iniciado en el puerto 8585...")

	// Arrancamos el servidor
	err := http.ListenAndServe(":8585", mux)
	if err != nil {
		log.Fatal("Error crítico iniciando el servidor: ", err)
	}
}
