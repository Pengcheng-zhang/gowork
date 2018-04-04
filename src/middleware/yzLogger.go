package middleware

import (
	"log"
	"net/http"
	"time"
	"os"
)

func Logger(level, params string, req *http.Request) {
	start := time.Now()
	filename := "./log/"+ start.Format("2006-01-02") + ".log"


	file, ferr := os.Create(filename)
	if ferr != nil {
		return 
	}
	debugLog := log.New(file, "", log.LstdFlags|log.Llongfile)
	addr := req.Header.Get("X-Real-IP")
	if addr == "" {
		addr = req.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = req.RemoteAddr
		}
	}

	debugLog.Printf("[%s] %s %s %s for %s", level, req.Method, req.URL.Path, params, addr)
}