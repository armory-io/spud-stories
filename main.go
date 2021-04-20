package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var version = "dev"

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// create N maps with byte slices of M
		dev := rand.Intn(deviation) + objectSize
		log.Printf("creating %d slices of standard size:%d with total size deviation: %d", numObjects, objectSize, dev)
		h := map[int][]byte{}
		for i := 0; i < numObjects; i++ {
			h[i] = make([]byte, dev)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"message": "ok"}`))
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"status": "healthy"}`)
	})

	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"version": "%s"}`, version)
	})
}

var (
	addr                              string
	numObjects, objectSize, deviation int
)

func main() {

	flag.StringVar(&addr, "addr", ":3000", "server address")
	flag.IntVar(&numObjects, "num-objects", 1000, "number of objects")
	flag.IntVar(&deviation, "deviation", 1024, "size by which to deviate the object size")
	flag.IntVar(&objectSize, "object-size", 1024, "fixed size for each object")
	flag.Parse()

	mux := http.NewServeMux()

	RegisterHandlers(mux)

	s := &http.Server{Addr: addr, Handler: mux}

	log.Printf("starting spud-stories server on %s. enjoy your 'taters.", addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("server closed unexpectedly: %s", err.Error())
	}
	log.Println("spud-stories server shut down gracefully.")
}
