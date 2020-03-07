package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	mapset "github.com/deckarep/golang-set"
	"github.com/longguikeji/arkfbp-go/intr"
)

// Route ...
type Route struct {
	Name    string
	Pattern string
	Handler intr.IFlow
}

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
func (svr *HTTPServer) RegisterRoutes(routes []Route) error {
	names := mapset.NewSet()

	for _, r := range routes {
		if names.Contains(r.Name) {
			return fmt.Errorf("duplicated route name: %s", r.Name)
		}
		http.HandleFunc(r.Pattern, func(w http.ResponseWriter, req *http.Request) {
			outputs := r.Handler.Run(nil)

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

// NewHTTPServer ...
func NewHTTPServer() *HTTPServer {
	s := &HTTPServer{}
	return s
}
