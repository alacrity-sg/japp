package main

import (
	"JApp-Backend/api"
	"JApp-Backend/middleware"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/placeholder", api.PlaceholderApi)
	// Add your routes as needed
	port := os.Getenv("PORT")
	address := fmt.Sprintf("0.0.0.0:%s", port)
	//srv := &http.Server{
	//	Addr:         address,
	//	WriteTimeout: time.Second * 15,
	//	ReadTimeout:  time.Second * 15,
	//	IdleTimeout:  time.Second * 60,
	//	Handler:      r,
	//}
	log.Println(port)
	r.Use(middleware.JsonMiddleware)
	r.Use(middleware.LoggingMiddleware)
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	log.Printf("Server running on at %s", address)
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil {
	//		log.Println(err)
	//	}
	//}()
	//
	//log.Printf("Server running on at %s", address)
	//
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//
	//<-c
	//
	//ctx, cancel := context.WithTimeout(context.Background(), wait)
	//defer cancel()
	//srv.Shutdown(ctx)
	//// Optionally, you could run srv.Shutdown in a goroutine and block on
	//// <-ctx.Done() if your application should wait for other services
	//// to finalize based on context cancellation.
	//log.Println("Shutting down application server")
	//os.Exit(0)
}
