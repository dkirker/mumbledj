/*
 * MumbleDJ
 * By Matthieu Grieger
 * services/mixcloud.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package services

import (
	"fmt"
	//"net/http"
	"regexp"
	"strings"
	"time"

	//"github.com/antonholmquist/jason"
	"layeh.com/gumble/gumble"
	"go.reik.pl/mumbledj/bot"
	"go.reik.pl/mumbledj/interfaces"
)

// Icecast is a wrapper around the Icecast API.
// https://www.mixcloud.com/developers/
type Icecast struct {
	*GenericService
}

// NewIcecastService returns an initialized Icecast service object.
func NewIcecastService() *Icecast {
	return &Icecast{
		&GenericService{
			ReadableName: "Icecast",
			Format:       "unknown",
			MediaSource:  "stream",
			TrackRegex: []*regexp.Regexp{
				regexp.MustCompile(`^https?:.*`),
			},
			// Playlists are currently unsupported
			PlaylistRegex: nil,
		},
	}
}

// CheckAPIKey performs a test API call with the API key
// provided in the configuration file to determine if the
// service should be enabled.
func (mc *Icecast) CheckAPIKey() error {
	// Icecast (at the moment) does not require an API key,
	// so we can just return nil.
	return nil
}

// GetTracks uses the passed URL to find and return
// tracks associated with the URL. An error is returned
// if any error occurs during the API call.
func (mc *Icecast) GetTracks(url string, submitter *gumble.User) ([]interfaces.Track, error) {
	var (
		//apiURL string
		//err    error
		//resp   *http.Response
		//v      *jason.Object
		tracks []interfaces.Track
	)

	//apiURL = strings.Replace(url, "www", "api", 1)

	// Track playback offset is not present in Icecast URLs,
	// so we can safely assume that users will not request
	// a playback offset in the URL.
	offset, _ := time.ParseDuration("0s")

	//resp, err = http.Get(apiURL)
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()

	//v, err = jason.NewObjectFromReader(resp.Body)
	//if err != nil {
	//	return nil, err
	//}

	//id, _ := v.GetString("slug")
	//trackURL, _ := v.GetString("url")
	//title, _ := v.GetString("name")
	//author, _ := v.GetString("user", "username")
	//authorURL, _ := v.GetString("user", "url")
	//durationSecs, _ := v.GetInt64("audio_length")
	//duration, _ := time.ParseDuration(fmt.Sprintf("%ds", durationSecs))
	//thumbnail, err := v.GetString("pictures", "large")
	//if err != nil {
	//	// Track has no artwork, using profile avatar instead.
	//	thumbnail, _ = v.GetString("user", "pictures", "large")
	//}

	urlParts := strings.Split(url, "/")
	icecastMountPoint := urlParts[len(urlParts)-1]

	id := icecastMountPoint
	title := icecastMountPoint
	author := ""
	authorURL := url
	durationSecs := -1
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", durationSecs))
	thumbnail := ""

	track := bot.Track{
		ID:             id,
		URL:            url,
		Title:          title,
		Author:         author,
		AuthorURL:      authorURL,
		Submitter:      submitter.Name,
		Service:        mc.ReadableName,
		ThumbnailURL:   thumbnail,
		Filename:       id + ".track",
		Duration:       duration,
		PlaybackOffset: offset,
		Playlist:       nil,
	}

	tracks = append(tracks, track)

	return tracks, nil
}
