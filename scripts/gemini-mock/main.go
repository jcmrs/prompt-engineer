package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	meta := map[string]interface{}{
		"type":  "meta",
		"model": "gemini-2.5-flash",
		"usage": map[string]interface{}{
			"tokens": 123,
		},
	}
	metaJSON, _ := json.Marshal(meta)
	fmt.Println(string(metaJSON))
	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 5; i++ {
		token := map[string]interface{}{
			"type":        "token",
			"data":        fmt.Sprintf("token-%d ", i),
			"chunk_index": i,
			"is_final":    false,
		}
		tokenJSON, _ := json.Marshal(token)
		fmt.Println(string(tokenJSON))
		time.Sleep(100 * time.Millisecond)

		progress := map[string]interface{}{
			"type":    "progress",
			"percent": (i + 1) * 20,
		}
		progressJSON, _ := json.Marshal(progress)
		fmt.Println(string(progressJSON))
		time.Sleep(100 * time.Millisecond)
	}

	final := map[string]interface{}{
		"type":    "final",
		"content": "This is the final content.",
		"metrics": map[string]interface{}{},
	}
	finalJSON, _ := json.Marshal(final)
	fmt.Println(string(finalJSON))
}
