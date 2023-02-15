package capture_to_rtsp

import (
	"fmt"
	"server/utils/cmd_unit"
	"time"
)

var (
	CaptureCmd = "ffmpeg.exe -r 20 -probesize 50M -rtbufsize 256M -f dshow -i video=\"Video (00 Pro Capture Mini HDMI)\" -s 1920x1080  -vcodec libx264 -pix_fmt nv12 -r 20 -b:v 2M -maxrate 2M  -bufsize 2M -bf 0 -g 20  -profile:v baseline  -preset veryfast -f rtsp -rtsp_transport tcp"
	CaptureURL = "rtsp://127.0.0.1:554/capture.sdp"
)

func Start(rtsp_port int) {

	go func() {

		for true {

			//capture_cmd := "ffmpeg.exe -r 20 -probesize 50M -rtbufsize 256M -f dshow -i video=\"Video (00 Pro Capture Mini HDMI)\" -s 1920x1080  -vcodec libx264 -pix_fmt nv12 -r 20 -b:v 2M -maxrate 2M  -bufsize 2M -bf 0 -g 20  -profile:v baseline  -preset veryfast -f rtsp -rtsp_transport tcp"
			capture_cmd := CaptureCmd + " " + CaptureURL

			cmd_err := cmd_unit.RunWait(capture_cmd)
			if cmd_err != nil {
				fmt.Println("采图出错,  capture_cmd:", capture_cmd)
			}
			time.Sleep(1 * time.Second)
		}

	}()

}
func Stop() {

}
