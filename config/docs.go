package config

import (
    "encoding/json"
    "io/ioutil"
    "path/filepath"
)

type Docs struct {
    Endpoints []map[string] interface{}`json:"endpoints"`
}

func GetDocs() (*Docs, error) {
    data, err := ioutil.ReadFile(filepath.Join("assets", "docs.json"))
    
    if err != nil {
        return nil, err
    }
    
    docs := &Docs{}
    
    json.Unmarshal(data, docs)
    
    return docs, nil
}