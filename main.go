package main

import (
	"fmt"
	"user-agent-api/useragent"
)

func main() {
	var repo useragent.Repository
	var err error
	repo, err = useragent.RepositoryFromFile("user-agents.json")
	fmt.Println(err , repo)
}
