package main

import (
	"strings"
	"testing"

	"github.com/patheath/weather/internal/model"
	"github.com/patheath/weather/internal/utils"
	"github.com/stretchr/testify/assert"
)

// Only runs when short option is not specified
func TestMainIntegration(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping integration test")
	}

	mainStdout := utils.CaptureStdout(main)

	// Will fail if API is down but that is by design
	num_providers := 2
	n := strings.Count(mainStdout, "The hour is")
	assert.Equal(t, n, num_providers*model.HOURS)
}
