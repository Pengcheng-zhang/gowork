package services

import (
	"crypto/aes"
	"bytes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

type EncryptionService struct {
	key []byte
}
//AES-128 Key长度：16， 24， 32 bytes 对应 AES-128,AES-192,AES-256

func (this *EncryptionService) initEncryptionService() error{
	this.key = []byte(GetConfigValue("encryption", "key"))
	if len(this.key) == 0 {
		return errors.New("失败")
	}
	return nil
}

func AesEncrypt(originData []byte) ([]byte, error) {
	service := &EncryptionService{}
	err := service.initEncryptionService()
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(service.key)
	if err != nil {
		Error("data encrypt failed:", err.Error())
		return nil,err
	}
	blockSize := block.BlockSize()
	originData = pkcs5Pading(originData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, service.key[:blockSize])
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return crypted, nil
}

func AesDecrypt(crypted []byte) ([]byte, error) {
	service := &EncryptionService{}
	err := service.initEncryptionService()
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(service.key)
	if err != nil {
		Error("data decrypt failed:", err.Error())
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, service.key[:blockSize])
	originData := make([]byte, len(crypted))
	blockMode.CryptBlocks(originData, crypted)
	originData = pkcs5UnPadding(originData)
	return originData, nil
}

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	paddingText := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, paddingText...)
}
func pkcs5Pading(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, paddingText...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	//去掉最后一个字节 unpadding 次
	unPadding := int(origData[length - 1])
	return origData[:(length - unPadding)]
}

func Test()  {
	result, err := AesEncrypt([]byte("zhangpch666"))
	if err != nil {
		panic(err)
	}
	Debug(base64.StdEncoding.EncodeToString(result))

	originData, err := AesDecrypt(result)

	if err != nil {
		panic(err)
	}
	Debug(string(originData))
}
