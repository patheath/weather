package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/patheath/weather/internal/utils"
)

func TestMain(t *testing.T) {
	mainStdout := utils.CaptureStdout(main)
	assert.Equal(t, mainStdout, "Hello, World!\n")
}
