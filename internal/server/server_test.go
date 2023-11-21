package server

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/thelemonday/smart-parking-iot-server/database"
	"github.com/thelemonday/smart-parking-iot-server/internal/server/presenter"
)

var (
	username = "02110882985"
	password = "G_2tEipy9ldDoqn"
	// url      = "ws://localhost:8080/ws"
)

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	// req, err := http.NewRequest("GET", "http://localhost:3000/ws", nil)
	// if err != nil {
	// 	t.Fatalf("could create a request for ws on %s: %v", url, err)
	// }

	// req.SetBasicAuth(username, password)

	h := http.Header{"Authorization": {"Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))}}
	ws, _, err := websocket.DefaultDialer.Dial(url, h)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func mustMakeSmartParkingIotServer() *SmartParkingIotService {
	accountsDatabase := database.NewAccountsDatabase(&database.CurrentCarParkingStatusDatabase{})
	return NewSmartParkingIotServer(accountsDatabase)
}

func TestWebsocket(t *testing.T) {
	server := httptest.NewServer(mustMakeSmartParkingIotServer())
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
	ws := mustDialWS(t, wsURL)
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			return
		}
	}(ws)

	var account presenter.AccountAuthenticationSuccessResponse
	if err := ws.ReadJSON(&account); err != nil {
		t.Fatalf("Cannot read account json data")
	}

	t.Log(account)

	t.Run("websocket send first message if authentication success", func(t *testing.T) {
		assertCorrectMessage(t, account.Type, "authentication")

		assertCorrectMessage(t, account.Username, username)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
