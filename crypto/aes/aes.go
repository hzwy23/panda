package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"sync"
)

// AES加密
// key 存储的是密钥信息，目前支持16位，24位，32位长度的密钥
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
type AES struct {
	key  []byte
	lock *sync.RWMutex
}

// 创建AES对象, 参数key是AES对象的密钥信息
// key长度必须是16,24,32位长度，否则无法创建AES对象实例。
func NewAES(key string) (*AES, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("The key length can only be 16,24,32")
	}
	r := &AES{
		key:  []byte(key),
		lock: new(sync.RWMutex),
	}
	return r, nil
}

// 对明文进行加密处理，plainText是明文信息
// 如果加密成功，返回两个值，第一个是明文对应的密文，第二个返回值是nil
// 如果加密失败，第一个返回值是脏数据，第二个返回值是错误信息
func (r *AES) Encrypt(plainText string) (string, error) {
	origData := []byte(plainText)

	r.lock.RLock()
	defer r.lock.RUnlock()
	block, err := aes.NewCipher(r.key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = r.pkcs5Padding(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, r.key[:blockSize])
	crypted := make([]byte, len(origData))

	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

// 解密数据，cryptotext 是密文信息
// 解密成功返回两个参数，第一个返回值是密文对应的明文，第二个是nil
// 解密失败，则第一个返回值是脏数据，第二个返回值是错误信息
func (r *AES) Decrypt(cryptotext string) (string, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	crypted, err := base64.StdEncoding.DecodeString(cryptotext)
	if err != nil {
		return "", errors.New("base64 decode failed.")
	}
	block, err := aes.NewCipher(r.key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, r.key[:blockSize])
	origData := make([]byte, len(crypted))

	blockMode.CryptBlocks(origData, crypted)
	origData, err = r.pkcs5UnPadding(origData)
	return string(origData), err
}

func (r *AES) SetKey(key string) error {
	switch len(key) {
	case 16, 24, 32:
		r.lock.Lock()
		r.key = []byte(key)
		r.lock.Unlock()
		return nil
	default:
		return errors.New("The key length can only be 16,24,32")
	}
}

func (r *AES) pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	rst := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, rst...)
}

func (r *AES) pkcs5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("cryptoText is empty")
	}
	padding := int(origData[length-1])

	if padding > length {
		return nil, errors.New("the key is invalid")
	}
	return origData[:length-padding], nil
}

var defaultAES *AES

func Encrypt(dt string) (string, error) {
	return defaultAES.Encrypt(dt)
}

func Decrypt(dt string) (string, error) {
	return defaultAES.Decrypt(dt)
}

func SetKey(key string) error {
	return defaultAES.SetKey(key)
}

func init() {
	defaultAES, _ = NewAES("hzwy23@hustwb09y")
}
