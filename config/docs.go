package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Docs struct {
	Endpoints []map[string]interface{} `json:"endpoints"`
}

func GetDocs() (*Docs, error) {
	if _, err := os.Stat("docs.json"); os.IsNotExist(err) {
		docs := &Docs{}

		bytes, _ := json.Marshal(docs)

		if err = ioutil.WriteFile("docs.json", bytes, 0777); err != nil {
			return nil, err
		}
	}

	data, err := ioutil.ReadFile("docs.json")

	if err != nil {
		return nil, err
	}

	docs := &Docs{}

	_ = json.Unmarshal(data, docs)

	return docs, nil
}
