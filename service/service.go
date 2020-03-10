package service

import (
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
// It returns the result of the fizzbuzz algorithm on success
func GetFizzbuzzStrings(ds datastore.Datastore, request *models.Request) ([]string, error) {
	strings, err := core.Fizzbuzz(request.Int1Param,
		request.Int2Param,
		request.LimitParam,
		request.Str1Param,
		request.Str2Param)
	if err != nil {
		return nil, err
	}

	// Store new request
	// TODO async db store
	err = ds.Store(request)
	if err != nil {
		return nil, err
	}
	return strings, nil
}
