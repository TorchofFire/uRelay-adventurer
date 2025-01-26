package profile

import (
	"encoding/json"
	"log"
	"os"
)

// TODO: rework this. This is temporary and profiles should be secured in the system's vault.

type Service struct {
	Profile profile
}

func NewService() *Service {
	s := &Service{}
	s.Init()
	return s
}

type profile struct {
	Name       string `json:"name"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

func (s *Service) Init() {
	file, err := os.Open("profile.json")
	if err != nil {
		log.Fatalf("couldn't find profile.json:%v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&s.Profile); err != nil {
		log.Fatalf("couldn't decode profile.json:%v", err)
	}
}
