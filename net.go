package api

import (
	"net/http"
	"strings"
)

// GetIP convierte el formato 192.168.0.8:34182  a 192.168.0.8
func GetIP(r *http.Request) (ip string) {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded == "" {
		forwarded = r.RemoteAddr
	}

	if strings.Contains(forwarded, "[::1]") || strings.Contains(forwarded, "127.0.0.1") {
		ip = "127.0.0.1"
	} else {
		nip := strings.Split(forwarded, ":") //separamos el contenido
		if len(nip) > 0 {
			ip = nip[0]
		} else {
			ip = forwarded
		}
	}

	return
}
