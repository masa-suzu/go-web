package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	fileName string
	template *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t.once.Do(func() {
		t.template = template.Must(template.ParseFiles(filepath.Join("templates", t.fileName)))
	})
	_ = t.template.Execute(w, req)
}

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{fileName: "chat.html"})
	http.Handle("/room", r)
	go r.run()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
