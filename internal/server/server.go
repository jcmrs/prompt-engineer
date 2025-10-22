package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jcmrs/prompt-engineer/internal/gemini"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewServer() *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/auth/check", authCheckHandler).Methods("GET")
	r.HandleFunc("/conversations", createConversationHandler).Methods("POST")
	r.HandleFunc("/conversations", getConversationsHandler).Methods("GET")
	r.HandleFunc("/conversations/{id}", getConversationHandler).Methods("GET")
	r.HandleFunc("/conversations/{id}/messages", createMessageHandler).Methods("POST")
	r.HandleFunc("/runs", createRunHandler).Methods("POST")
	r.HandleFunc("/runs/{id}", getRunHandler).Methods("GET")
	r.HandleFunc("/attachments", createAttachmentHandler).Methods("POST")
	r.HandleFunc("/ws/run/{run_id}", wsRunHandler)

	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func authCheckHandler(w http.ResponseWriter, r *http.Request) {
	// [TODO-JULES] Implement auth check
	json.NewEncoder(w).Encode(map[string]interface{}{"authenticated": true, "message": "Authenticated"})
}

func createConversationHandler(w http.ResponseWriter, r *http.Request) {
	// [TODO-JULES] Implement create conversation
}

func getConversationsHandler(w http.ResponseWriter, r *http.Request) {
	// [TODO-JULES] Implement get conversations
}

func getConversationHandler(w http.ResponseWriter, r *http.Request) {
	// [TODO-JULES] Implement get conversation
}

func createMessageHandler(w http.ResponseWriter, r *http.Request) {
	// [TODO-JULES] Implement create message
}

func createRunHandler(w http.ResponseWriter, r *http.Request) {
	// [TODO-JULES] Implement create run
}

func getRunHandler(w http.ResponseWriter, r *http.Request) {
	// [TODO-JULES] Implement get run
}

func createAttachmentHandler(w http.ResponseWriter, r *http.Request) {
	// [TODO-JULES] Implement create attachment
}

func wsRunHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	wpr := gemini.NewWrapperFromEnv()
	_, err = wpr.RunChatStreaming(r.Context(), "", "", nil, func(token string, idx int, isFinal bool) {
		msg := map[string]interface{}{
			"type":        "token",
			"data":        token,
			"chunk_index": idx,
			"is_final":    isFinal,
		}
		jsonMsg, _ := json.Marshal(msg)
		conn.WriteMessage(websocket.TextMessage, jsonMsg)
	})

	if err != nil {
		msg := map[string]interface{}{
			"type":    "error",
			"message": err.Error(),
		}
		jsonMsg, _ := json.Marshal(msg)
		conn.WriteMessage(websocket.TextMessage, jsonMsg)
	}
}
