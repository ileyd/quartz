package main

import (
	"path/filepath"
	"time"
)

// RegisterSonarrEvent ... self exlanatory
func RegisterSonarrEvent(event SonarrEvent) (err error) {
	// if this is a download started event rather than a download complete event, we don't need to do anything more
	if event.EventType == SonarrEventDownloadBegin {
		return
	}

	// if is mkv remux to mp4
	extension := filepath.Ext(event.EpisodeFile.RelativePath)
	dir := filepath.Dir(event.Series.Path + "/" + event.EpisodeFile.RelativePath)
	if extension == "mkv" || extension == ".mkv" {
		time.Sleep(time.Second * 15) // conversion seems to get triggered too early
		RemuxMKVToMP4(dir, event.Series.Path+"/"+event.EpisodeFile.RelativePath)
		SonarrClient.RescanSeries(event.Series.ID)
	}

	return err
}
