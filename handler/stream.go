package handler

import (
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func Stream(c *gin.Context) {
	cmd := exec.Command("ffmpeg", "-i", "./Media/Cartoon.mp4", "-c:v",
		"libx264", "-c:a", "aac", "-b:v", "1000k", "-b:a", "128k", "-map", "0", "-f",
		"segment", "-segment_time", "10", "-segment_list", "./Out/outputlist.m3u8", "-segment_format",
		"mpegts", "./Out/output%03d.ts")
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

	log.Println("Video encoding complete")
	c.Redirect(302,"/video")
}
