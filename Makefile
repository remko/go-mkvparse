check:
	go test .

lint:
	staticcheck .
	go vet .

example-live+junk.mkv:
	ffmpeg -t 1 -s 320x240 -f rawvideo -r 25 -pix_fmt rgb24 -i /dev/zero -metadata title="Live + Junk" -metadata author="John Doe" -c:v libx264 -pix_fmt yuv420p dirty.$@
	mkclean --live dirty.$@ $@
	-rm -rf dirty.$@
