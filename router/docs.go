package router

import (
    "net/http"
    "github.com/bonbot-team/bonapi/config"
    "github.com/julienschmidt/httprouter"
)

func DocsRoute(res http.ResponseWriter, req *http.Request, p httprouter.Params){
    t, err := GetTemplate("docs")
    
    if err != nil {
        res.Write([]byte("An error occured\n" + err.Error()))
    }
    
    docs, err := config.GetDocs()
    
    if err != nil {
        res.Write([]byte("An error occured\n" + err.Error()))
    }
    
    t.Execute(res, docs)
}