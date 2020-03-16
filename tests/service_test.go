package tests

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lbcfizzbuzz/fizzbuzz/models"
	"github.com/lbcfizzbuzz/fizzbuzz/service"
	"testing"
)

func TestService(t *testing.T) {
	datastore, err := GetInitializedDatastore()
	if err != nil {
		t.Errorf(err.Error())
	}

	// Create the request
	request := models.Request{
		Int1Param:  3,
		Int2Param:  5,
		LimitParam: 10,
		Str1Param:  "fizz",
		Str2Param:  "buzz"}

	// Use this request to do a fizzbuzz several times
	for i := 0; i < 3; i++ {
		_, err := service.GetFizzbuzzStrings(datastore, &request)
		if err != nil {
			t.Errorf("An unhandled error occured")
		}
	}

	// Change the request and use this new version more than the first one
	request.Int2Param = 6
	for i := 0; i < 4; i++ {
		_, err := service.GetFizzbuzzStrings(datastore, &request)
		if err != nil {
			t.Errorf("An unhandled error occured in GetFizzbuzzStrings()")
		}
	}

	mostUsedRequest, err := service.GetMostUsedQueryString(datastore)
	if err != nil {
		t.Errorf("An unhandled error occured in GetMostUsedQueryString()")
	}

	// We need to set the count to be able to compare correctly
	request.Count = mostUsedRequest.Count
	if mostUsedRequest != request {
		t.Errorf("The most used request is not the expected one")
	}
}
