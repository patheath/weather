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

	// Will fail is API is down but that is by design
	n := strings.Count(mainStdout, "\n")
	assert.Equal(t, n, 2*model.HOURS)

	n = strings.Count(mainStdout, "The hour is")
	assert.Equal(t, n, 2*model.HOURS)
}
