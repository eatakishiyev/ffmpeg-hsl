package service

import (
	"bytes"
	"hsl_proxy/dto"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func Hls(c echo.Context) (err error) {
	var request *dto.Request
	c.Bind(&request)

	playlist_name := strings.Join([]string{request.MeasurementPointUUID, "m3u8"}, ".")
	go startffmpeg(request, playlist_name)
	time.Sleep(time.Second * 5)
	return c.JSON(http.StatusCreated, dto.Response{M3U8Filename: playlist_name})
}

func startffmpeg(request *dto.Request, playlist_name string) {
	var outputbuff, errbuff bytes.Buffer
	cmd := exec.Command("/home/administrator/ffmpeg.sh", request.RTSPUrl, playlist_name, request.Transport)
	cmd.Stdout = &outputbuff
	cmd.Stderr = &errbuff

	err := cmd.Run()

	if err != nil {
		log.Println(err)
		log.Printf("%s , %s", outputbuff.String(), errbuff.String())
	}
}
