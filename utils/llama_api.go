package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func LlamaAPI(prompt string) (string, error) {
	requestBody := map[string]interface{}{
		"model":  "llama3.2",
		"prompt": prompt,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("error to access llama: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	var resultText string

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error reading response line: %v", err)
			return "", err
		}

		var result map[string]interface{}
		if err := json.Unmarshal(line, &result); err != nil {
			log.Printf("Error unmarshalling line: %v", err)
			return "", err
		}

		if responsePart, ok := result["response"].(string); ok {
			resultText += responsePart
		}

		if done, ok := result["done"].(bool); ok && done {
			break
		}
	}

	return resultText, nil
}
