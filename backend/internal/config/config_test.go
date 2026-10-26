package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func fakeEnv(m map[string]string) func(string) string {
	return func(k string) string { return m[k] }
}

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		want    *Config
		wantErr error
	}{
		{
			name: "valid config",
			env: map[string]string{
				"SPEECHMATICS_API_KEY": "abc123",
				"PORT":                 "8080",
			},
			want: &Config{
				SpeechmaticsAPIKey: "abc123",
				Port:               "8080",
			},
			wantErr: nil,
		},
		{
			name:    "missing api key",
			env:     map[string]string{"PORT": "8080"},
			want:    nil,
			wantErr: ErrMissingAPIKey,
		},
		{
			name: "missing port",
			env:  map[string]string{"SPEECHMATICS_API_KEY": "abc123"},
			want: &Config{
				SpeechmaticsAPIKey: "abc123",
				Port:               "8080",
			},
			wantErr: nil,
		},
		{
			name:    "blank api key",
			env:     map[string]string{"SPEECHMATICS_API_KEY": "", "PORT": "8080"},
			want:    nil,
			wantErr: ErrMissingAPIKey,
		},
		{
			name:    "whitespace api key",
			env:     map[string]string{"SPEECHMATICS_API_KEY": "       ", "PORT": "8080"},
			want:    nil,
			wantErr: ErrMissingAPIKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Load(fakeEnv(tt.env))

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestKeyRedacted(t *testing.T) {
	const canaryKey = "testkey-3fa"
	cfg := Config{canaryKey, "8080"}

	formats := []string{"%v", "%+v", "%s", "%#v"}
	values := map[string]any{
		"value":   cfg,
		"pointer": &cfg,
	}

	for form, v := range values {
		for _, f := range formats {
			t.Run(form+" "+f, func(t *testing.T) {
				out := fmt.Sprintf(f, v)

				if strings.Contains(out, canaryKey) {
					t.Fatalf("%s leaked secret: %s", f, out)
				}

				if !strings.Contains(out, "[REDACTED]") {
					t.Errorf("%s missing redaction marker: %s", f, out)
				}
			})
		}
	}
}
