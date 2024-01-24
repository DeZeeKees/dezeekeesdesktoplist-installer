package main

import (
	"encoding/json"
	"fmt"
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
	TagName    string `json:"tag_name"`
	Prerelease bool   `json:"prerelease"`
}

func DownloadLatestRelease(isPrerelease bool) error {

	var release Release
	var err error

	if isPrerelease {
		release, err = GetLatestPrerelease()
		if err != nil {
			return err
		}
	} else {
		release, err = GetLatestRelease()
		if err != nil {
			return err
		}
	}

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

	return nil
}

func GetLatestPrerelease() (Release, error) {
	resp, err := http.Get("https://api.github.com/repos/DeZeeKees/dezeekeesdesktoplist-app/releases")
	if err != nil {
		return Release{}, err
	}
	defer resp.Body.Close()

	var releases []Release
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return Release{}, err
	}

	for _, release := range releases {

		fmt.Println("Downloading Pre-release version " + release.TagName)
		if release.Prerelease {
			return release, nil
		}
	}

	return Release{}, nil
}

func GetLatestRelease() (Release, error) {
	resp, err := http.Get("https://api.github.com/repos/DeZeeKees/dezeekeesdesktoplist-app/releases/latest")
	if err != nil {
		return Release{}, err
	}
	defer resp.Body.Close()

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return Release{}, err
	}

	if release.TagName == "" {
		return Release{}, fmt.Errorf("failed to get latest release")
	}

	fmt.Println("Downloading release version " + release.TagName)

	return release, nil
}
