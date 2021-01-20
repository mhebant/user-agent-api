package main

import (
	"fmt"
	"user-agent-api/useragent"
	"net/http"
	"log"
)

const repositoryFile = "user-agents.json"
const addr = ":8080"

func main() {
	repo, repoErr := useragent.RepositoryFromFile(repositoryFile)
	if repoErr != nil {
		log.Fatal(repoErr)
	}
	log.Printf("user-agent repository loaded from file %v", repositoryFile)
	
	http.HandleFunc("/info", func(resp http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			resp.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(resp, "only GET method is suported")
			return
		}

		ua := req.URL.Query().Get("ua")
		if ua == "" {
			resp.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(resp, "query parameter required")
			return
		}

		uaInfo := repo.Query(ua)
		resp.WriteHeader(http.StatusOK)
		fmt.Fprint(resp, uaInfo.ToJson())
		log.Printf("query user-agent \"%v\"", ua)
	})
	
	log.Printf("starting listening on %v", addr)
	httpErr := http.ListenAndServe(addr, nil)
	if httpErr != nil {
		log.Fatal(httpErr)
	}
}
