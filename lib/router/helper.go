package router

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/justinas/alice"
)

var (
	routeList []string
	listMutex sync.RWMutex
)


// Record stores the method and path.
func record(method, path string) {
	listMutex.Lock()
	routeList = append(routeList, fmt.Sprintf("%v\t%v", method, path))
	listMutex.Unlock()
}

// Get is a shortcut for router.Handle("GET", path, handle).
func Get(path string, fn http.HandlerFunc, c ...alice.Constructor) {
	record("GET", path)

	infoMutex.Lock()
	r.Get(path, alice.New(c...).ThenFunc(fn).(http.HandlerFunc))
	infoMutex.Unlock()
}


// ChainHandler returns a handler of chained middleware.
func ChainHandler(h http.Handler, c ...alice.Constructor) http.Handler {
	return alice.New(c...).Then(h)
}
