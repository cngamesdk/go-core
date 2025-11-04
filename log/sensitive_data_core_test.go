package log

import "testing"

func TestReplaceStrSensitiveData(t *testing.T) {
	str := "{\"timestamp\":123,\"game_id\":1,\"sign\":\"2086d7e79b21ee4d148d872bb648b867\",\"user_name\":\"test\", \"password\":\"123456\"}"
	println(ReplaceStrSensitiveData(str, []byte("1234567890123456")))
}
