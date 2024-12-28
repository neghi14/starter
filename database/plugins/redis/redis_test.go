package redisdb

import "testing"

func TestRedisDB(t *testing.T) {
	type TestModel struct{}
	_, err := New(Opts(), TestModel{})
	if err != nil {
		t.Log(err)
	}

	/// future test here
}
