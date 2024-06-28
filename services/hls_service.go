package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hsl_proxy/dto"
	"math/rand"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func Hls(c echo.Context) (err error) {
	var request *dto.Request
	c.Bind(&request)
	//"rtsp://admin:Admin123@10.113.8.152:554/ISAPI/Streaming/Channels/102",
	playlist_name := strings.Join([]string{"/usr/local/nginx/html/stream/playlist", strconv.Itoa(rand.Int()), "m3u8"}, ".")
	cmd := exec.Command("/usr/bin/ffmpeg", "-i", request.RTSPUrl,
		"-vcodec", "copy", "-acodec", "copy", "-f", "hls", "-hls_time", "10", "-hls_list_size", "3", playlist_name)
	//cmd := exec.Command("psa", "-a")
	cmd.Start()
	if cmd.Err != nil {
		fmt.Printf("error occurred %s", err)
	}
	time.Sleep(time.Second * 10)
	if cmd.Process != nil {
		cmd.Process.Kill()
	}
	return c.JSON(http.StatusCreated, dto.Response{M3U8Filename: playlist_name})
}
