package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qooodr/fitness-tracker/models"
	"github.com/qooodr/fitness-tracker/storage"
)

func GetWorkoutsHandler(w http.ResponseWriter, r *http.Request) {
	workouts, err := storage.LoadWorkouts()
	if err != nil {
		http.Error(w, "Failed to load workouts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workouts)
}

func AddWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	var newWorkout models.Workout
	err := json.NewDecoder(r.Body).Decode(&newWorkout)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	workouts, err := storage.LoadWorkouts()
	if err != nil {
		http.Error(w, "Failed to load workouts", http.StatusInternalServerError)
		return
	}

	workouts = append(workouts, newWorkout)
	err = storage.SaveWorkouts(workouts)
	if err != nil {
		http.Error(w, "Failed to save workouts", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Workout added successfully")
}
