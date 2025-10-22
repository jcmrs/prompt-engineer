package gemini

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Wrapper interface {
	CheckAuth(ctx context.Context) error
	RunChatStreaming(ctx context.Context, model string, input string, settings map[string]interface{}, onToken func(token string, idx int, isFinal bool)) (finalResult string, err error)
	Embeddings(ctx context.Context, text string) ([]float32, error)
}

func NewWrapperFromEnv() Wrapper {
	if os.Getenv("PEA_GEMINI_MOCK") == "true" {
		return &MockWrapper{}
	}
	return &realWrapper{}
}

type MockWrapper struct{}

func (m *MockWrapper) CheckAuth(ctx context.Context) error {
	return nil
}

func (m *MockWrapper) RunChatStreaming(ctx context.Context, model string, input string, settings map[string]interface{}, onToken func(token string, idx int, isFinal bool)) (finalResult string, err error) {
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
			onToken(fmt.Sprintf("token-%d ", i), i, false)
			time.Sleep(100 * time.Millisecond)
		}
	}
	onToken("This is the final content.", 5, true)
	return "This is the final content.", nil
}

func (m *MockWrapper) Embeddings(ctx context.Context, text string) ([]float32, error) {
	return []float32{0.1, 0.2, 0.3}, nil
}

type realWrapper struct{}

func (r *realWrapper) CheckAuth(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "gemini", "whoami", "--format=json")
	if err := cmd.Run(); err != nil {
		cmd = exec.CommandContext(ctx, "gemini", "auth", "status", "--format=json")
		if err := cmd.Run(); err != nil {
			cmd = exec.CommandContext(ctx, "gemini", "chat", "--model=gemini-2.5-flash", "--format=json")
			cmd.Stdin = os.NewFile(0, "stdin")
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("failed to authenticate with Gemini CLI: %w", err)
			}
		}
	}
	return nil
}

func (r *realWrapper) RunChatStreaming(ctx context.Context, model string, input string, settings map[string]interface{}, onToken func(token string, idx int, isFinal bool)) (finalResult string, err error) {
	// [TODO-JULES] Implement real Gemini CLI streaming, including process group management for cancellation.
	return "", fmt.Errorf("not implemented")
}

func (r *realWrapper) Embeddings(ctx context.Context, text string) ([]float32, error) {
	// [TODO-JULES] Implement real Gemini CLI embeddings.
	return nil, fmt.Errorf("not implemented")
}
