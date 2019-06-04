package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	key        *Key
	message    *Message
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpwIBAAKCAQIAwXeoExjEyeKFu1Oj0Z9O48kC/uwZeNJSm5NEtEQ+ayMihRtv
Gi7DOAiiyHuVAo3nMKhtqiU71EVQ6BLxzCKQbPcc+dggpO5pX0HyuA3kdtliQD0J
z+VIvHAhXUa/i2Mit6aWIwgt+ytUvBd8sYFnuKjl/3bb5hEPC9/u93tFgSfPp86h
Udy97IjiEUxBXoKJ8ehaKgNJEzfvW1DoxUKcYCF0yb4pyNNPB92fbvXlPjZe8Te3
gUy4JeGo3eJr7Xa078aWgWOK+7pcdzYTHM6MXQpXjVLAoXLlCpCr0FjAXJqLEsNh
iMsEukoY9RhNfFymxPDQAx3Hbqjl+vAqIRu/RtUCAwEAAQKCAQIAu8vq2p6wZ0f7
iFsoKdL6QSJeRhXoo9+FUH8jsdiMvnLcj1iSAFhkJ4A6g2Fyw4f7YsAbs41xBhxC
7QN2szDaAOvetKeXCIJkxpK9iOvzWWqqdLDLVYK7mC9AM8r5I9SXXq3WDythdu4Y
0nv6DlQO2rEYkWJPEoR8lopI7PI8BaMEPYf+h5Nsdl6jsuGhQEop3vsrrBeknX6u
vMhrw/q/GiP6Gujr/JdqKutGD1TXXuFzN6E8m0R/N5Bu+CHN9z6pEJB39Jqh/uIo
0XCOMeXEBRtXnOtfYttBL2EhmyNhq4MkjHf12dlvQukxfHCkYR7dZ1Fi85rF03sA
bK3NXk7jMAECgYEMWYQWt8UoQcEIbEcWp4scD+uxCeJqV3lkP1U54GsNQnVLosgV
l7ncz10Ti2afrbqq1Nhw3EhMt8uuRIhXCLorvM+C+6wPtq+3rA9Y0F7ErQHdYfoY
jA2Nx4dKe5HMfdpk/89ZCD88QdEk/rzZfY7XJWyXbZ4cevlfUs36viKdm60CgYEP
qnFyVHuPclVlKeRgqagxc5vH33igMY7ZKe0Y8+HZNHPvNYAYa0p7ObQdrmWmtVdR
HM4RWMeISntL8uj+pNB+lBXyiT9f3gG6OD+VEsuBuGfpM9VSx2GXFx2obb0twgPR
iU8x5S9JpO3s229o2XJjbeEnTvDBr/KqS78N22QBvMkCgYAYppYCws/Ii3fEWF2N
2uDSIvVTbWeE3RZyA/kajdshnIaFc2fvsexN1Zz0Zk1yblUsqa9fmFS4zibCtAlx
sPnsU/Xifnn047Pb7Ja1sTd5Xd1bCTctyGFFoAFtqzpb9Nr6v6QjE1Ml9DqEnfZY
K0f0K7+WhDgWoWEj5SVCXES8VQKBgQxqwuUmRZHbxAgdfmGH2ELqKa3xWYFQBrPm
4YGHvZoWU1Zlh5TTZgPqJvPnybarwfwO4t8pCW7j6nq2nStJo+DQq9zEILFyHNhn
wS396cR2UBat+QZV9up1bhKUeQCN6czqExWvXR34VoYJIHNw95QMAgzQK1C6j5Of
2l23abte0QKBgQIXhnG/2+oJAFZeY61TiXYFlaF47UhzGeugivo0NHG9b8D8CTAO
WSbGswDZLE+hV974OjRrZyALOsChPa1ALzF8pJDK/o1RsYqSEyMZdAd4gRGaeCn+
WIm4OmWzxPpj9nr1xDUm4FpTaiGejnzHdc5L7CQ2bH8uPnorLF6o6SuROQ==
-----END RSA PRIVATE KEY-----`
)

func TestCreatePrivateKey(t *testing.T) {
	key, err := CreatePrivateKey("private.pem")
	assert.Nil(t, err)
	assert.NotEmpty(t, key.PublicKey())
}

func TestOpenKey(t *testing.T) {
	var err error
	key, err = OpenKey("private.pem", nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, key.Private)
	assert.NotEmpty(t, key.Public)

	_, err = OpenKey("missingPrivate.pem", nil)
	assert.Error(t, err)
}

func TestKeyFromData(t *testing.T) {
	key, err := KeyFromData(privateKey, nil)
	assert.Nil(t, err)
	assert.Equal(t, privateKey+"\n", key.PrivateKey())

	_, err = KeyFromData("----RSA INVALID---- 0x0x0x0", nil)
	assert.Error(t, err)
}

func TestEncrypt(t *testing.T) {
	var err error
	message, err = key.Encrypt("hello world")
	assert.Nil(t, err)
	assert.Equal(t, "hello world", message.Message)
	assert.NotEmpty(t, message.Signature)

	_, err = key.Encrypt("HsZt2r2UB2iQ2Nafv7sioavSnNfiPR82f7nIKbY1ZSAkYw9raE3SFCGguzBJehfKPe4lST73lILefdMpb77uP9FlweYTcKy9vWt2sSVZrXfr3vLYoUuY9Ivk8YYwd3c7zqIEnILGRbkGQZCSxwJFAIw0hvtvVJgC7xjYFYHE7KR82f7nIKbY1ZSAkYw9raE3SFCGguzBJehfKPe4lST73lILefdMpb77uP9FlweYTcKy9vWt2sSVZrXfr3vLYoUuY9Ivk8YYwd3c7zqIEnILGRbkGQZCSxwJFAIw0hvtvVJgC7xjY")
	assert.Error(t, err)
}

func TestToBase64(t *testing.T) {
	input := []byte("hello world")
	output := ToBase64(input)
	assert.Equal(t, "aGVsbG8gd29ybGQ=", output)
}

func TestMessage_JSON(t *testing.T) {
	jsonMessage := message.JSON()
	assert.Equal(t, 876, len(jsonMessage))

	var trueMessage *Message
	json.Unmarshal([]byte(jsonMessage), &trueMessage)
	assert.Equal(t, "hello world", trueMessage.Message)
}
