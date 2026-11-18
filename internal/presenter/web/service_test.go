package web

import (
	"os"
	"testing"
)

func TestGetAllowOriginsInReleaseMode(t *testing.T) {
	oldOrigin, existed := os.LookupEnv("CORS_ORIGIN")
	oldMode, modeExisted := os.LookupEnv("GIN_MODE")

	defer func() {
		if existed {
			os.Setenv("CORS_ORIGIN", oldOrigin)
		} else {
			os.Unsetenv("CORS_ORIGIN")
		}

		if modeExisted {
			os.Setenv("GIN_MODE", oldMode)
		} else {
			os.Unsetenv("GIN_MODE")
		}
	}()

	os.Setenv("GIN_MODE", "release")
	os.Setenv("CORS_ORIGIN", "https://arcade.tleung.net")

	got := getAllowOrigins()

	if len(got) != 1 || got[0] != "https://arcade.tleung.net" {
		t.Fatalf("getAllowOrigins() = %v, want [https://arcade.tleung.net]", got)
	}
}
