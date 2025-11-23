package util

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"run/global"
	"run/models"
	"run/models/types"
)

func LoadStandardsFromCSV(path string) ([]models.Standard, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			global.Log.Error(err.Error())
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV records: %w", err)
	}

	if len(records) == 0 {
		return nil, errors.New("CSV file is empty")
	}

	var standards []models.Standard
	for _, row := range records[1:] {
		if len(row) < 5 {
			return nil, fmt.Errorf("invalid CSV row: %v", row)
		}

		eventID := parseUint(row[0])
		gender := types.Gender(parseUint(row[1]))
		level := types.Level(parseUint(row[2]))
		threshold := parseFloat(row[3])
		sys := types.StandardSystem(parseUint(row[4]))

		if eventID == 0 {
			return nil, fmt.Errorf("invalid event_id in row: %v", row)
		}
		if !gender.Valid() || !level.Valid() || !sys.Valid() {
			return nil, fmt.Errorf("invalid enum in row: %v", row)
		}

		std := models.Standard{
			EventID:        eventID,
			Gender:         gender,
			Level:          level,
			Threshold:      threshold,
			StandardSystem: sys,
		}
		standards = append(standards, std)
	}

	return standards, nil
}
