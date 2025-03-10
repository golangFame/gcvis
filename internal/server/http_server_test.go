package server

import (
	"bytes"
	"encoding/json"
	"github.com/golangFame/gcvis/pkg/graph"
	"github.com/golangFame/gcvis/pkg/trace"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestHttpServerListener(t *testing.T) {
	graph := graph.NewGraph("fake title", graph.GCVIS_TMPL)
	server := NewHttpServer("127.0.0.1", "0", &graph)

	url := server.Url()

	if !strings.Contains(url, "http://127.0.0.1") {
		t.Fatalf("Server URL didn't contain localhost address: %v", url)
	}
}

func TestHttpServerResponse(t *testing.T) {
	graph := graph.NewGraph("fake title", graph.GCVIS_TMPL)
	graph.AddGCTraceGraphPoint(&trace.Gctrace{})
	server := NewHttpServer("127.0.0.1", "0", &graph)

	go server.Start()
	defer server.Close()

	response, err := http.Get(server.Url())
	if err != nil {
		t.Errorf("HTTP request returned an error: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error while reading response body: %v", err)
	}

	w := &bytes.Buffer{}

	if err = graph.Write(w); err != nil {
		t.Errorf("Error while writing template: %v", err)
	}

	expectedBody, err := io.ReadAll(w)
	if err != nil {
		t.Errorf("Error while reading buffer: %v", err)
	}

	if !bytes.Equal(expectedBody, body) {
		t.Fatalf(
			"Expected response body to equal parsed template.\nExpected: %v\nGot: %v",
			string(expectedBody),
			string(body),
		)
	}
}

func TestHttpServerJsonEndpoint(t *testing.T) {
	graph := graph.NewGraph("fake title", graph.GCVIS_TMPL)
	graph.AddGCTraceGraphPoint(&trace.Gctrace{Heap1: 10})
	server := NewHttpServer("127.0.0.1", "0", &graph)

	go server.Start()
	defer server.Close()

	/*jsonEndpoint, err := url.JoinPath(server.Url(), "graph")

	if err != nil {
		t.Fatal("error building jsonEndpoint")
	}
	*/
	jsonEndpoint := server.Url() + "graph"
	response, err := http.Get(jsonEndpoint)

	if err != nil {
		t.Errorf("HTTP request returned an error: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error while reading response body: %v", err)
	}

	result, err := json.Marshal(graph)
	if err != nil {
		t.Errorf("Error marshalling graph: %v", err)
	}

	if string(result) != strings.TrimRight(string(body), "\r\n") {
		t.Errorf("Expected graph to be a json string.\nExpected: %v\nGot: %v", string(result), string(body))
	}
}
