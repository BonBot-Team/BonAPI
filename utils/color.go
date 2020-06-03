package utils

import (
    "regexp"
    "gopkg.in/go-playground/colors.v1"
)

func GetColor(color *string) {
    matched, err := regexp.Match("^[0-9A-F]{6}$", []byte(*color))
    
    if err == nil {
        if matched {
            *color = "#" + *color
        }
    }
}

func ConvertColors(ca []string) []string {
    
    for index, color := range ca {
        GetColor(&color)
        c, err := colors.Parse(color)
        
        if err != nil {
            ca[index] = "#4287f5"
        } else {
            hex := c.ToHEX()
            ca[index] = hex.String()
        }
    }
    
    return ca
}
