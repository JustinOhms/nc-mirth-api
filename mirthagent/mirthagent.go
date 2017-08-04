package mirthagent

import (
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
