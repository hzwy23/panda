package aes_test

import (
	"fmt"
	"testing"

	"github.com/hzwy23/panda/crypto/aes"
)

func TestDefaultCrypt(t *testing.T) {
	cryptoText, err := aes.Encrypt("123456")
	fmt.Println(cryptoText, err)

	plainText, err := aes.Decrypt(cryptoText)
	fmt.Println(plainText, err)

	err = aes.SetKey("1234567890123456")
	fmt.Println(err)
	cryptoText, err = aes.Encrypt("123456")
	fmt.Println(cryptoText, err)

	plainText, err = aes.Decrypt(cryptoText)
	fmt.Println(plainText, err)
}

func TestCustomerCrypto(t *testing.T) {
	aesobj, err := aes.NewAES("helloworld")
	fmt.Println(aesobj, err)

	aesobj, err = aes.NewAES("hzwy231234567654")
	fmt.Println(aesobj, err)

	aesobj, err = aes.NewAES("hzwy231234f67654")
	fmt.Println(aesobj, err)

	a, er := aesobj.Encrypt("weffdsdfadf")
	fmt.Println(a, er)
	s, e := aesobj.Decrypt(a)
	fmt.Println(s, e)
	aesobj.SetKey("hzwy231234f67654")
	s, e = aesobj.Decrypt(a)
	fmt.Println(s, e)

	aesobj, err = aes.NewAES("fdaefdertydsadfg")
	fmt.Println(aesobj, err)
	fmt.Println(a)
	fmt.Println(aesobj.Decrypt(a))

	aesobj, err = aes.NewAES("123456hzwyfdsadf3erfdsc3")
	fmt.Println(aesobj, err)
	fmt.Println(a)
	fmt.Println(aesobj.Decrypt(a))
}
