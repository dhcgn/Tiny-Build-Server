package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	listenAddrPtr := flag.String("port", "5000", "The port which the build server should listen on")
	listenAddr := fmt.Sprintf(":%v", *listenAddrPtr)
	fmt.Printf("Server is ready to handle requests at port %v\n", *listenAddrPtr)

	if _, err := loadSysConfig(); err != nil {
		log.Fatal("could not handle config/app.yaml file; something went wrong")
	}

	router := mux.NewRouter()
	router.HandleFunc("/bitbucket-receive", bitBucketReceiveHandler).Methods("POST")
	router.HandleFunc("/github-receive", gitHubReceiveHandler).Methods("POST")
	router.HandleFunc("/gitlab-receive", gitLabReceiveHandler).Methods("POST")
	router.HandleFunc("/gitea-receive", giteaReceiveHandler).Methods("POST")
	router.HandleFunc("/ping", pingHandler)

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      limit(router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		fmt.Println("\nServer is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("Could not gracefully shutdown the server: %v\n", err)
			os.Exit(-1)
		}
		close(done)
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Could not listen on %s: %v\n", listenAddr, err)
		os.Exit(-1)
	}
	<-done
	fmt.Println("Server shutdown complete. Have a nice day!")
}

