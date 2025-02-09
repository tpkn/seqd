package seqd

import (
	"os"
)

// StdoutWriter writer for production
type StdoutWriter struct{}

// WriteString writes new-line separated results to STDOUT
func (o StdoutWriter) WriteString(s string) (int, error) {
	return os.Stdout.WriteString(s + "\n")
}
