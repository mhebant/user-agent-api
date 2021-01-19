package useragent

import (
	"io/ioutil"
	"encoding/json"
)

type Info struct {
	User_Agents []string
	App *string
	Device *string
	Bot *bool
}

type Repository []Info

func (*Repository) Query(ua string) *Info {
	return nil
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