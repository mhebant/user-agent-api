package useragent

import (
	"io/ioutil"
	"encoding/json"
	"regexp"
)

type Info struct {
	User_Agents []string
	App string
	Device string
	Bot bool
}

func (info *Info) ToJson() string {
	m := make(map[string]interface{})
	m["app"] = info.App
	m["device"] = info.Device
	m["bot"] = info.Bot
	b, _ := json.Marshal(m)
	return string(b)
}

type Repository []Info

func (repo *Repository) Query(ua string) *Info {
	for _, info := range *repo {
		for _, re := range info.User_Agents {
			if match, err := regexp.MatchString(re, ua); err==nil && match {
				return &info
			} 
		} 
	}
	return &Info{}
}

func RepositoryFromFile(file string) (Repository, error) {
	var data []byte
	var err error

	data, err = ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	
	var repo Repository
	err = json.Unmarshal(data, &repo)
	if err != nil {
		return nil, err
	}

	return repo, nil
}