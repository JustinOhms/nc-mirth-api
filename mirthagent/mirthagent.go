package mirthagent

import (
	"bytes"
	"errors"

	"github.com/NavigatingCancer/mirth-api/mirthagent/model/extendederror"
	"github.com/caimeo/stickyjar/tracer"
	"github.com/parnurzeal/gorequest"
)

var Tracer tracer.Tracer
var commonErrorChannel chan error

func traceCurl(r *gorequest.SuperAgent) {
	if Tracer.IsVerbose() {
		cmd, _ := r.AsCurlCommand()
		Tracer.Verbose(cmd)
	}
}

func checkErrorAndPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func CommonErrorChannel() chan error {
	if commonErrorChannel == nil {
		commonErrorChannel = make(chan error, 100)
	}
	return commonErrorChannel
}

func checkErrorAndLog(e error) {
	if e == nil || commonErrorChannel == nil {
		return
	}
	select {
	case commonErrorChannel <- e:
	default: //common error channel is full, use this so we don't block
	}
}

func checkErrorAndChannelLog(e error, ec chan error) {
	if e == nil {
		return
	}
	checkErrorAndLog(e)
	select {
	case ec <- e:
	default: //channel is full or otherwise unavailable, so we don't block
	}
}

func responseErrors(ec chan error, errs []error, text string) bool {
	if len(errs) > 0 {
		e := *extendederror.New(text, errs)
		checkErrorAndChannelLog(e, ec)
		return true
	}
	return false
}

func statusErrors(ec chan error, r gorequest.Response, text string) bool {
	if r.StatusCode != 200 {
		var ea []error
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		ea = append(ea, errors.New(buf.String()))
		e := *extendederror.New(text, ea)
		checkErrorAndChannelLog(e, ec)
		return true
	}
	return false
}

func responseOrStatusErrors(ec chan error, r gorequest.Response, errs []error, text string) bool {
	return responseErrors(ec, errs, text) || statusErrors(ec, r, text)
}
