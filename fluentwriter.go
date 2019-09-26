package fluentwriter

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type FluentWriter struct {
	input chan []byte
	baseURL string
	client *http.Client
}

func NewFluentWriter(host string, port int, tag string, timeout time.Duration, bufferSize int) *FluentWriter {
	w := &FluentWriter{
		input:   make(chan []byte, bufferSize),
		baseURL: fmt.Sprintf("http://%s:%d/%s", host, port, tag),
		client: &http.Client{
			Timeout: timeout,
		},
	}
	go func() {
		for {
			w.doWrite(<-w.input)
		}
	}()
	return w
}

func (w *FluentWriter) doWrite(p []byte) {
	req, err := http.NewRequest("GET", w.baseURL, nil)
	if err != nil {
		return
	}
	q := req.URL.Query()
	q.Add("json", fmt.Sprintf("%s", p))
	req.URL.RawQuery = q.Encode()

	_, err = w.client.Do(req)
}

func (w *FluentWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	select {
	case w.input <- p:
	default:
		err = errors.New("log buffer is full")
		n = 0
	}
	return
}