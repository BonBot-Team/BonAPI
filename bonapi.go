package main

import (
    "log"
    "net/http"
    "github.com/bonbot-team/bonapi/config"
    "github.com/bonbot-team/bonapi/router"
)

func main(){
    Config, err := config.GetConfig()
    
    if err != nil {
        log.Println("Cannot get config")
        return 
    }
    
    router := router.Init()
    
    srv := &http.Server {
        Addr: ":" + Config.Port,
        Handler: router,
    }
    
    log.Println("Web server started on port " + Config.Port)
    
    log.Fatal(srv.ListenAndServe())
}
