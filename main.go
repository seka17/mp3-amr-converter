package main

import (
	"fmt"
	"github.com/nats-io/nats"
)

func main() {
	natsURL := flag.String("nats", nats.DefaultURL, "nats URL")
	toMP3 := flag.String("mp3", "toMP3", "subject for mp3 conversion")
	toAMR := flag.String("amr", "toAMR", "subject for amr conversion")
	flag.Parse()

	for {
		n, err := nats.Connect(*natsURL)
		if err != nil {
			time.Sleep(time.Second * 5)
			continue
		} else {
			nc, _ := nats.NewEncodedConn(n, nats.JSON_ENCODER)
			server.nc = nc
			defer nc.Close()
			break
		}
	}
	// Request to convert from AMR to MP3
	if _, err := nc.Subscribe(*toMP3, func(_, reply, msg string) {
	}); err != nil {
		return
	}
}
