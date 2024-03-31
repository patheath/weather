package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/patheath/weather/internal/utils"
)

// TODO seperate out as an integration/e2e test not unit test
// as it reaches out to api's

func TestMain(t *testing.T) {
	mainStdout := utils.CaptureStdout(main)
	assert.Equal(t, mainStdout, "Hello, World!\n")
}
