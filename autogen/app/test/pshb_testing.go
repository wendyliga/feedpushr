// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "feedpushr": pshb TestHelpers
//
// Command:
// $ goagen
// --design=github.com/ncarlier/feedpushr/v2/design
// --out=/home/nicolas/workspace/fe/feedpushr/autogen
// --version=v1.4.3

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/ncarlier/feedpushr/v2/autogen/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

// PubPshbBadRequest runs the method Pub of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func PubPshbBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PshbController) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v2/pshb"),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PshbTest"), rw, req, prms)
	pubCtx, _err := app.NewPubPshbContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.Pub(pubCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// PubPshbOK runs the method Pub of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func PubPshbOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PshbController) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v2/pshb"),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PshbTest"), rw, req, prms)
	pubCtx, _err := app.NewPubPshbContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.Pub(pubCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}

// SubPshbBadRequest runs the method Sub of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func SubPshbBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PshbController, hubChallenge string, hubLeaseSeconds *int, hubMode string, hubTopic string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{hubChallenge}
		query["hub.challenge"] = sliceVal
	}
	if hubLeaseSeconds != nil {
		sliceVal := []string{strconv.Itoa(*hubLeaseSeconds)}
		query["hub.lease_seconds"] = sliceVal
	}
	{
		sliceVal := []string{hubMode}
		query["hub.mode"] = sliceVal
	}
	{
		sliceVal := []string{hubTopic}
		query["hub.topic"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v2/pshb"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{hubChallenge}
		prms["hub.challenge"] = sliceVal
	}
	if hubLeaseSeconds != nil {
		sliceVal := []string{strconv.Itoa(*hubLeaseSeconds)}
		prms["hub.lease_seconds"] = sliceVal
	}
	{
		sliceVal := []string{hubMode}
		prms["hub.mode"] = sliceVal
	}
	{
		sliceVal := []string{hubTopic}
		prms["hub.topic"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PshbTest"), rw, req, prms)
	subCtx, _err := app.NewSubPshbContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.Sub(subCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// SubPshbOK runs the method Sub of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func SubPshbOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PshbController, hubChallenge string, hubLeaseSeconds *int, hubMode string, hubTopic string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{hubChallenge}
		query["hub.challenge"] = sliceVal
	}
	if hubLeaseSeconds != nil {
		sliceVal := []string{strconv.Itoa(*hubLeaseSeconds)}
		query["hub.lease_seconds"] = sliceVal
	}
	{
		sliceVal := []string{hubMode}
		query["hub.mode"] = sliceVal
	}
	{
		sliceVal := []string{hubTopic}
		query["hub.topic"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v2/pshb"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{hubChallenge}
		prms["hub.challenge"] = sliceVal
	}
	if hubLeaseSeconds != nil {
		sliceVal := []string{strconv.Itoa(*hubLeaseSeconds)}
		prms["hub.lease_seconds"] = sliceVal
	}
	{
		sliceVal := []string{hubMode}
		prms["hub.mode"] = sliceVal
	}
	{
		sliceVal := []string{hubTopic}
		prms["hub.topic"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PshbTest"), rw, req, prms)
	subCtx, _err := app.NewSubPshbContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.Sub(subCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}
