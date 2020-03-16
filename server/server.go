package server

import (
	"encoding/json"
	cfg "github.com/lbcfizzbuzz/fizzbuzz/config"
	ds "github.com/lbcfizzbuzz/fizzbuzz/datastore"
	"github.com/lbcfizzbuzz/fizzbuzz/internal/constants"
	models "github.com/lbcfizzbuzz/fizzbuzz/models"
	"github.com/lbcfizzbuzz/fizzbuzz/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strconv"
)

// Server represents a server listening for requests
type Server struct {
	Config *cfg.Configuration
	Db     ds.Datastore
	Logger *log.Logger
}

// Error represents an error that will be sent to the client
type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (s *Server) statisticsHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Infoln("received request: " + r.RequestURI)

	w.Header().Set("Content-Type", "application/json")

	response, err := service.GetMostUsedQueryString(s.Db)
	if err != nil {
		s.Logger.Errorln(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: constants.ErrorMostUsedQueryFailed, Status: http.StatusInternalServerError})
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (s *Server) fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Infoln("received request: " + r.RequestURI)

	w.Header().Set("Content-Type", "application/json")

	params, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil {
		s.Logger.Errorln(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: constants.ErrorQueryString, Status: http.StatusBadRequest})
		return
	}

	// Get numeric params
	int1, err := strconv.ParseUint(params.Get("int1"), 10, 64)
	if err != nil {
		s.Logger.Errorln(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: constants.ErrorInt1Param, Status: http.StatusBadRequest})
		return
	}
	int2, err := strconv.ParseUint(params.Get("int2"), 10, 64)
	if err != nil {
		s.Logger.Errorln(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: constants.ErrorInt2Param, Status: http.StatusBadRequest})
		return
	}

	limit, err := strconv.ParseUint(params.Get("limit"), 10, 32)
	if err != nil {
		s.Logger.Errorln(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: constants.ErrorLimitParam, Status: http.StatusBadRequest})
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
		s.Logger.Errorln(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: constants.ErrorFizzbuzz, Status: http.StatusInternalServerError})
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
