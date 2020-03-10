package server

import (
	"encoding/json"
	conf "github.com/lbcfizzbuzz/fizzbuzz/config"
	ds "github.com/lbcfizzbuzz/fizzbuzz/datastore"
	models "github.com/lbcfizzbuzz/fizzbuzz/models"
	"github.com/lbcfizzbuzz/fizzbuzz/service"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Server represents a server listening for requests
type Server struct {
	Config *conf.Configuration
	Db     ds.Datastore
}

// Error represents an error that will be sent to the client
type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (s *Server) statisticsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := service.GetMostUsedQueryString(s.Db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error(), Status: http.StatusInternalServerError})
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (s *Server) fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	// Get numeric params
	int1, err := strconv.ParseUint(params.Get("int1"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error(), Status: http.StatusBadRequest})
		return
	}
	int2, err := strconv.ParseUint(params.Get("int2"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	limit, err := strconv.ParseUint(params.Get("limit"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	str1 := params.Get("str1")
	str2 := params.Get("str2")

	request := models.Request{
		Int1Param:  int1,
		Int2Param:  int2,
		LimitParam: limit,
		Str1Param:  str1,
		Str2Param:  str2}
	strList, err := service.GetFizzbuzzStrings(s.Db, &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error(), Status: http.StatusInternalServerError})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(strList)
}

// Run launch the server handlers and make it listen for requests
func (s *Server) Run() {
	http.HandleFunc("/fizzbuzz/", s.fizzbuzzHandler)
	http.HandleFunc("/statistics/", s.statisticsHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.Config.Port), nil))
}
