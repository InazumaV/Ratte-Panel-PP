package main

import (
	"Ratte-Panel-PP/impl"
	"github.com/InazumaV/Ratte-Interface/panel"
	log "github.com/sirupsen/logrus"
)

func main() {
	p, err := panel.NewServer(nil, impl.New())
	if err != nil {
		log.Fatalln(err)
	}
	if err = p.Run(); err != nil {
		log.Fatalln(err)
	}
}
