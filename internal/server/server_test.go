package server

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	// req, err := http.NewRequest("GET", "http://localhost:3000/ws", nil)
	// if err != nil {
	// 	t.Fatalf("could create a request for ws on %s: %v", url, err)
	// }

	// req.SetBasicAuth(username, password)

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func TestWebsocket(t *testing.T) {
	server := httptest.NewServer(NewSmartParkingIotServer())
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
	ws := mustDialWS(t, wsURL)
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			return
		}
	}(ws)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
