package clientip

import (
	"net/http"
	"strings"
)

func ProcessClientIp(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	} else {
		ip = strings.Split(ip, ",")[0]
	}
	ip = strings.Split(ip, ":")[0]

	return ip
}
