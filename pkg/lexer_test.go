package samael

import "testing"

func interfaceToProject(filename string, read map[interface{}]interface{}) (result interface{}, fail error) {
	return result, fail
}

func TestLex(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		value, fail := LexProject("samael", "../test/config/default/", interfaceToProject)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})
}
