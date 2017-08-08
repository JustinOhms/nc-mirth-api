package mirthagent

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/errorhandler"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/extendederror"
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

func errorCheck(onErr errorhandler.Handler, errs []error, text string) bool {
	if len(errs) > 0 {
		onErr(*extendederror.New(text, errs))
		return true
	}
	return false
}
