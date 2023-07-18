package api

import "strings"

func getMethodAndObjectFromPath(path string) (string, string) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 || parts[1] == "" {
		return "", ""
	}
	return parts[1], parts[2]
}
