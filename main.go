package main

import (
	"fmt"
	"net/http"

	"github.com/qooodr/fitness-tracker/handlers"
)

func main() {
	http.HandleFunc("/workouts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetWorkoutsHandler(w, r)
		case http.MethodPost:
			handlers.AddWorkoutHandler(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
