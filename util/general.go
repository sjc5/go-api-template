package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sjc5/go-api-template/global"
)

func UnmarshalAndValidateFromRequest(r *http.Request, destination interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(destination); err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}
	if err := global.Validate.Struct(destination); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	return nil
}

func UnmarshalAndValidateFromString(data string, destination interface{}) error {
	if err := json.Unmarshal([]byte(data), destination); err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}
	if err := global.Validate.Struct(destination); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	return nil
}
