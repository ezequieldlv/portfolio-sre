package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// =====================================================================
// 1. ESTRUCTURAS DE DATOS
// =====================================================================
type SystemStats struct {
	Status       string `json:"status"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
	GoVersion    string `json:"go_version"`
	ServerTime   string `json:"server_time"`
	Message      string `json:"message"`
}

// =====================================================================
// 2. MIDDLEWARES (Filtros e Interceptores)
// =====================================================================
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		allowedOrigins := map[string]bool{
			"https://ez-lab.site":           true,
			"https://www.ez-lab.site":       true,
			"https://ezequieldlv.github.io": true,
			"http://localhost:1313":         true, // Desarrollo local
		}

		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// =====================================================================
// 3. CONTROLADORES (Handlers de Negocio)
// =====================================================================

// healthCheck: Responde a las peticiones del frontend (Hugo)
func healthCheck(w http.ResponseWriter, r *http.Request) {
	stats := SystemStats{
		Status:       "Prime 🟢",
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		GoVersion:    runtime.Version(),
		ServerTime:   time.Now().Format(time.Kitchen),
		Message:      "Live from SRE Microservice",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

// chaosHandler: Simula una falla crítica (OOM o Fatal Panic) para probar resiliencia
func chaosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("⚠️ ALERTA: Iniciando secuencia de Chaos Engineering. Matando proceso...")
	// Enviamos un código 1 al OS, simulando que el programa crasheó por un error grave
	os.Exit(1) 
}

// =====================================================================
// 4. MOTOR PRINCIPAL (Router y Servidor)
// =====================================================================
func main() {
	mux := http.NewServeMux()

	// RUTAS PÚBLICAS (Con protección CORS para el Frontend)
	mux.HandleFunc("/api/health", enableCORS(healthCheck))
	
	// RUTAS DE INGENIERÍA (Sin CORS, solo para uso interno del SRE)
	mux.HandleFunc("/api/chaos/crash", enableCORS(chaosHandler)) // Botón de autodestrucción
	
	// RUTA DE OBSERVABILIDAD (Prometheus Scrape Target)
	// Manejador automático: expone memoria RAM, hilos de Go, y recolector de basura
	mux.Handle("/metrics", promhttp.Handler())

	log.Println("🚀 Backend SRE iniciado en el puerto 8585...")
	log.Println("📊 Endpoint de métricas disponible en /metrics")
	log.Println("💀 Endpoint de Chaos disponible en /api/chaos/crash")

	err := http.ListenAndServe(":8585", mux)
	if err != nil {
		log.Fatal("Error crítico iniciando el servidor: ", err)
	}
}