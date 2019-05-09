package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

type LoginData struct {
	Msg      string `json:"msg"`
	TokenStr string `json:"tokenStr"`
}

type Login struct {
	Data LoginData `json:"data"`
	Code string    `json:"code"`
}

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	MaxAge     int
	Secure     bool
	HttpOnly   bool
	Raw        string
	Unparsed   []string
}

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"build/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	_ = templates.Execute(w, "index")
}

func randToken() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func login(w http.ResponseWriter, r *http.Request) {
	var user map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &user)
	fmt.Println(user["userAccount"])
	fmt.Println(user["pwd"])

	w.Header().Set("Content-Type", "application/json")
	post := &Login{ // 伪造数据
		Code: "1",
		Data: LoginData{
			Msg:      "OK",
			TokenStr: randToken(),
		},
	}
	json2, _ := json.Marshal(post)
	_, _ = w.Write(json2)
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("build/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/cas/user/login", login)

	server := &http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: mux,
	}
	_ = server.ListenAndServe()
}
