package messaging

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type HTTPRequester interface {
	SendRequest() ([]byte, error)
}

type RestyClient struct {
	QueryParams map[string]string
	UrlValues   url.Values
	Url         string
}

func (h *RestyClient) SendRequest() ([]byte, error) {
	client := resty.New()
	log.Info().Msgf("sending new request to %v", h.Url)
	resp, err := client.R().
		EnableTrace().
		SetQueryParams(map[string]string{
			"latitude":  "-21.31",
			"longitude": "-157.85",
			"timezone":  "GMT",
		}).
		SetQueryParamsFromValues(url.Values{
			"hourly": []string{"temperature_2m", "showers"},
		}).
		Get(h.Url)
	if err != nil {
		return resp.Body(), fmt.Errorf("error during request to %v: %v", h.Url, err)
	}
	if statusCode := resp.StatusCode(); statusCode != http.StatusOK {
		return resp.Body(), fmt.Errorf("response status code doesn't equal 200: %v", statusCode)
	}
	log.Info().Msgf("request to %v was succesful", resp.Request.URL)
	return resp.Body(), err
}
