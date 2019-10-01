package main

import (
	"flag"
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
	addr := flag.String("addr", ":8080", "app address")
	flag.Parse()

	r := newRoom()
	http.Handle("/", &templateHandler{fileName: "chat.html"})
	http.Handle("/room", r)
	go r.run()

	log.Println("Started with Port", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
