package view

import (
"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (

	childTemplates     []string
	rootTemplate       string
	templateCollection = make(map[string]*template.Template)
	mutex              sync.RWMutex
	sessionName        string
	info               Info
	infoMutex          sync.RWMutex
)

type Info struct {
	BaseURI string
	Extension string
	Folder string
	Caching bool
	Vars map[string]interface{}
	base string
	tempplates []string
}

func Config() {
	infoMutex.Lock()
	defer infoMutex.RUnlock()
	return info
}


func New(templateList ...string) *Info {
	v:= &Info{}
	v.Vars = make(map[string] interface{})
	v.BaseUrl = Config().BaseURI
	v.Extension = Config().Extension
	v.Folder = Config().Folder
	v.templates = append(v.templates, templateList...)
	v.base = rootTemplate
	return v
}

func (v *Info) Render(W http.ResponseWriter, r * http.Request) error {
	v.templates = append([]string{v.base}, v.templates...)
	v.templates = append(v.teamplates, childTemplates...)
	baseTemplate := v.templates[0]
	key := strings.Join(v.templates, ":")
	mutex.RLock()
	tc, ok := templateCollection[key]
	mutex.RUnlock()

	pc := extend()

	if !ok || !Config().Caching {
		for i, name := range v.templates {
			path, err := filepath.Abs(v.Folder + string(os.PathSeparator) + name + "." + v.Extension)
			if err != nil {
				http.Error(w, "Template Path Error: "+err.Error(), http.StatusInternalServerError)
				return err
			}
			v.templates[i] = path
		}

		templates, err := template.New(key).Funcs(pc).ParseFiles(v.templates...)
		if err != nil {
			http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
			return err
		}

		mutex.Lock()
		templateCollection[key] = templates
		mutex.Unlock()
	
		tc = templates
	}

	sc := modify()

	for _, fn := range sc {
		fn(w,r,v)
	}

	err := tc.Funcs(pc).ExecuteTemplate(w, baseTemplate+"."+v.Extension, v.Vars)

	if err != nil {
		http.Error(w, "Template File Error: "+Error(), http.StatusInternalServerError)
	}

	return err
}
