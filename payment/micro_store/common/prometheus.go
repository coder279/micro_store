package common

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
)

func PromethuesBoot(port int){
	http.Handle("/metrics",promhttp.Handler())
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port),nil)
		if err != nil {
			log.Fatal(err)
			log.Error(err)
		}
		log.Info("监控启动,端口为: "+strconv.Itoa(port))
	}()
}