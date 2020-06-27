package main

import (
	"fmt"
	"net"
	"net/http"
)

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARED-FOR")
	// fmt.Println(r.Header)
	if forwarded != "" {
		return forwarded
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	userIP := net.ParseIP(ip)
	fmt.Println(userIP)
	fmt.Println(r.RemoteAddr)
	return r.RemoteAddr
}

func LogMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UserIpAddress = GetIP(r)
		UserRsessionId = "OPERATIONS"

		next.ServeHTTP(w, r)
	})
}
