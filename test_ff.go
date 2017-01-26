package main

import (
	"os/exec"
	"os"
)

func main() {
	cmd := exec.Command("ffmpeg", "-i", "rtsp://***@10.157.197.46/axis-media/media.amp", "-f",
		"segment", "-map", "0:0", "-vcodec", "copy", "-reset_timestamps", "1", "-f", "dash", "manifest.mpd")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		println(err.Error())
		return
	}
}
