package cipher

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)

// NewRSAFromPEM 从 PEM 编码的公钥和私钥创建 RSA 实例
func NewRSAFromPEM(publicKey, privateKey []byte) (*RSA, error) {
	privatePem, _ := pem.Decode(privateKey)
	private, err := x509.ParsePKCS1PrivateKey(privatePem.Bytes)
	if err != nil {
		return nil, err
	}
	return &RSA{privateKey: private}, nil
}

// NewRSAFromPublicKeyPEM 从 PEM 编码的公钥创建 RSA 实例
func NewRSAFromPublicKeyPEM(publicKey []byte) (*RSA, error) {
	publicPem, _ := pem.Decode(publicKey)
	public, err := x509.ParsePKCS1PublicKey(publicPem.Bytes)
	if err != nil {
		return nil, err
	}
	return &RSA{privateKey: &rsa.PrivateKey{PublicKey: *public}}, nil
}

// NewRSAFromPrivateKeyPEM 从 PEM 编码的私钥创建 RSA 实例
func NewRSAFromPrivateKeyPEM(privateKey []byte) (*RSA, error) {
	privatePem, _ := pem.Decode(privateKey)
	private, err := x509.ParsePKCS1PrivateKey(privatePem.Bytes)
	if err != nil {
		return nil, err
	}
	return &RSA{privateKey: private}, nil
}

// NewRSA 创建一个新的 RSA 实例，使用指定的位数生成密钥对
func NewRSA(bits int) (*RSA, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return &RSA{privateKey: privateKey}, nil
}

type RSA struct {
	privateKey *rsa.PrivateKey
}

// Encrypt 使用 RSA 公钥加密数据
func (r *RSA) Encrypt(data []byte) ([]byte, error) {
	pubKey := &r.privateKey.PublicKey
	return rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
}

// Decrypt 使用 RSA 私钥解密数据
func (r *RSA) Decrypt(encrypted []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, encrypted)
}

// PublicKey 返回 RSA 公钥的 PEM 编码。
func (r *RSA) PublicKey() []byte {
	pubKey := r.privateKey.PublicKey
	pubKeyDer := x509.MarshalPKCS1PublicKey(&pubKey)
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyDer,
	})
}

// PrivateKey 返回 RSA 私钥的 PEM 编码。
func (r *RSA) PrivateKey() []byte {
	privateKeyDer := x509.MarshalPKCS1PrivateKey(r.privateKey)
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyDer,
	})
}

// Sign 使用 RSA 私钥对数据进行签名
func (r *RSA) Sign(data []byte) ([]byte, error) {
	hashed := r.sha256Sum(data)
	return rsa.SignPKCS1v15(rand.Reader, r.privateKey, crypto.SHA256, hashed)
}

// Verify 使用 RSA 公钥验证签名
func (r *RSA) Verify(data, signature []byte) error {
	hashed := r.sha256Sum(data)
	return rsa.VerifyPKCS1v15(&r.privateKey.PublicKey, crypto.SHA256, hashed, signature)
}

// sha256Sum 计算数据的 SHA-256 哈希值
func (r *RSA) sha256Sum(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}
