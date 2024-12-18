package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Estructura para pasar datos a las plantillas HTML
type Datos struct {
	Titulo  string
	Mensaje string
	Nombre  string
}

func main() {
	// Ruta principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		datos := Datos{
			Titulo:  "Página Principal",
			Mensaje: "¡Bienvenido a mi sitio web!",
		}
		renderTemplate(w, "index.html", datos)
	})

	// Ruta dinámica con parámetro
	http.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
		nombre := r.URL.Query().Get("nombre")
		if nombre == "" {
			nombre = "Visitante"
		}
		datos := Datos{
			Titulo:  "Saludo",
			Mensaje: "¡Hola, " + nombre + "!",
			Nombre:  nombre,
		}
		renderTemplate(w, "saludo.html", datos)
	})

	// Servir archivos estáticos (CSS, JS, imágenes)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Iniciar el servidor
	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Función para renderizar plantillas HTML
func renderTemplate(w http.ResponseWriter, tmpl string, datos Datos) {
	t, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		http.Error(w, "No se pudo cargar la plantilla", http.StatusInternalServerError)
		return
	}
	t.Execute(w, datos)
}
