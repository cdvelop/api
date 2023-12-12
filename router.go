package api

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cdvelop/model"
	out "github.com/cdvelop/output"
)

func (c config) ServeMuxAndRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		action_type, api_name := getMethodAndObjectFromPath(r.URL.Path)

		// fmt.Println("ACTION", action_type, "API NAME", api_name)

		var registered_user bool
		u, err := c.BackendCheckUser(r)
		if err != "" {
			u = &model.User{
				Ip: GetIP(r), //USUARIO NO IDENTIFICADO
			}
		} else {
			registered_user = true
		}

		// fmt.Println("registered_user:", registered_user, u)

		p := &petition{
			u:   u,
			o:   &model.Object{ObjectName: "error"},
			r:   r,
			w:   w,
			t:   time.Now(),
			err: err,
		}

		if action_type != "" && r.Method != "GET" {
			out.PrintInfo(fmt.Sprintf("método:[%v]: acción:[%v]: objeto[%v]\n", r.Method, action_type, api_name))
		}

		if r.Method == "POST" {

			c.Log("OK HANDLER AUTH", c.NameOfAuthHandler())
			c.Log("-- API NAME", api_name)
			// time.Sleep(1 * time.Second)

			if c.NameOfAuthHandler() != api_name {

				if !registered_user {
					c.unauthorized(p, "realizar operaciones de lectura o escritura")
					return
				}
			}

			err := c.isHandlerOk(p, action_type, api_name)
			if err != "" {
				c.error(p, err)
				return
			}

			p.object_response = api_name
			if c.NameOfAuthHandler() == api_name {
				p.object_response = "" // cuando es login permitimos que responda multiples datos
				// agregamos el writer al usuario para que el manejador de authentication pueda crear la cookie
				u.W = p.w
			}

			switch p.action {
			case "create":
				c.create(p)

			case "read":
				c.read(p)

			case "update":
				c.update(p)

			case "delete":
				c.delete(p)

			case "upload":
				c.upload(p)

			default:
				c.error(p, "acción: "+p.action+" no permitida con método "+r.Method)
			}

		} else if r.Method == "GET" {

			switch action_type {

			case "file":
				if !registered_user {
					c.unauthorized(p, "leer archivos")
					return
				}

				c.readFile(p)

			case "static":
				c.static(w, r)

			default:
				if r.URL.Path == "/" {

					index_content, e := os.ReadFile(filepath.Join(INDEX_FOLDER, "index.html"))
					if e != nil {
						c.error(p, "ReadFile error. pagina index.html no encontrada", http.StatusInternalServerError)
						c.Log(e)
						return
					}

					t, e := template.New("").Parse(string(index_content))
					if e != nil {
						c.error(p, "template html error", http.StatusInternalServerError)
						c.Log(e)
						return
					}

					var boot_data_st = "none"
					if registered_user {
						boot_data_st, err = c.BackendLoadBootData(p.u)
						if err != "" {
							c.error(p, err, http.StatusInternalServerError)
							return
						}
					}

					var actions = model.BootPageData{
						JsonBootActions: boot_data_st,
					}

					e = t.Execute(w, actions)
					if e != nil {
						c.error(p, "error al renderizar sitio", http.StatusInternalServerError)
						c.Log(e)
						return
					}
					// w.Header.Set()
					// fmt.Fprint(w, "¡Hola! Esta es la página principal.")
					// w.Write()
					// http.ServeFile(w, r, INDEX_FOLDER+"/index.html")
				} else {
					c.error(p, "error not found "+r.URL.Path)
					http.NotFound(w, r)
				}
			}

		} else {
			c.error(p, "método "+r.Method+" no permitido", http.StatusMethodNotAllowed)
		}

	})

	return mux
}
