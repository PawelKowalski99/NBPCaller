package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"
)

const (
	nbpEurLast100URL = "http://api.nbp.pl/api/exchangerates/rates/a/eur/last/100/?format=json"
)

func NewHost(writer io.Writer) *Handler {
	return &Handler{W: writer}
}

// Makes API Call
type Checker interface {
	Check(host string, x, y int) error
}

type Handler struct {
	W io.Writer
}

type HostResponse struct {
	Message     string
	Host        string
	Method      string
	TFFBTime    time.Duration
	StatusCode  int
	ContentType string
	Body        []byte
}

type CallResponse struct {
	Message     string
	Host        string
	Method      string
	TFFBTime    time.Duration
	StatusCode  int
	ContentType string
}

func (h *Handler) CheckHost(host string, x, y int) error {

	for i := 0; i < x; i++ {
		hResp, err := h.get(host)
		if err != nil {
			return err
		}

		cResp := CallResponse{
			Message:     hResp.Message,
			Host:        hResp.Host,
			Method:      hResp.Method,
			TFFBTime:    hResp.TFFBTime,
			StatusCode:  hResp.StatusCode,
			ContentType: hResp.ContentType,
		}

		b, err := json.Marshal(&cResp)
		if err != nil {
			return err
		}

		b = append(b, "\n"...)

		// Write json marshaled information via writer
		_, err = h.W.Write(b)
		if err != nil {
			return fmt.Errorf("could not write to host writer: %w", err)

		}

		time.Sleep(time.Duration(y) * time.Second)

	}

	return nil
}

func (h *Handler) get(host string) (*HostResponse, error) {
	req, err := http.NewRequest(http.MethodGet, host, nil)
	if err != nil {
		return nil, err
	}

	var start time.Time
	var TFFBTime time.Duration

	// Prepare ClientTrace to know TFFBTime
	trace := &httptrace.ClientTrace{

		GotFirstResponseByte: func() {
			TFFBTime = time.Since(start)
			fmt.Printf("Time from start to first byte: %v\n", time.Since(start))
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	// Set User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:10.0) Gecko/20100101 Firefox/10.0")

	start = time.Now()
	// Call host
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Check for status code and behave accordingly
	switch s := resp.StatusCode; s {
	case http.StatusOK:

	case http.StatusNoContent:

	case http.StatusNotFound:

	case http.StatusInternalServerError:

	}

	ct := resp.Header.Get("Content-Type")

	// Check if content-type is application/json; charset=utf-8
	if !strings.Contains("application/json; charset=utf-8", ct) {

		return nil, fmt.Errorf("response content-type is not json, content-type: %s", ct)
	}

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// close response body
	err = resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("could not close body: %w", err)
	}

	data := HostResponse{
		Host:        host,
		Method:      req.Method,
		TFFBTime:    TFFBTime,
		StatusCode:  resp.StatusCode,
		ContentType: ct,
		Body:        body,
	}

	return &data, nil

}
