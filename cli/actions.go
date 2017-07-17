package cli

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/corpix/filelist/config"
	"github.com/corpix/filelist/matcher"
)

const (
	PATHS_CAP = 1024
)

func walk(paths []string, stream chan string) {
	for _, p := range paths {
		ap, err := filepath.Abs(p)
		if err != nil {
			log.Error(err)
			return
		}
		err = filepath.Walk(ap, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			stream <- path
			return nil
		})
		if err != nil {
			log.Error(err)
			return
		}
	}
	close(stream)
}

func filterAction(ctx *cli.Context) {
	var err error
	c, err := config.FromFile(ctx.String("config"))
	if err != nil {
		log.Error(err)
		return
	}
	config.SetCurrent(c)
	ex := matcher.New(c.Excludes)
	in := matcher.New(c.Includes)
	stream := make(chan string, PATHS_CAP)
	done := make(chan bool)
	go walk(ctx.Args(), stream)
	go func() {
		for p := range stream {
			if !ex.Match(p) || in.Match(p) {
				fmt.Println(p)
			}
		}
		done <- true
	}()

	<-done
}
