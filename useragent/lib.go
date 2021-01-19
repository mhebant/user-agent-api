package useragent

import (
	"io/ioutil"
	"encoding/json"
)

type Info struct {
	app *string
	device *string
	bot *bool
}

type Repository interface {
	Query(ua string) *Info
}

type JsonRepository struct {
	useragents []map[string]interface{}
}

func (*JsonRepository) Query(ua string) *Info {
	return nil
}

func RepositoryFromFile(file string) (*JsonRepository, error) {
	var data []byte
	var err error

	data, err = ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	
	var repo JsonRepository
	err = json.Unmarshal(data, &repo.useragents)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}