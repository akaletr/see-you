package app

import (
	"cmd/main/main.go/internal/storage"
	"net/http"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want App
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_app_AddUser(t *testing.T) {
	type fields struct {
		storage storage.Storage
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &app{
				storage: tt.fields.storage,
			}
			app.AddUser(tt.args.w, tt.args.r)
		})
	}
}

func Test_app_GetUser(t *testing.T) {
	type fields struct {
		storage storage.Storage
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &app{
				storage: tt.fields.storage,
			}
			app.GetUser(tt.args.w, tt.args.r)
		})
	}
}

func Test_app_Start(t *testing.T) {
	type fields struct {
		storage storage.Storage
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &app{
				storage: tt.fields.storage,
			}
			if err := app.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
