package service

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "正常系",
			args:    args{text: "password"},
			want:    "08fd14c5f4250407728c41a9839fc5e3",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	type args struct {
		plainText            string
		storedHashedPassword string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "正常系_ハッシュ一致",
			args: args{plainText: "password", storedHashedPassword: "08fd14c5f4250407728c41a9839fc5e3"},
			want: true,
		},
		{
			name: "正常系_ハッシュ不一致",
			args: args{plainText: "password", storedHashedPassword: "08fd14c5f4250407728c41a9839fc5e3aaa"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyPassword(tt.args.plainText, tt.args.storedHashedPassword); got != tt.want {
				t.Errorf("VerifyPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
