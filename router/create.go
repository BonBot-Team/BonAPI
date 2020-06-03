package router

import (
    "net/http"
    "strings"
    "github.com/bonbot-team/bonapi/generator"
    "github.com/julienschmidt/httprouter"
)

func CreateRoute(res http.ResponseWriter, req *http.Request, p httprouter.Params){
    
    genName := strings.ToLower(p.ByName("generator"))
    genPtr := generator.GetMgr().Get(genName)
    
    if genPtr == nil {
        res.Write([]byte("Generator : " + genName + " not found"))
        return
    }
    
    gen := *genPtr
    
    args := req.URL.Query()
    
    bytes, err := gen.Generate(args)
    
    if err != nil {
        res.Write([]byte(err.Error()))
        return
    }
    
    res.Write(bytes)
}