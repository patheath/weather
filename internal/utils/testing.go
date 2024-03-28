package utils

import (
	"bytes"
	"io"
	"os"
)

// Use this to capture the cli output to check for tests
// For this toy app this is fine for a large app this will have
// issues: buffer could overflow, not concurrent, etc.
// Source https://gist.github.com/mindscratch/0faa78bd3c0005d080bf
func CaptureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil{
		println(err)
	}
	return buf.String()
}
