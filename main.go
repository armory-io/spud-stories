package main

import (
	"flag"
	"log"
	"net/http"
)

type SpudStoriesAPIConfig struct {
	Objects int `json:"objects" yaml:"objects"`
	Size    int `json:"size" yaml:"size"`
}

type SpudStoriesAPI struct {
	config SpudStoriesAPIConfig
}

func NewSpudStoriesAPI(configPath string) (*SpudStoriesAPI, error) {
	// TODO - make configuration dynamic
	return &SpudStoriesAPI{
		config: SpudStoriesAPIConfig{
			Objects: 1000,
			Size:    1000,
		},
	}, nil
}

func (s *SpudStoriesAPI) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// create N maps with byte slices of M
		log.Printf("creating %d slices of size %d", s.config.Objects, s.config.Size)
		h := map[int][]byte{}
		for i := 0; i < s.config.Objects; i++ {
			h[i] = make([]byte, s.config.Size)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"message": "ok"}`))
	})
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":3000", "server address")
	flag.Parse()

	mux := http.NewServeMux()

	api := &SpudStoriesAPI{}
	api.RegisterHandlers(mux)

	s := &http.Server{Addr: addr, Handler: mux}

	log.Println("starting spud-stories server. enjoy your 'taters.")
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("server closed unexpectedly: %s", err.Error())
	}
	log.Println("spud-stories server shut down gracefully.")
}
