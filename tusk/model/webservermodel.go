package model

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type PageHandler struct {
	Parent *WebServer
}

func NewPageHandler(ws *WebServer) *PageHandler {
	p := new(PageHandler)
	p.Parent = ws
	return p
}

func (p *PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if f, ok := p.Parent.Muxer[r.URL.String()]; ok {
		f(w, r)
	}
}

type WebServer struct {
	Server   http.Server
	PHandler *PageHandler
	Config   *TuskConfig
	DB       *DatabaseManager
	Muxer    map[string]func(http.ResponseWriter, *http.Request)
}

func NewWebServer(tc *TuskConfig) *WebServer {
	ws := new(WebServer)
	ws.PHandler = NewPageHandler(ws)
	ws.Config = tc
	ws.Server = http.Server{
		Addr:    tc.Socket,
		Handler: ws.PHandler,
	}
	ws.Muxer = make(map[string]func(http.ResponseWriter, *http.Request))
	ws.DB = NewDatabaseManager(tc.DBUser, tc.DBPass, tc.DBSocket, "mysql", tc.DBName)
	ws.DB.Init()
	return ws
}

func (ws *WebServer) Listen() {
	ws.Server.ListenAndServe()
}

func (ws *WebServer) RegisterPage(uri string, f func(http.ResponseWriter, *http.Request)) {
	ws.Muxer[uri] = f
}

func (ws *WebServer) Welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Tusk: Long Term Test Result Store and Analytics")
}

func (ws *WebServer) Query(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	HandleError(err)
	var results = []*TestResult{}
	var query *[]Query
	json.Unmarshal(body, &query)
	for _, q := range *query {
		queryFunc := ws.DB.GetQueryFunction(q.QueryType)
		results = queryFunc(&q)
	}
	json.NewEncoder(w).Encode(results)
}

func (ws *WebServer) Upload(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	HandleError(err)
	var results *[]TestResult
	json.Unmarshal(body, &results)
	for _, r := range *results {
		fmt.Println(r)
		ws.DB.Insert(&r)
	}
}
