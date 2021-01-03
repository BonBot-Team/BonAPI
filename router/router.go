package router

import (
	"github.com/gorilla/mux"
	"path/filepath"
	"text/template"
)

func Init() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/docs", DocsRoute).Methods("GET")
	router.HandleFunc("/api/create/{generator}", CreateRoute).Methods("GET")
	router.HandleFunc("/api/download/{generator}", DownloadRoute).Methods("GET")

	return router
}

func GetTemplate(name string) (*template.Template, error) {
	t, err := template.ParseFiles(filepath.Join("views", name+".templ"))

	if err != nil {
		return nil, err
	}

	return t, nil
}
