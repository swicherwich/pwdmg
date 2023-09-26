package secutil

import "testing"

func TestHashStr(t *testing.T) {
	input := "hello"
	expectedHash := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"

	hash := HashStr(input)

	if hash != expectedHash {
		t.Errorf("HashStr(%s) = %s; want %s", input, hash, expectedHash)
	}
}

func TestEncodeDecodeBase64(t *testing.T) {
	input := "hello, world!"
	encoded := EncodeBase64(input)

	decoded, err := DecodeBase64(encoded)
	if err != nil {
		t.Errorf("DecodeBase64(%s) returned an error: %v", encoded, err)
	}

	if decoded != input {
		t.Errorf("EncodeBase64/DecodeBase64(%s) = %s; want %s", input, decoded, input)
	}
}
