package test

import (
	"net/http"
	"testing"

	"secretary/alpha/api"
)

func TestResponser(t *testing.T) {
	type args struct {
		w          http.ResponseWriter
		r          *http.Request
		status     bool
		statusCode int
		response   map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api.Responser(tt.args.w, tt.args.r, tt.args.status, tt.args.statusCode, tt.args.response)
		})
	}
}
