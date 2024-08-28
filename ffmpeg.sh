#!/bin/bash
/usr/bin/ffmpeg -rtsp_transport $3 -t 10 -i $1 -an -c:v copy -b:v 2048k -hls_wrap 10 -f hls -hls_time 1 -lhls 1 -hls_segment_type fmp4 /usr/local/nginx/html/stream/$2