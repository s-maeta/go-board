package service

import (
	"crypto/md5"
	"fmt"
	"io"
)

/*
文字列を暗号化する（MD5)
*/
func Encrypt(text string) (string, error) {
	var err error

	//ハッシュ値生成
	hash := md5.New()
	_, err = io.WriteString(hash, text)
	if err != nil {
		return "", err
	}
	//ハッシュパスワード生成
	passwordMd5 := fmt.Sprintf("%x", hash.Sum(nil))

	//ソルトでをつけて再度暗号化
	passwordSalt := "soltfirst"
	_, err = io.WriteString(hash, passwordSalt)
	if err != nil {
		return "", err
	}
	_, err = io.WriteString(hash, passwordMd5)
	if err != nil {
		return "", err
	}

	password := fmt.Sprintf("%x", hash.Sum(nil))

	return password, nil

}

/*
暗号化したハッシュと比較する
*/
func VerifyPassword(plainText, storedHashedPassword string) bool {
	hashedPassword, err := Encrypt(plainText)
	if err != nil {
		return false
	}

	// 保存済みのハッシュと生成したハッシュを比較
	return hashedPassword == storedHashedPassword
}
