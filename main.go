package main

import (
	"os"

	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/friendly"
	"github.com/lcpluginmaker/gilc"
)

func pcommand(data gilc.IData, args []string) {
	if len(args) < 1 {
		arrowprint.Err0("you need to pass a url to download")
		return
	}

	dlDir := getDownloadDir(data)
	if !friendly.IsDir(dlDir) {
		err := os.MkdirAll(dlDir, 0700)
		if err != nil {
			arrowprint.Err0("cannot create download folder: %s", err.Error())
			return
		}
	}

	ok := false
	for i, u := range args {
		arrowprint.Info0("downloading file %d of %d: '%s'...", i, len(args), u)
		err := download(u, dlDir)
		if err == nil {
			ok = true
		}
	}

	if ok {
		arrowprint.Info0("your file(s) have been downloaded to %s", dlDir)
	}
}

func main() {
	gilc.Setup("download file from it's url", func(i gilc.IData) {}, pcommand, func(i gilc.IData) {})
	gilc.Run()
}
