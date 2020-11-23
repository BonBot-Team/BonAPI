package main

import (
    "github.com/bonbot-team/bonapi/config"
    "github.com/bonbot-team/bonapi/router"
    "log"
    "net/http"
)

func main(){
    Config, err := config.GetConfig()
    
    if err != nil {
        log.Println("Cannot get config")
        return
    }

    handler := router.Init()

    srv := &http.Server {
        Addr: ":" + Config.Port,
        Handler: handler,
    }
    
    log.Println("Web server started on port " + Config.Port)
    
    log.Fatal(srv.ListenAndServe())
}
