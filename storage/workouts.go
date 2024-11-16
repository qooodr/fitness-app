package storage

import (
	"encoding/json"
	"os"

	"github.com/qooodr/fitness-tracker/models"
)

const filePath = "workouts.json"

func LoadWorkouts() ([]models.Workout, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Workout{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var workouts []models.Workout
	err = json.NewDecoder(file).Decode(&workouts)
	return workouts, err
}

func SaveWorkouts(workouts []models.Workout) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(workouts)
}
