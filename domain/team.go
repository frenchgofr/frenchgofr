package domain

import (
	"log"
	"sync"

	"github.com/nlopes/slack"
)

// Team information
type Team struct {
	mu      sync.RWMutex
	iconURL string
	name    string
	domain  string
	cocURL  string
}

func (t *Team) Update(s *slack.TeamInfo) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.name = s.Name
	t.domain = s.Domain
	if v, ok := s.Icon["image_default"]; ok {
		if b, ok := v.(bool); ok && b {
			t.iconURL = ""
			return
		}
	}
	var icons = []string{"132", "102", "88", "68", "44", "34"}
	for _, i := range icons {
		img, ok := s.Icon["image_"+i]
		if ok {
			if str, ok := img.(string); ok {
				t.iconURL = str
			}
			return
		}
	}
	log.Println("Unable to determine icon image")
}

//Icon information for the teams
func (t *Team) Icon() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.iconURL
}

// Name of the Team
func (t *Team) Name() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.name
}

// Domain of the Team
func (t *Team) Domain() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.domain
}
