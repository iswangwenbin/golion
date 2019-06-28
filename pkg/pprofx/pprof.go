package pprofx

import (
	"github.com/micro/go-config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/pprof"
)

func NewPProf() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	addr := config.Get("ServiceSettings", "PProfListenAddress").String("")
	go func() {
		if err := http.ListenAndServe(addr, mux); err != nil {
			log.WithFields(log.Fields{
				"addr":  addr,
				"error": err,
			}).Info("http.ListenAndServe err")
		}
	}()
}
