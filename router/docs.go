package router

import (
    "github.com/bonbot-team/bonapi/config"
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func DocsRoute(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
    t, err := GetTemplate("docs")
    
    if err != nil {
        _, _ = res.Write([]byte("An error occured\n" + err.Error()))
    }
    
    docs, err := config.GetDocs()
    
    if err != nil {
        _, _ = res.Write([]byte("An error occured\n" + err.Error()))
    }

    _ = t.Execute(res, docs)
}