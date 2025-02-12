package util
import (
	"fmt"
	"os"
	"encoding/json"
)


func WriteJSONToFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print JSON
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}



func AppendJSONToFile(filename string, data interface{}) error {
	var existingData []interface{}

	// Check if file exists
	fileBytes, err := os.ReadFile(filename)
	if err == nil && len(fileBytes) > 0 {
		// Try to unmarshal as an array
		err = json.Unmarshal(fileBytes, &existingData)
		if err != nil {
			// If it's an object instead of an array, wrap it in an array
			var singleObject map[string]interface{}
			if json.Unmarshal(fileBytes, &singleObject) == nil {
				existingData = append(existingData, singleObject)
			} else {
				return fmt.Errorf("failed to decode existing JSON: %w", err)
			}
		}
	}

	existingData = append(existingData, data)

	// Write updated data
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print JSON
	if err := encoder.Encode(existingData); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}