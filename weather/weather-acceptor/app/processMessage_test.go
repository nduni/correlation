package app

import (
	"context"
	"testing"

	"github.com/nduni/correlation/common/messaging"
	"github.com/nduni/correlation/weather/weather-acceptor/app/resources"
)

type httpClientMock struct {
	resp []byte
}

func (h *httpClientMock) SendRequest() ([]byte, error) {
	return h.resp, nil
}

type senderMock struct {
	sentMsg string
}

func (s *senderMock) Send(ctx context.Context, msg []byte) error {
	s.sentMsg = string(msg)
	return nil
}

func TestProcessWeather(t *testing.T) {
	tests := []struct {
		name           string
		weatherResp    string
		httpClientMock *httpClientMock
		senderMock     *senderMock
		expectedMsg    string
		wantErr        bool
	}{
		{
			name:           "empty api response",
			weatherResp:    "",
			httpClientMock: &httpClientMock{},
			senderMock:     &senderMock{},
			expectedMsg:    "",
			wantErr:        true,
		},
		{
			name:           "valid api response",
			weatherResp:    resources.WeatherResp,
			httpClientMock: &httpClientMock{},
			senderMock:     &senderMock{},
			expectedMsg:    resources.BrokerMessage,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// init http client mock
			tt.httpClientMock.resp = []byte(tt.weatherResp)
			HTTPClient = tt.httpClientMock

			// init sender senderMock
			Senders = make(map[string]messaging.Sender)
			Senders[TOPIC_WEATHER] = tt.senderMock

			if err := ProcessWeather(context.Background()); (err != nil) != tt.wantErr {
				t.Errorf("ProcessWeather() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.expectedMsg != tt.senderMock.sentMsg {
				t.Fatalf("expected message sent to message broker: %v, got %v", tt.expectedMsg, tt.senderMock.sentMsg)
			}
		})
	}
}
