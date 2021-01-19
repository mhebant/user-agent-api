package main

import (
	"fmt"
	"user-agent-api/useragent"
	"encoding/json"
)

func main() {
	var repo *useragent.Repository
	var err error
	repo, err = useragent.RepositoryFromFile("user-agents.json")
	b, _ := json.Marshal(repo)
	fmt.Println(err , string(b))
}
