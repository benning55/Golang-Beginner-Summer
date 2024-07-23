package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveData(tracker *IncomeExpenseTracker) error {
	data, err := json.Marshal(tracker)
	if err != nil {
		return fmt.Errorf("error marshaling data: %v", err)
	}

	err = os.WriteFile("transactions.json", data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Println("Data saved successfully!")
	return nil
}

func LoadData(tracker *IncomeExpenseTracker) error {
	data, err := os.ReadFile("transactions.json")
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	err = json.Unmarshal(data, tracker)
	if err != nil {
		return fmt.Errorf("error unmarshaling data: %v", err)
	}

	fmt.Println("Data loaded successfully!")
	return nil
}
