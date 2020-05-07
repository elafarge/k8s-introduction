package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
)

func sayHello(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	personToGreet := os.Getenv("PERSON_TO_GREET")

	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("Hello " + personToGreet))
	if err != nil {
		log.Println(err)
	}
}

func isPrime(test int) bool {
	for i := 2; i <= int(math.Sqrt(float64(test))); i++ {
		if test%i == 0 {
			return false
		}
	}
	return true
}

func countPrimesUntil(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	stringTarget := params.ByName("target")
	target, err := strconv.Atoi(stringTarget)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(fmt.Sprintf("could not parse %s to int64", stringTarget)))
		return
	}

	primeCount := 1 // 1 is prime, we all know that
	for i := 2; i <= target; i++ {
		if isPrime(i) {
			primeCount++
		}
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(fmt.Sprintf("found %d prime numbers between 1 and %d", primeCount, target)))
}

func main() {
	router := httprouter.New()

	// our routes
	router.GET("/hello", sayHello)
	router.GET("/countprimesuntil/:target", countPrimesUntil)

	// a little health check
	router.GET("/healthz", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, I'm healthy AF. !"))
	})

	// add a logging middleware
	requestID := 0
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID++
		startTime := time.Now()
		log.Printf("received request %d: %s %s", requestID, r.Method, r.URL.Path)
		router.ServeHTTP(w, r)
		log.Printf("request %d took %s", requestID, time.Since(startTime))
	})

	server := http.Server{
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		Handler:      handler,
	}

	// catch SIGINT/SIGTERM and gracefully close the server in the background
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("received SIGINT/SIGTERM, waiting 6 seconds for health check to fail")
		time.Sleep(6 * time.Second)

		log.Println("processing pending requests and exiting")
		if err := server.Close(); err != nil {
			log.Fatalf("error closing server: %v", err)
		}
	}()

	log.Println("starting server...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
