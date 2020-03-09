package server

import (
	"encoding/json"
	ds "github.com/samyy321/fizzbuzz/datastore"
	models "github.com/samyy321/fizzbuzz/models"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Server represents a server listening for requests
type Server struct {
	// config *Config
	Db ds.Datastore
}

func fizzbuzz(int1, int2, limit uint64, str1, str2 string) []string {
	var result []string

	for i := uint64(1); i <= limit; i++ {
		currentStr := ""
		if i%int1 == 0 {
			currentStr = str1
		}
		if i%int2 == 0 {
			currentStr += str2
		}
		if currentStr == "" {
			currentStr = strconv.FormatUint(i, 10)
		}
		result = append(result, currentStr)
	}

	return result
}

func (s *Server) statisticsHandler(w http.ResponseWriter, r *http.Request) {
	request, err := s.Db.FindByMostUsedQueryString()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(request)
}

func (s *Server) fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	params, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get numeric params
	int1, err := strconv.ParseUint(params.Get("int1"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	int2, err := strconv.ParseUint(params.Get("int2"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if int1 == 0 || int2 == 0 {
		http.Error(w, "The int1 and int2 parameters must be greater than 0", http.StatusInternalServerError)
		return
	}

	limit, err := strconv.ParseUint(params.Get("limit"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str1 := params.Get("str1")
	str2 := params.Get("str2")

	strList := fizzbuzz(int1, int2, limit, str1, str2)

	// Store new request
	// TODO async db store
	err = s.Db.Store(&models.Request{Int1Param: int1, Int2Param: int2, LimitParam: limit, Str1Param: str1, Str2Param: str2})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(strList)
}

// Run launch the server handlers and make it listen for requests
func (s *Server) Run() {
	http.HandleFunc("/fizzbuzz/", s.fizzbuzzHandler)
	http.HandleFunc("/statistics/", s.statisticsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
