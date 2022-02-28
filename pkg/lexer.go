package samael

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type projectFunc func(string, map[interface{}]interface{}) (interface{}, error)

// checkConfig looks for the config file
// It returns the file or an error
func checkConfig(filename string) (data []byte, fail error) {
	read, fail := ioutil.ReadFile(filename)

	if nil != fail {
		return nil, fail
	}

	return read, fail
}

// LexProject just reads the given Succubus config file and checks it whether
// or not it's a valid one.
// It returns whether or not the config file is valid and any error encountered.
func LexProject(defaultName string,
	projectPath string,
	interfaceToProject projectFunc) (
	project interface{},
	fail error) {
	filename, fail := readFileOrDir(defaultName, projectPath)

	if nil != fail {
		return nil, fail
	}

	data, fail := checkConfig(filename)

	if nil != fail {
		return project, fail
	}

	read := make(map[interface{}]interface{})
	fail = yaml.Unmarshal([]byte(data), &read)

	if nil != fail {
		return project, fail
	}

	return interfaceToProject(filename, read)
}
