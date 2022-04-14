package finnhub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"exporter/internal/entities"
	"github.com/D3vR4pt0rs/logger"
	"github.com/gorilla/websocket"
)

const (
	FINNHUB_WEBSOCKET_URL = "wss://ws.finnhub.io"
	FINNHUB_API_URL       = "https://finnhub.io/api/v1/"
)

type client struct {
	token           string
	httpClient      *http.Client
	websocketClient *websocket.Dialer
	currentData     map[string]map[time.Time]CurrentStockValue
}

type CurrentStockMap map[time.Time]CurrentStockValue

func New(token string) *client {
	return &client{
		token:           token,
		httpClient:      &http.Client{},
		websocketClient: &websocket.Dialer{},
		currentData:     make(map[string]map[time.Time]CurrentStockValue),
	}
}

func (c client) OpenWebSocketConnection(symbols []string) {
	url := fmt.Sprintf("%s?token=%s", FINNHUB_WEBSOCKET_URL, c.token)
	logger.Info.Println("Open connection with Finnhub by websocket")
	w, _, err := c.websocketClient.Dial(url, nil)
	if err != nil {
		logger.Error.Println(err.Error())
	}
	defer w.Close()

	for _, symbol := range symbols {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": symbol})
		w.WriteMessage(websocket.TextMessage, msg)
	}

	var msg Message
	for {
		logger.Info.Println("Collecting information from channel")
		err := w.ReadJSON(&msg)
		logger.Info.Println(msg)
		if err != nil {
			logger.Error.Println(err.Error())
		}
		for _, result := range msg.Data {
			currentStock := CurrentStockValue{result.LastPrice, result.Volume}
			timeStamp := c.parseTimestamp(result.Timestamp)
			c.currentData[result.Symbol][timeStamp] = currentStock
			logger.Info.Println(c.currentData[result.Symbol])
		}

		time.Sleep(1 * time.Minute)
	}
}

func (c client) GetQuoteStock(ticker string) (entities.Stock, error) {
	url := fmt.Sprintf("%squote?symbol=%s", FINNHUB_API_URL, ticker)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Error.Println(err.Error())
		return entities.Stock{}, err
	}
	request.Header.Add("X-Finnhub-Token", c.token)

	resp, err := c.httpClient.Do(request)
	if err != nil {
		logger.Error.Println(err.Error())
		return entities.Stock{}, err
	}

	logger.Info.Println(url, resp.StatusCode)
	defer resp.Body.Close()

	var result QuoteStock
	json.NewDecoder(resp.Body).Decode(&result)

	logger.Info.Println(result)

	return entities.Stock{
		Symbol:    ticker,
		LastPrice: result.CurrentPrice,
		Timestamp: time.Unix(result.Timestamp, 0),
	}, nil
}

func (c client) parseTimestamp(timestamp int64) time.Time {
	tUnixNanoRemainder := (timestamp % int64(time.Microsecond)) * int64(time.Millisecond)
	return time.Unix(timestamp/int64(time.Microsecond), tUnixNanoRemainder)
}
