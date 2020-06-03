package router

import (
    "net/http"
    "encoding/json"
    "strings"
    "github.com/bonbot-team/bonapi/utils"
    "github.com/bonbot-team/bonapi/generator"
    "github.com/julienschmidt/httprouter"
)

func CreateRoute(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
    
    genName := strings.ToLower(p.ByName("generator"))
    genPtr := generator.GetMgr().Get(genName)
    
    if genPtr == nil {
        err := utils.NewError(400, genName + " generator not found")
        json, _ := json.Marshal(err)
        
        res.WriteHeader(err.Code)
        res.Write(json)
        return
    }
    
    gen := *genPtr
    
    args := req.URL.Query()
    
    bytes, err := gen.Generate(args)
    
    if err != nil {
        json, _ := json.Marshal(err)
        
        res.WriteHeader(err.Code)
        res.Write(json)
        return
    }
    
    res.Header().Set("Content-Type", "image/png")
    res.WriteHeader(200)
    res.Write(bytes)
}
