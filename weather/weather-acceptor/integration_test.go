package main

import (
	"context"
	"errors"
	"testing"

	"github.com/nduni/correlation/common/messaging"
	"github.com/nduni/correlation/weather/weather-acceptor/app"
)

type htttpClientMock struct {
	resp []byte
}

func (h *htttpClientMock) SendRequest() ([]byte, error) {
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
		name            string
		weatherResp     string
		htttpClientMock *htttpClientMock
		senderMock      *senderMock
		expetedMsg      string
		wantErr         error
	}{
		{
			name:            "invalid_weather_message_process",
			weatherResp:     "",
			htttpClientMock: &htttpClientMock{},
			senderMock:      &senderMock{},
			expetedMsg:      "",
			wantErr:         errors.New("empty response body"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// init http client mock
			tt.htttpClientMock.resp = []byte(tt.weatherResp)
			app.HTTPClient = tt.htttpClientMock

			// init sender mock
			app.Senders = make(map[string]messaging.Sender)
			app.Senders[app.TOPIC_WEATHER] = tt.senderMock

			// run tested function
			err := app.ProcessWeather(context.Background())

			// asserstions
			if err.Error() != tt.wantErr.Error() {
				t.Fatalf("TestProcessWeather wanted error: %v, but got: %v", tt.wantErr, err)
			}
			if tt.expetedMsg != tt.senderMock.sentMsg {
				t.Fatalf("TestProcessWeather expected messega to be sent to message broker: %v, got %v", tt.expetedMsg, tt.senderMock.sentMsg)
			}
		})
	}
}
