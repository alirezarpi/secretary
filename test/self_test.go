package test

import (
	"net/http"
	"testing"

	"secretary/alpha/api"
)

func TestSelfAPI(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api.SelfAPI(tt.args.w, tt.args.r)
		})
	}
}
