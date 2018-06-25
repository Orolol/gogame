package main

// import (
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func Logger(inner http.Handler, name string) http.Handler {
// 	return http.HandlerFunc(func(c *gin.Context) {
// 		start := time.Now()

// 		inner.ServeHTTP(w, r)

// 		log.Printf(
// 			"%s\t%s\t%s\t%s",
// 			r.Method,
// 			r.RequestURI,
// 			name,
// 			time.Since(start),
// 		)
// 	})
// }
