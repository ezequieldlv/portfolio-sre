package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Definimos la estructura del JSON de respuesta
type SystemStatus struct {
	App       string `json:"app"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

// Funcion que maneja la peticion
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🔔 Peticion recibida en /api/health")

	data := SystemStatus{
		App:       "Ez-Lab Backend",
		Status:    "Prime 🟢",
		Timestamp: time.Now().Format(time.RFC3339),
		Message:   "Zero Trust Architecture is Active.",
	}

	w.Header().Set("Content-Type", "application/json")

	// Permite que cualquier origen consulte
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/api/health", healthHandler)

	port := ":8585"
	fmt.Println("🚀 Servidor Go escuchando en el puerto", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("❌ Error al iniciar", err)
	}
}
