package useragent

import (
	"encoding/json"
)

type Info struct {
	User_Agents []string
	App string
	Device string
	Bot bool
}

func (info *Info) ToJson() string {
	m := map[string]interface{}{
		"app": info.App,
		"device": info.Device,
		"bot": info.Bot,
	}
	b, _ := json.Marshal(m)
	return string(b)
}