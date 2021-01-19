package main

import (
	"fmt"
	"user-agent-api/useragent"
	"net/http"
)

func main() {
	repo, _ := useragent.RepositoryFromFile("user-agents.json")
	
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
	})
	http.ListenAndServe(":8080", nil)
}
