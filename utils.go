package main

import (
	"net/url"
	"path"

	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/friendly/v2"
	"github.com/lcpluginmaker/gilc"
)

func getDownloadDir(data gilc.IData) string {
	user := data.Username
	if user == "" || user == "/" {
		user = "_"
	}
	return path.Join(data.SavePath, "home", user, "downloads")
}

func download(u string, dlDir string) error {
	URL, err := url.Parse(u)
	if err != nil {
		arrowprint.Err0("invalid url: '%s'", u)
		return err
	}

	fileName := path.Base(URL.Path)
	if fileName == "" || fileName == "/" {
		fileName = "index.html"
	}
	err = friendly.DownloadFile(u, path.Join(dlDir, fileName))
	if err != nil {
		arrowprint.Err0("download for '%s' failed: %s", u, err.Error())
		return err
	}

	arrowprint.Suc1("downloaded '%s' successfully", u)
	return nil
}
