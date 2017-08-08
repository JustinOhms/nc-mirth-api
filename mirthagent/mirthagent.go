package mirthagent

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/handle"
	"github.com/caimeo/stickyjar/tracer"
	"github.com/parnurzeal/gorequest"
)

var Tracer tracer.Tracer

func traceCurl(r *gorequest.SuperAgent) {
	if Tracer.IsVerbose() {
		cmd, _ := r.AsCurlCommand()
		Tracer.Verbose(cmd)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func errorCheck(onErr handle.Error, errs []error, text string) bool {
	if len(errs) > 0 {
		onErr(*model.NewRequestError(text, errs))
		return true
	}
	return false
}
