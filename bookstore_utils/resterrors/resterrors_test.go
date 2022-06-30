package resterrors

import (
	"reflect"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *RestError
	}{
		{
			name: "correct newBadRequestError",
			args: args{
				message: "invalid json",
			},
			want: NewBadRequestError("invalid json"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBadRequestError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBadRequestError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInternalServerError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *RestError
	}{
		// TODO: Add test cases.
		{
			name: "correct internal server error",
			args: args{
				message: "error here",
			},
			want: NewInternalServerError("error here"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInternalServerError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInternalServerError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUnauthorizedError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *RestError
	}{
		// TODO: Add test cases.
		{
			name: "correct unauthorizedError",
			args: args{
				message: "unauthorized",
			},
			want: NewUnauthorizedError("unauthorized"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUnauthorizedError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnauthorizedError() = %v, want %v", got, tt.want)
			}
		})
	}
}
