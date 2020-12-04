package input

import (
	"io/ioutil"
)

func ReadAll(fname string) (string, error) {
	data, err := ioutil.ReadFile(fname)

	return string(data), err
}
