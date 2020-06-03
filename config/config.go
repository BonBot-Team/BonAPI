package config

import (
    "encoding/json"
    "io/ioutil"
    "path/filepath"
)

type Config struct {
    Port string `json:"port"`
}

func GetConfig() (*Config, error) {
    data, err := ioutil.ReadFile(filepath.Join("assets", "config.json"))
    
    if err != nil {
        return nil, err
    }
    
    config := &Config{}
    
    json.Unmarshal(data, config)
    
    return config, nil
}