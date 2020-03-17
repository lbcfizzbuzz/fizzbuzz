package service

import (
	"errors"
	"github.com/lbcfizzbuzz/fizzbuzz/core"
	"github.com/lbcfizzbuzz/fizzbuzz/datastore"
	"github.com/lbcfizzbuzz/fizzbuzz/models"
)

// GetMostUsedQueryString returns the most used requests from a datastore
func GetMostUsedQueryString(ds datastore.Datastore) (models.Request, error) {
	return ds.FindByMostUsedQueryString()
}

// GetFizzbuzzStrings launch the fizzbuzz algorithm using the given parameters and
// store the request in the given datastore.
// It returns the a slice containing each result of the fizzbuzz algorithm on success
func GetFizzbuzzStrings(ds datastore.Datastore, request *models.Request) ([]string, error) {
	if request.Int1Param == 0 || request.Int2Param == 0 {
		return nil, errors.New("the int1 and int2 parameters must be greater than 0")
	}
	if request.LimitParam > 1000000 { // arbitrary limit our server doesn't scale :(
		return nil, errors.New("the limit parameter must be smaller than 1000000")
	}

	var result []string
	for i := uint64(1); i <= request.LimitParam; i++ {
		result = append(result, core.Fizzbuzz(i, request.Int1Param,
			request.Int2Param,
			request.Str1Param,
			request.Str2Param))
	}

	// Store new request
	err := ds.Store(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
