package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type SpotifyConnectionResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

func SpotifyConnection(spotifyClientId string, spotifyClientSecret string) (token SpotifyConnectionResponse, err error) {
	form := url.Values{"grant_type": {"client_credentials"}}
	request, err := http.PostForm("https://accounts.spotify.com/api/token", form)
	if err != nil {
		return
	}
	something := fmt.Sprintf("%s:%s", spotifyClientId, spotifyClientSecret)
	// authorizationKey := base64.StdEncoding.EncodeToString([]byte(something))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", something)

	defer request.Body.Close()

	responseData, err := io.ReadAll(request.Body)
	if err != nil {
		return
	}

	data := SpotifyConnectionResponse{}
	if err = json.Unmarshal(responseData, &data); err != nil {
		return
	}

	return data, nil
}
