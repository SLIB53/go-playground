package api

import (
	"context"
	"testing"
)

type TestLogger struct{}

func (l *TestLogger) LogInfo(message string)             {}
func (l *TestLogger) LogError(message string, err error) {}

func TestSayHello(t *testing.T) {
	testctx := context.Background()
	testlogger := &TestLogger{}

	type args struct {
		in0    context.Context
		logger Logger
		name   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "empty name",
			args: args{
				in0:    testctx,
				logger: testlogger,
				name:   "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "name too long",
			args: args{

				in0:    testctx,
				logger: testlogger,
				name:   "123456789012345678901234567890123",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "valid name",
			args: args{

				in0:    testctx,
				logger: testlogger,
				name:   "John",
			},
			want:    "Hello, John!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SayHello(tt.args.in0, tt.args.logger, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createHello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "empty name",
			args: args{
				name: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "name too long",
			args: args{
				name: "123456789012345678901234567890123",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "valid name",
			args: args{
				name: "Akil",
			},
			want:    "Hello, Akil!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeHello(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("createHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
