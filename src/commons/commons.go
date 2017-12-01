package commons

import (
	"os"
	"bufio"
	"strings"
)

func ReadFile(path string, callback func(string)) {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		raw := scanner.Text()
		line := strings.TrimSpace(raw)

		if len(line) == 0 {
			continue
		}

		callback(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
