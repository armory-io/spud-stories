package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// create N maps with byte slices of M
		log.Printf("creating %d slices of size:%d with deviation: %d", numObjects, objectSize, deviation)
		h := map[int][]byte{}
		for i := 0; i < numObjects; i++ {
			h[i] = make([]byte, objectSize+rand.Intn(deviation))
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"message": "ok"}`))
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
