// Package config loads the server's settings from environment variables once
// at startup and keeps the API key out of anything that prints the result.
package config

import (
	"errors"
	"fmt"
	"strings"
)

const defaultPort = "8080"

type Config struct {
	SpeechmaticsAPIKey string
	Port               string
}

var ErrMissingAPIKey = errors.New("SPEECHMATICS_API_KEY is required")

// String and GoString keep the API key out of logs and error messages.
// fmt uses String for %v/%s and GoString for %#v.
func (c Config) String() string {
	return fmt.Sprintf("Config{SpeechmaticsAPIKey: [REDACTED], Port: %s}", c.Port)
}

func (c Config) GoString() string {
	return c.String()
}

func Load(getenv func(string) string) (*Config, error) {
	key := strings.TrimSpace(getenv("SPEECHMATICS_API_KEY"))
	if key == "" {
		return nil, ErrMissingAPIKey
	}

	port := strings.TrimSpace(getenv("PORT"))
	if port == "" {
		port = defaultPort
	}

	return &Config{SpeechmaticsAPIKey: key, Port: port}, nil
}
