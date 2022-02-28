package samael

import (
	"testing"
)

func TestReadFileOrDir(t *testing.T) {
	t.Run("missing", func(t *testing.T) {
		value, fail := readFileOrDir("shojo", "../test/config/missing")

		if nil == fail {
			t.Errorf("Expected to got a invalid error but got the following value:\n%v", value)
		}
	})
}
