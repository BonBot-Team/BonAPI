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
    
    router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("Access-Control-Request-Method") != "" {
            header := w.Header()
            
            header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
            header.Set("Access-Control-Allow-Origin", "*")
        }
        
        w.WriteHeader(http.StatusNoContent)
    })
    
    srv := &http.Server {
        Addr: ":" + Config.Port,
        Handler: router,
    }
    
    log.Println("Web server started on port " + Config.Port)
    
    log.Fatal(srv.ListenAndServe())
}
