package integration

import (
	"encoding/json"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/jcmrs/prompt-engineer/internal/server"
)

func TestBasicRun(t *testing.T) {
	os.Setenv("PEA_GEMINI_MOCK", "true")
	defer os.Unsetenv("PEA_GEMINI_MOCK")

	s := server.NewServer()
	ts := httptest.NewServer(s.Handler)
	defer ts.Close()

	wsURL := "ws" + strings.TrimPrefix(ts.URL, "http") + "/ws/run/demo-run?token=test"
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("failed to connect to websocket: %v", err)
	}
	defer conn.Close()

	var receivedFinal bool
	for i := 0; i < 6; i++ {
		_, p, err := conn.ReadMessage()
		if err != nil {
			t.Fatalf("failed to read message from websocket: %v", err)
		}

		var msg map[string]interface{}
		if err := json.Unmarshal(p, &msg); err != nil {
			t.Fatalf("failed to unmarshal message: %v", err)
		}

		if msg["type"] != "token" {
			t.Fatalf("unexpected message type: %s", msg["type"])
		}

		if isFinal, ok := msg["is_final"].(bool); ok && isFinal {
			receivedFinal = true
		}
	}

	if !receivedFinal {
		t.Fatal("did not receive final message")
	}
}
