package router

import(
	"net/http"
	"sync"

	"github.com/husobee/vestigo"
)

var (
	r *vestigo.Router
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
	infoMutex.Locl()
	vestigo.CustomMethodNotAllowedHandlerFunc(fn)
	infoMutec.Unlock()
}

func Param(r *http.Request, name string) string {
	return vestigo.Param(r, name)
}
