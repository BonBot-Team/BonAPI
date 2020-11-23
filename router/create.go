package router

import (
    "encoding/json"
    "github.com/bonbot-team/bonapi/generator"
    "github.com/bonbot-team/bonapi/utils"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "strings"
)

func CreateRoute(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
    
    genName := strings.ToLower(p.ByName("generator"))
    genPtr := generator.GetMgr().Get(genName)
    
    if genPtr == nil {
        err := utils.NewError(400, genName + " generator not found")
        bytes, _ := json.Marshal(err)

        res.WriteHeader(err.Code)
        _, _ = res.Write(bytes)
        return
    }

    gen := *genPtr
    
    args := req.URL.Query()
    
    bytes, err := gen.Generate(args)
    
    if err != nil {
        bytes, _ := json.Marshal(err)
        
        res.WriteHeader(err.Code)
        _, _ = res.Write(bytes)
        return
    }
    
    res.Header().Set("Content-Type", "image/png")
    res.WriteHeader(200)
    _, _ = res.Write(bytes)
}
