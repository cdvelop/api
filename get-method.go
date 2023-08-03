package api

import (
	"strings"
)

func getMethodAndObjectFromPath(path string) (string, string) {

	// fmt.Println("Path: ", path)

	if path == "/file" {
		return "file", ""
	}

	parts := strings.Split(path, "/")
	// fmt.Println("Path: ", path, "PARTES: TAMAÑO: ", len(parts), " ", parts)
	if len(parts) < 3 || parts[1] == "" {
		return "", ""
	}
	return parts[1], parts[2]
}
