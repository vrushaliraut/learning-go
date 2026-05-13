package main

import (
	"encoding/json"
	"fmt"
)

type webhookPod struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func main() {
	webhookPodString := `{"name"": "test-pod", "status":"Running"}`
	var pod webhookPod
	err := json.Unmarshal([]byte(webhookPodString), &pod)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pod)
}
