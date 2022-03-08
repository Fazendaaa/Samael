package samael

import (
	"fmt"

	"github.com/theckman/yacspin"
)

type ResponseChannel struct {
	Response error
	Process  string
}

func ConsumeChannel(params []string,
	spinner *yacspin.Spinner,
	resultChannel chan ResponseChannel) (fail error) {
	for i := 0; i < len(params); i++ {
		result := <-resultChannel

		spinner.Message(fmt.Sprintf("'%s'", result.Process))

		if nil != result.Response {
			return fmt.Errorf(`\n%v;
error while processing '%s'`, result.Response, result.Process)
		}
	}

	return fail
}
