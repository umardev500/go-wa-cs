package sse

import (
	"sync"

	"github.com/rs/zerolog/log"
)

// Store all SSE clients categorized by event type
var (
	sseClients = make(map[string]map[string]chan string) // userid:resource:channel
	clientLock sync.Mutex
)

func GetSseResource(userId string, resource string) chan string {
	if _, ok := sseClients[userId]; !ok {
		return nil
	}

	if _, ok := sseClients[userId][resource]; !ok {
		return nil
	}

	return sseClients[userId][resource]
}

// Add a new client to a specific event type
func AddClient(userId string, resource string) chan string {
	clientLock.Lock()
	defer clientLock.Unlock()

	// Initialize the map for the eventType if it doesn't exist
	if _, ok := sseClients[userId]; !ok {
		sseClients[userId] = make(map[string]chan string)
	}

	// Add the client to the eventType map
	channel := make(chan string)
	sseClients[userId][resource] = channel

	return channel
}

// Remove a resource from a specific the client
func RemoveResource(userId string, resource string) {
	clientLock.Lock()
	defer clientLock.Unlock()

	if res, exists := sseClients[userId]; exists {
		delete(res, resource)
		close(res[resource])
		log.Info().Msgf("Clien unsubscribe from %s", resource)
	}
}

// Remove the client from sse data
func RemoveClient(userId string) {
	clientLock.Lock()
	defer clientLock.Unlock()

	delete(sseClients, userId)
}
