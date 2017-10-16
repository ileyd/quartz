package main

import (
	"crypto/sha1"
	"time"
)

// Sonarr EventTypes
const (
	SonarrEventDownloadBegin    = "Grab"
	SonarrEventDownloadComplete = "Download"
	SonarrEventUpgrade          = "Download"
	SonarrEventRename           = "Rename"
)

// Sonarr SeriesTypes
const (
	SonarrSeriesAnime    = "Anime"
	SonarrSeriesDaily    = "Daily"
	SonarrSeriesStandard = "Standard"
)

// SonarrEvent describes a Sonarr event
type SonarrEvent struct {
	Episodes []struct {
		ID             int       `json:"id"`
		EpisodeNumber  int       `json:"episodenumber"`
		SeasonNumber   int       `json:"seasonnumber"`
		Title          string    `json:"title"`
		AirDate        string    `json:"airdate"`
		AirDateUtc     time.Time `json:"airdateutc"`
		Quality        string    `json:"quality"`
		QualityVersion int       `json:"qualityversion"`
		ReleaseGroup   string    `json:"releasegroup"`
		SceneName      string    `json:"scenename"`
	} `json:"episodes" bson:"episodes"`
	EpisodeFile struct {
		ID             int    `json:"id"`
		RelativePath   string `json:"relativepath"`
		Path           string `json:"path"`
		Quality        string `json:"quality"`
		QualityVersion int    `json:"qualityversion"`
		ReleaseGroup   string `json:"releasegroup"`
		SceneName      string `json:"scenename"`
	} `json:"episodefile" bson:"episodefile"`
	Release struct {
		Quality        string `json:"quality"`
		QualityVersion int    `json:"qualityversion"`
		ReleaseGroup   string `json:"releasegroup"`
		ReleaseTitle   string `json:"releasetitle"`
		Indexer        string `json:"indexer"`
		Size           int    `json:"size"`
	} `json:"release" bson:"release"`
	IsUpgrade bool   `json:"isupgrade" bson:"isupgrade"`
	EventType string `json:"eventtype" bson:"eventtype"`
	Series    struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Path   string `json:"path"`
		TvdbID int    `json:"tvdbid"`
	} `json:"series" bson:"series"`
}

// EpisodeInfoID returns a hash of a concatenation of some information that is consistent between Grab/Download Sonarr event pairs
func (e *SonarrEvent) EpisodeInfoID() string {
	var info string
	for _, ep := range e.Episodes {
		info = info + string(ep.ID) + string(ep.EpisodeNumber) + string(ep.SeasonNumber) + ep.ReleaseGroup
	}
	hash := sha1.Sum([]byte(info))
	return string(hash[:20])
}
