package handle

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
)

type ExtendedError func(model.ExtendedError)

type Error func(error)
