package envparser

import (
	"bufio"
	"os"
	"strings"

	"github.com/crewcrew23/envguard/internal/envtypes"
)

func ParseFile(filePath string) (*envtypes.EnvMap, error) {
	envMap := envtypes.NewEnvMap()

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, val := parseVar(line)
		if key != "" {
			envMap.Set(key, val)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return envMap, nil
}

func parseVar(line string) (key, value string) {
	line = strings.TrimPrefix(line, "export ")
	parts := strings.SplitN(line, "=", 2)

	if len(parts) < 2 {
		return strings.TrimSpace(parts[0]), ""
	}

	key = strings.TrimSpace(parts[0])
	value = strings.TrimSpace(parts[1])
	value = strings.Trim(value, `"`)

	return key, value
}
