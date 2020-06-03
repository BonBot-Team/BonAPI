package main

import (
    "fmt"
    "net/http"
    "github.com/bonbot-team/bonapi/config"
    "github.com/bonbot-team/bonapi/router"
)

func main(){
    Config, err := config.GetConfig()

    router := router.Init()
    
    if err != nil {
        fmt.Println("Cannot get config")
    }
    
	http.ListenAndServe(":" + Config.Port, router)
	
	log("Web server started on port " + Config.Port)
}

func log(msg string){
    fmt.Println("[BONAPI]", msg)
}