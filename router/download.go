package router

import (
	"encoding/json"
	"github.com/bonbot-team/bonapi/generator"
	"github.com/bonbot-team/bonapi/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func DownloadRoute(res http.ResponseWriter, req *http.Request) {
	p := mux.Vars(req)

	genName := strings.ToLower(p["generator"])
	genPtr := generator.GetMgr().Get(genName)

	if genPtr == nil {
		err := utils.NewError(400, genName+" generator not found")
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

	names, _ := args["name"]
	name := names[0]

	res.Header().Set("Content-Disposition", "attachment; filename=bon_"+strings.ToLower(strings.ReplaceAll(name, " ", "_"))+".png")
	res.Header().Set("Content-Type", "image/png")

	res.WriteHeader(200)
	_, _ = res.Write(bytes)
}
