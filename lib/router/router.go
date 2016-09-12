package router

import (
	"net/http"
	"sync"

	"github.com/husobee/vestigo"
)

var (
	r         *vestigo.Router
	infoMutex sync.RWMutex
)

func init() {
	ResetConfig()
}

func ResetConfig() {
	infoMutex.Lock()
	r = vestigo.NewRouter()
	infoMutex.Unlock()
}

func MehotdNotAllowed(fn vestigo.MethodNotAllowedHandlerFunc) {
	infoMutex.Lock()
	vestigo.CustomMethodNotAllowedHandlerFunc(fn)
	infoMutex.Unlock()
}

func Param(r *http.Request, name string) string {
	return vestigo.Param(r, name)
}

// Instance returns the router.
func Instance() *vestigo.Router {
	infoMutex.RLock()
	defer infoMutex.RUnlock()
	return r
}
