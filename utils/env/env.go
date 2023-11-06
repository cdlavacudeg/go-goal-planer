package env

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/joho/godotenv"
)

type Config struct {
	Env       string
	DynConfig DynamoDBConfig `required:"true"`
	AwsConfig AwsConfig      `required:"true"`
}

type DynamoDBConfig struct {
	Endpoint string `required:"true"`
}

type AwsConfig struct {
	Region string `required:"true"`
}

var config *Config

// checkRequiredEnv verifies that the required environment variables are set

func checkRequiredEnv(config *Config) error {
	// Get the reflect.Value of the Config struct and dereference it to access its fields
	val := reflect.ValueOf(config).Elem()

	// Get the reflect.Type of the Config struct
	typ := val.Type()

	// Loop through all the fields of the Config struct
	for i := 0; i < val.NumField(); i++ {
		// Get the reflect.Value and reflect.Type of the current field
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Check if the field has a "required" tag and if it is set to "true"
		if fieldType.Tag.Get("required") == "true" {
			// If the field is a zero value (empty, uninitialized), it is missing in the environment
			if field.IsZero() {
				// Create an error with the message indicating the missing environment variable
				return errors.New("missing environment variable: " + fieldType.Name)
			}
			// If the field is a nested structure with fields of its own
			if field.NumField() > 0 {
				// Loop through all the fields of the nested structure
				for j := 0; j < field.NumField(); j++ {
					// Get the reflect.Value and reflect.Type of the nested field
					subField := field.Field(j)
					subFieldType := field.Type().Field(j)

					// Check if the nested field has a "required" tag and if it is set to "true"
					if subFieldType.Tag.Get("required") == "true" && subField.IsZero() {
						// Create an error with the message indicating the missing environment variable in the nested structure
						errMsg := fmt.Sprintf(
							"missing environment variable: %v -> %v\n",
							fieldType.Name,
							subFieldType.Name,
						)
						return errors.New(errMsg)
					}
				}
			}
		}
	}

	// If all required fields are present and have valid values, return nil (no error)
	return nil
}

func LoadConfig() (*Config, error) {
	if config != nil {
		// If the config is already loaded, return it
		return config, nil
	}

	if err := godotenv.Load(filepath.Join("utils", "env", ".env")); err != nil {
		return nil, err
	}

	config = &Config{
		Env: os.Getenv("GO_ENV"),
		DynConfig: DynamoDBConfig{
			Endpoint: os.Getenv("DYNAMO_ENDPOINT"),
		},
		AwsConfig: AwsConfig{
			Region: os.Getenv("AWS_REGION"),
		},
	}

	if err := checkRequiredEnv(config); err != nil {
		return nil, err
	}

	return config, nil
}

func GetConfig() (*Config, error) {
	if config == nil {
		_, err := LoadConfig()
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
