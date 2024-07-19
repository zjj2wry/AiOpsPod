package tools

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
)

// StartMockPrometheusServer starts a mock Prometheus server on the specified port.
func StartMockPrometheusServer(port int) (*httptest.Server, error) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": "success",
			"data": {
				"resultType": "matrix",
				"result": [
					{
						"metric": {
							"__name__": "up",
							"instance": "localhost:9090",
							"job": "prometheus"
						},
						"values": [
							[1609459200, "1"],
							[1609459260, "1"]
						]
					}
				]
			}
		}`))
	})

	server := httptest.NewServer(handler)

	// Change the port if necessary
	if port != 0 {
		server.Listener.Close()
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			return nil, err
		}
		server.Listener = listener
	}

	return server, nil
}
