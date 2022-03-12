package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/sirupsen/logrus"
	"github.com/wallissonmarinho/challenge_digital_republic/internal/endpoint"
	"github.com/wallissonmarinho/challenge_digital_republic/internal/service"
	trans "github.com/wallissonmarinho/challenge_digital_republic/internal/transport/http"
)

const (
	banner = ` 
	________  ___  ________  ___  _________  ________  ___               ________  _______   ________  ___  ___  ________  ___       ___  ________     
	|\   ___ \|\  \|\   ____\|\  \|\___   ___\\   __  \|\  \             |\   __  \|\  ___ \ |\   __  \|\  \|\  \|\   __  \|\  \     |\  \|\   ____\    
	\ \  \_|\ \ \  \ \  \___|\ \  \|___ \  \_\ \  \|\  \ \  \            \ \  \|\  \ \   __/|\ \  \|\  \ \  \\\  \ \  \|\ /\ \  \    \ \  \ \  \___|    
	 \ \  \ \\ \ \  \ \  \  __\ \  \   \ \  \ \ \   __  \ \  \            \ \   _  _\ \  \_|/_\ \   ____\ \  \\\  \ \   __  \ \  \    \ \  \ \  \       
	  \ \  \_\\ \ \  \ \  \|\  \ \  \   \ \  \ \ \  \ \  \ \  \____        \ \  \\  \\ \  \_|\ \ \  \___|\ \  \\\  \ \  \|\  \ \  \____\ \  \ \  \____  
	   \ \_______\ \__\ \_______\ \__\   \ \__\ \ \__\ \__\ \_______\       \ \__\\ _\\ \_______\ \__\    \ \_______\ \_______\ \_______\ \__\ \_______\
		\|_______|\|__|\|_______|\|__|    \|__|  \|__|\|__|\|_______|        \|__|\|__|\|_______|\|__|     \|_______|\|_______|\|_______|\|__|\|_______|
																																																																						                                                                                                                                                   
`
)

func main() {

	logrus.Info(banner)

	// initialize our OpenCensus configuration and defer a clean-up
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger)
	defer level.Info(logger)

	var (
		context    context.Context
		services   = service.NewServiceFactory(logger)
		endpoint   = endpoint.MakeEndpoints(services, logger)
		serverHTTP = trans.NewService(context, &endpoint, &logger)
		httpAddr   = flag.String("http.addr", ":1707", "HTTP listen address")
		err        = make(chan error)
	)

	go func() {
		server := &http.Server{
			Addr:         *httpAddr,
			Handler:      serverHTTP,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		}
		err <- server.ListenAndServe()
	}()

	fatal := level.Error(logger).Log("exit", <-err)
	if fatal != nil {
		logrus.Error(fatal)
	}

}
