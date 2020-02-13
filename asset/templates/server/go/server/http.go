package server

import (
	"encoding/json"
	"fmt"
	"git.intra.longguikeji.com/longguikeji/arkfbp-go/flow"
	"net/http"
)

// HTTPServer ...
type HTTPServer struct {
}

// Serve ...
func (svr *HTTPServer) Serve(host string, port int) error {
	s := fmt.Sprintf("%s:%d", host, port)
	err := http.ListenAndServe(s, nil)
	if err != nil {
		return err
	}

	return nil
}

// RegisterRoutes ...
func (svr *HTTPServer) RegisterRoutes(handlers map[string]flow.IFlow) error {
	for pattern, handler := range handlers {
		http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
			outputs := handler.Run()

			if outputs != nil {
				data, err := json.Marshal(outputs)
				if err != nil {
					w.WriteHeader(500)
					w.Write([]byte(fmt.Sprintf("%v", err)))
					return
				}

				w.Header().Add("Content-Type", "application/json")
				w.Write(data)
			} else {
				w.Write([]byte(""))
			}
		})
	}

	return nil
}

func process(w http.ResponseWriter, req *http.Request) {

}

// NewHTTPServer ...
func NewHTTPServer() *HTTPServer {
	s := &HTTPServer{}
	return s
}
