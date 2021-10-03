package api_test

import (
	"goModul/api"
	"testing"
)

func Test(t *testing.T) {
	_, err := api.GetAllCategories()
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

//go test -coverprofile=coverage.out && go tool cover -html=coverage.out
