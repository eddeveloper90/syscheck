package httpServer

import (
	"fmt"
	"net"
	"net/http"
	"xcheck/config"
	"strings"
	"time"

	"github.com/felixge/httpsnoop"
)

const DATETIME_LOG_FORMAT = "2006-01-02 15:04:05"

func GetIP(r *http.Request) string {
	nginx := r.Header.Get("X-NGINX-IP")
	cloudflare := r.Header.Get("X-FORWARDED-FOR")
	arvan := r.Header.Get("X-Real-IP")
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	var realIp string
	if cloudflare != "" {
		realIp = cloudflare
	} else if arvan != "" {
		realIp = arvan
	} else if nginx != "" {
		realIp = nginx
	} else {
		realIp = ip
	}

	if err != nil {
		fmt.Println(time.Now().Format(DATETIME_LOG_FORMAT), r.Method, realIp, port, err)
	} else {
		fmt.Println(time.Now().Format(DATETIME_LOG_FORMAT), r.Method, realIp, port)
	}

	return realIp
}

func home(w http.ResponseWriter, r *http.Request) {
	GetIP(r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("hello"))
}

// HTTPReqInfo describes info about HTTP request
type HTTPReqInfo struct {
	// GET etc.
	method  string
	url     string
	referer string
	ipaddr  string
	// response code, like 200, 404
	code int
	// number of bytes of the response sent
	size int64
	// how long did it take to
	duration  time.Duration
	userAgent string
}

func (info HTTPReqInfo) print() {
	fmt.Println(time.Now().Format(DATETIME_LOG_FORMAT),
		info.method,
		info.ipaddr,
		info.url,
		info.referer,
		info.code,
		info.size,
		info.duration,
		info.userAgent)
}

// Request.RemoteAddress contains port, which we want to remove i.e.:
// "[::1]:58292" => "[::1]"
func ipAddrFromRemoteAddr(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}

// requestGetRemoteAddress returns ip address of the client making the request,
// taking into account http proxies
func requestGetRemoteAddress(r *http.Request) string {
	hdr := r.Header
	hdrRealIP := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")
	if hdrRealIP == "" && hdrForwardedFor == "" {
		return ipAddrFromRemoteAddr(r.RemoteAddr)
	}
	if hdrForwardedFor != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		parts := strings.Split(hdrForwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		// TODO: should return first non-local address
		return parts[0]
	}
	return hdrRealIP
}

// return true if this request is a websocket request
func isWsRequest(r *http.Request) bool {
	uri := r.URL.Path
	return strings.HasPrefix(uri, "/ws/")
}

// simplest possible server that returns url as plain text
func handleIndex(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("You've called url %s", r.URL.String())
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK) // 200
	w.Write([]byte(msg))
}

func logRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// websocket connections won't work when wrapped
		// in RecordingResponseWriter, so just pass those through
		if isWsRequest(r) {
			h.ServeHTTP(w, r)
			return
		}

		ri := &HTTPReqInfo{
			method:    r.Method,
			url:       r.URL.String(),
			referer:   r.Header.Get("Referer"),
			userAgent: r.Header.Get("User-Agent"),
		}

		ri.ipaddr = requestGetRemoteAddress(r)

		// this runs handler h and captures information about
		// HTTP request
		m := httpsnoop.CaptureMetrics(h, w, r)

		ri.code = m.Code
		ri.size = m.Written
		ri.duration = m.Duration
		//fmt.Println(*ri)
		//logHTTPReq(ri)
		ri.print()
	}
	return http.HandlerFunc(fn)
}

func makeHTTPServer() *http.Server {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handleIndex)
	var handler http.Handler = mux

	handler = logRequestHandler(handler)

	srv := &http.Server{
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		IdleTimeout:  120 * time.Second, // introduced in Go 1.8
		Handler:      handler,
		Addr:         "0.0.0.0:" + (*config.CONFIG).HttpServer.Port,
	}

	return srv
}

func StartServer() {
	httpSrv := makeHTTPServer()
	httpSrv.ListenAndServe()
	//http.HandleFunc("/", home)
	//log.Fatal(http.ListenAndServe("0.0.0.0:"+(*config.CONFIG).HttpServer.Port, nil))
}
