package errors

import (
	"bytes"
	"errors"

	"github.com/caimeo/console"

	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
	"github.com/parnurzeal/gorequest"
)

var Console console.Console
var commonErrorChannel chan error

func TraceCurl(r *gorequest.SuperAgent) {
	if Console.IsVerbose() {
		cmd, _ := r.AsCurlCommand()
		Console.Verbose(cmd)
	}
}

func CheckErrorAndPanic(e error) {
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

func CheckErrorAndLog(e error) {
	if e == nil || commonErrorChannel == nil {
		return
	}
	select {
	case commonErrorChannel <- e:
	default: //common error channel is full, use this so we don't block
	}
}

func CheckErrorAndChannelLog(e error, ec chan error) {
	if e == nil {
		return
	}
	CheckErrorAndLog(e)
	select {
	case ec <- e:
	default: //channel is full or otherwise unavailable, so we don't block
	}
}

func ResponseErrors(ec chan error, errs []error, text string) bool {
	if len(errs) > 0 {
		e := *model.NewExtendedError(text, errs)
		CheckErrorAndChannelLog(e, ec)
		return true
	}
	return false
}

func StatusErrors(ec chan error, r gorequest.Response, text string) bool {
	if r.StatusCode > 299 || r.StatusCode < 200 {
		var ea []error
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		ea = append(ea, errors.New(buf.String()))
		e := *model.NewExtendedError(text, ea)
		CheckErrorAndChannelLog(e, ec)
		return true
	}
	return false
}

func ResponseOrStatusErrors(ec chan error, r gorequest.Response, errs []error, text string) bool {
	return ResponseErrors(ec, errs, text) || StatusErrors(ec, r, text)
}
