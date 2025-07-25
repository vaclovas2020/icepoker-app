package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lpernett/godotenv"
)

const (
	pongWait = 60 * time.Second
)

func isValidUUIDv4(uuid string) bool {
	var re = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	return re.MatchString(uuid)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	if idString == "" || !isValidUUIDv4(idString) {
		w.WriteHeader(422)
		w.Write([]byte("id path param is empty"))
		return
	}

	log.Printf("ID: %s", idString)

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			log.Printf("Origin: %s", origin)
			return origin == fmt.Sprintf("https://%s/%s", os.Getenv("APP_HOST"), idString)
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}

	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
	}
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Printf("APP_HOST: %s\n", os.Getenv("APP_HOST"))
	log.Println("Starting server on port 7788")

	http.HandleFunc("GET /ws/{id}", handleWebSocket)
	http.ListenAndServe(":7788", nil)
}
