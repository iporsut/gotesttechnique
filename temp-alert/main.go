package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	server := NewServer()
	http.ListenAndServe(":8080", server)
}

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("POST /temp", func(w http.ResponseWriter, r *http.Request) {
		type TempRequest struct {
			SensorID    string  `json:"sensor_id"`
			Temperature float64 `json:"temperature_celsius"`
		}

		var tempReq TempRequest
		err := json.NewDecoder(r.Body).Decode(&tempReq)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		sqlDB, err := sql.Open("postgres", "user=temp password=temp dbname=temp sslmode=disable")
		if err != nil {
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer sqlDB.Close()

		_, err = sqlDB.Exec("INSERT INTO temperatures (sensor_id, temperature_celsius) VALUES ($1, $2)", tempReq.SensorID, tempReq.Temperature)
		if err != nil {
			http.Error(w, "Failed to save temperature", http.StatusInternalServerError)
			return
		}

		if tempReq.Temperature > 30.0 {
			err := SendSMSAlert(tempReq.SensorID, tempReq.Temperature)
			if err != nil {
				http.Error(w, "Failed to send SMS alert", http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"status":"temperature saved"}`))
	})
	return mux
}

func SendSMSAlert(sensorID string, temperature float64) error {
	// Placeholder for SMS sending logic
	// In a real implementation, integrate with an SMS service provider
	return nil
}
