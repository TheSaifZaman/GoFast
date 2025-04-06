package helpers

import "strings"

func ToPascalCase(s string) string {
	parts := strings.Split(s, "_")
	var pascal string
	for _, p := range parts {
		if len(p) == 0 {
			continue
		}
		pascal += strings.ToUpper(string(p[0])) + p[1:]
	}
	return pascal
}

func ToSnakeCase(s string) string {
	var snake string
	for i, c := range s {
		if i > 0 && isUpper(c) {
			snake += "_"
		}
		snake += strings.ToLower(string(c))
	}
	return snake
}

func ToPlural(s string) string {
	parts := strings.Split(s, "_")
	if len(parts) == 0 {
		return s
	}
	lastPart := parts[len(parts)-1]
	if lastPart[len(lastPart)-1] == 's' {
		return s
	}
	if lastPart[len(lastPart)-1] == 'y' {
		return strings.Join(parts[:len(parts)-1], "_") + "_" + lastPart[:len(lastPart)-1] + "ies"
	}
	return s + "s"
}
