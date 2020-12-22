package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type SpudStoriesAPIConfig struct {
	Objects int `json:"objects" yaml:"objects"`
	Size    int `json:"size" yaml:"size"`
}

func (s *SpudStoriesAPIConfig) EnsureDefaults() {
	if s.Objects == 0 {
		s.Objects = 1000
	}
	if s.Size == 0 {
		s.Size = 1024
	}
}

type SpudStoriesAPI struct {
	objectSize int
	numObjects int
}

func NewSpudStoriesAPI(objectSize, numObjects int) (*SpudStoriesAPI, error) {
	return &SpudStoriesAPI{
		objectSize: objectSize,
		numObjects: numObjects,
	}, nil
}

func parseConfig(configPath string) (*SpudStoriesAPIConfig, error) {
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var config SpudStoriesAPIConfig
	if err := yaml.NewDecoder(bytes.NewReader(b)).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (s *SpudStoriesAPI) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// create N maps with byte slices of M
		log.Printf("creating %d slices of size %d", s.numObjects, s.objectSize)
		h := map[int][]byte{}
		for i := 0; i < s.numObjects; i++ {
			h[i] = make([]byte, s.objectSize)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"message": "ok"}`))
	})
}

func main() {
	var (
		addr                   string
		numObjects, objectSize int
	)
	flag.StringVar(&addr, "addr", ":3000", "server address")
	flag.IntVar(&numObjects, "num-objects", 1000, "number of objects")
	flag.IntVar(&objectSize, "object-size", 1024, "fixed size for each object")
	flag.Parse()

	mux := http.NewServeMux()

	api, err := NewSpudStoriesAPI(objectSize, numObjects)
	if err != nil {
		log.Fatalf("failed to create new spudstories server: %s", err.Error())
	}

	api.RegisterHandlers(mux)

	s := &http.Server{Addr: addr, Handler: mux}

	log.Printf("starting spud-stories server on %s. enjoy your 'taters.", addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("server closed unexpectedly: %s", err.Error())
	}
	log.Println("spud-stories server shut down gracefully.")
}
