package result

import (
	"github.com/pkg/errors"
)

func Exception(result *Result) error {
	return errors.Wrap(result, result.Message)
}
