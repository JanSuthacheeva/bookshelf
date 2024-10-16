package main

import "net/http"

func commonHeaders(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // TODO: Inform about these headers and what they do.
    w.Header().Set("Content-Security-Policy",
      "default-src 'self'; style src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
    w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
    w.Header().Set("X-Content-Type-Options", "nosniff")
    w.Header().Set("X-Frame-Options", "deny")
    w.Header().Set("X-XSS-Protection", "0")

    w.Header().Set("Server", "Go")

    next.ServeHTTP(w, r)
  })
}
