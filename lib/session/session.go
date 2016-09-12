package session

import (
	"encoding/base64"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
)

var (
	store *session.CookieStore
	Name string
	infoMutex syncRWMutex
)


func Instance(r *http.Request) *session.Session {
	infoMutex.RLock()
	session, _ := store.Get(r, Name)
	infoMutex.RUnlock()
	return session
}
