package main

import (
	"fmt"
	"user-agent-api/useragent"
)

func main() {
	repo, _ := useragent.RepositoryFromFile("user-agents.json")
	fmt.Println(repo.Query("Overcast ( http://overcast.fm/; Apple Watch podcast app)"))
	fmt.Println(repo.Query("Overcast/1.0 Podcast Sync"))
	fmt.Println(repo.Query("de.danoeh.antennapod/1.7.3b (Linux;Android 8.0.0) ExoPlayerLib/2.9.3"))
}
