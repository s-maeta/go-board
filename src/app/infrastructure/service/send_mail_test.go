package service

import "testing"

func TestSendEmail(t *testing.T) {
	type args struct {
		from       string
		recipients []string
		subject    string
		body       string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "正常系",
			args: args{
				from:       "test@test.com",
				recipients: []string{"testto1@test.go"},
				subject:    "テストタイトル",
				body:       "テスト内容です。これはテスト内容です。",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendEmail(tt.args.from, tt.args.recipients, tt.args.subject, tt.args.body)
		})
	}
}
