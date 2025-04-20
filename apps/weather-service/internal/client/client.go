package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Porsche-ths/weather-4-u/weather-service/internal/config"
	"github.com/Porsche-ths/weather-4-u/weather-service/internal/werror"
	"github.com/Porsche-ths/weather-4-u/weather-service/internal/dto"
)

type Client struct {
	url *url.URL
}

func NewClient(cfg *config.Config) (*Client, error) {
	url, err := url.Parse(cfg.WeatherApi.Endpoint)
	if err != nil {
		return nil, err
	}

	url.RawQuery = fmt.Sprintf("key=%s", cfg.WeatherApi.Key)

	return &Client{
		url,
	}, nil
}

func (c *Client) parseError(res *http.Response) error {
	var errRes dto.ErrorResponse
	err := json.NewDecoder(res.Body).Decode(&errRes)
	if err != nil {
		return err
	}
	return werror.New(errRes.Message, werror.WithCode(res.StatusCode), werror.WithExternalSource())
}

func (c *Client) FetchWeather(ctx context.Context, q string, days string) (*dto.WeatherResponse, error) {
	url := c.url

	query := url.Query()
	query.Set("q", q)
	query.Set("days", days)
	url.RawQuery = query.Encode()
	fmt.Printf("url: %s\n", url.String())

	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, c.parseError(res)
	}

	var weatherRes dto.WeatherResponse
	err = json.NewDecoder(res.Body).Decode(&weatherRes)
	if err != nil {
		return nil, err
	}

	return &weatherRes, nil
}

// func (c *Client) SearchFromAddress(ctx context.Context, text string) (*dto.GeocodeResponse, error) {
// 	url := c.url.JoinPath("/v1/geocode/search")

// 	query := url.Query()
// 	query.Set("text", text)
// 	url.RawQuery = query.Encode()

// 	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		return nil, c.parseError(res)
// 	}

// 	var geocodeRes dto.GeocodeResponse
// 	err = json.NewDecoder(res.Body).Decode(&geocodeRes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &geocodeRes, nil
// }

// func (c *Client) CoordinateFromIP(ctx context.Context, ip string) (*dto.GeoipResponse, error) {
// 	url := c.url.JoinPath("/v1/ipinfo")

// 	query := url.Query()
// 	query.Set("ip", ip)
// 	query.Set("lang", "en")
// 	url.RawQuery = query.Encode()

// 	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		return nil, c.parseError(res)
// 	}

// 	var geoipRes dto.GeoipResponse
// 	err = json.NewDecoder(res.Body).Decode(&geoipRes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &geoipRes, nil
// }
