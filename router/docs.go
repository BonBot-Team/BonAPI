package router

import (
	"github.com/bonbot-team/bonapi/config"
	"net/http"
)

func DocsRoute(res http.ResponseWriter, req *http.Request) {
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
