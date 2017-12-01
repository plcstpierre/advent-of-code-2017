package commons

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var buffer []string

func TestReadFile(t *testing.T) {
	ReadFile("../../input/test/test.txt", appendAndTest)

	assert.Equal(t, len(buffer), 3)

	assert.Equal(t, "a", buffer[0])
	assert.Equal(t, "b", buffer[1])
	assert.Equal(t, "d", buffer[2])
}

func appendAndTest(line string) {
	buffer = append(buffer, line)
}
