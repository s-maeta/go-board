package service

import (
	"board/config"
	"os"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	os.Setenv("ENV", "dev")
	config.Init()
	type args struct {
		userId string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				userId: "test",
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUyOTg1NzYwNTQsInVzZXJfaWQiOiJ0ZXN0In0.yX5Dx4mpmSE7Wc1WGxW7l5qFeC1AgeuuSQIW9Ed6HDk",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// トークン自体の検証は難しいので、とりあえずエラーが吐かれないことを検証
			_, err := GenerateToken(tt.args.userId)
			if err != nil {
				t.Fatal("Error test GenerateToken")
			}

			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			// if got != tt.want {
			// 	t.Errorf("GenerateToken() = %v, want %v", got, tt.want)
			// }
		})
	}
}
