package user

import (
	"fmt"
	"testing"
)

func Test_OK(t *testing.T) {
	newEmail := "ken_esp@icloud.com"
	email, err := NewEmail(newEmail)
	if err != nil {
		t.Fatal(email, err)
	}
	if *email != Email(newEmail) {
		t.Fatal(err, "作成したアドレスと引数のアドレスが一致しません。")
	}
}

func Test_ValidateRegexp(t *testing.T) {
	newEmail := "ken_espicliud.com"
	_, err := NewEmail(newEmail)

	if err == nil {
		t.Fatal("バリデーションでエラーになりませんでした")
	}
	fmt.Println(err)
	if err.Error() != "メールアドレスの形式が間違っています。" {
		t.Fatal("異なるバリデーションでエラーとなりました。", err)
	}
}
