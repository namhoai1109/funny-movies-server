package link

import (
	"context"
	"encoding/json"
	"fmt"
	"funnymovies/internal/model"
	"funnymovies/util/server"
	websocketutil "funnymovies/util/websocket"
	"io"
	"net/http"
)

func (s *Link) Create(ctx context.Context, user *model.AuthoUser, url string) error {
	if err := s.linkRepository.Create(s.db, &model.Link{Url: url, UserID: user.ID}); err != nil {
		server.NewHTTPInternalError(fmt.Sprintf("Error when create link: %v", err.Error()))
	}

	title, err := s.getTitle(url)
	if err != nil {
		return err
	}

	s.ws.BroadcastMessage(websocketutil.Message{
		EmailSender: user.Email,
		VideoTitle:  title,
	})

	return nil
}

func (s *Link) getTitle(url string) (string, error) {
	reqUrl := fmt.Sprintf(`https://www.youtube.com/oembed?url=%s&format=json`, url)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
	}
	linkData := &YtbOembedResponse{}
	if err := json.Unmarshal(body, linkData); err != nil {
		return "", err
	}
	return linkData.Title, nil
}
