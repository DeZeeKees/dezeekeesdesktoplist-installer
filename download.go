package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Release struct {
	Assets []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func DownloadLatestRelease() error {
	resp, err := http.Get("https://api.github.com/repos/DeZeeKees/dezeekeesdesktoplist-app/releases")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var releases []Release
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return err
	}

	for _, release := range releases {
		for _, asset := range release.Assets {
			if asset.Name == "dezeekeesdesktoplist.exe" {
				resp, err := http.Get(asset.BrowserDownloadURL)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				out, err := os.Create(filepath.Join(installPath, asset.Name))
				if err != nil {
					return err
				}
				defer out.Close()

				_, err = io.Copy(out, resp.Body)
				return err
			}
		}
	}

	return nil
}
