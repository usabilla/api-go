package gobilla

import (
	"testing"

	"github.com/usabilla/gobilla/test"
)

func Test_Hash(t *testing.T) {
	spec := test.Spec(t)
	hashData := string(hash([]byte("test")))
	expected := string([]byte{159, 134, 208, 129, 136, 76, 125, 101, 154, 47, 234, 160, 197, 90, 208, 21, 163, 191, 79, 27, 43, 11, 130, 44, 209, 93, 108, 21, 176, 240, 10, 8})
	spec.Expect(hashData).ToEqual(expected)
}

func Test_HexHash(t *testing.T) {
	spec := test.Spec(t)
	hashData := hexHash([]byte("test"))
	expected := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	spec.Expect(hashData).ToEqual(expected)
}

func Test_KeyedHash(t *testing.T) {
	spec := test.Spec(t)
	keyedHashData := string(keyedHash([]byte("key"), []byte("test")))
	expected := string([]byte{2, 175, 181, 99, 4, 144, 44, 101, 111, 203, 115, 124, 221, 3, 222, 98, 5, 187, 109, 64, 29, 162, 129, 46, 253, 155, 45, 54, 160, 138, 241, 89})
	spec.Expect(keyedHashData).ToEqual(expected)
}

func Test_HexKeyedHash(t *testing.T) {
	spec := test.Spec(t)
	hashData := hexKeyedHash([]byte("key"), []byte("test"))
	expected := "02afb56304902c656fcb737cdd03de6205bb6d401da2812efd9b2d36a08af159"
	spec.Expect(hashData).ToEqual(expected)
}
