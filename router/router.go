package router

import (
    "text/template"
    "path/filepath"
    "github.com/julienschmidt/httprouter"
)

func Init() (*httprouter.Router) {
    router := httprouter.New()
    
    router.GET("/docs", DocsRoute)
    router.GET("/api/create/:generator", CreateRoute)
    
    return router
}

func GetTemplate(name string) (*template.Template, error){
    t, err := template.ParseFiles(filepath.Join("views", name + ".templ"))
    
    if err != nil {
        return nil, err
    }
    
    return t, nil
}