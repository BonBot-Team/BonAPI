package router

import (
    "github.com/julienschmidt/httprouter"
    "path/filepath"
    "text/template"
)

func Init() *httprouter.Router {
    router := httprouter.New()
    
    router.GET("/docs", DocsRoute)
    router.GET("/api/create/:generator", CreateRoute)
    router.GET("/api/download/:generator", DownloadRoute)
    
    return router
}

func GetTemplate(name string) (*template.Template, error) {
    t, err := template.ParseFiles(filepath.Join("views", name + ".templ"))
    
    if err != nil {
        return nil, err
    }
    
    return t, nil
}