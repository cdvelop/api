package api

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const INDEX_FOLDER = "frontend/built"
const STATIC_FOLDER = "frontend/built/static"

func (c config) static(w http.ResponseWriter, r *http.Request) {

	// Obtener el nombre del archivo solicitado
	file := r.URL.Path[len("/static/"):]
	// fmt.Printf("archivo estáticos %s solicitado\n", file)
	// Leer el contenido del archivo
	content, err := os.ReadFile(STATIC_FOLDER + "/" + file)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Obtener el tipo de archivo correspondiente a la extensión
	mimeType := mime.TypeByExtension(filepath.Ext(file))

	// Agregar la cabecera de tipo de archivo a la respuesta
	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
	}

	w.Header().Set("Cache-Control", c.static_cache)

	// Escribir el contenido en la respuesta
	w.Write(content)

}
