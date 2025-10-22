package unit

import (
	"context"
	"os"
	"testing"

	"github.com/jcmrs/prompt-engineer/internal/gemini"
)

func TestMockWrapper(t *testing.T) {
	os.Setenv("PEA_GEMINI_MOCK", "true")
	defer os.Unsetenv("PEA_GEMINI_MOCK")

	w := gemini.NewWrapperFromEnv()
	if _, ok := w.(*gemini.MockWrapper); !ok {
		t.Fatal("expected mock wrapper")
	}

	if err := w.CheckAuth(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	finalResult, err := w.RunChatStreaming(context.Background(), "", "", nil, func(token string, idx int, isFinal bool) {})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if finalResult != "This is the final content." {
		t.Fatalf("unexpected final result: %s", finalResult)
	}
}
