# Conociendo Go - Curso Victor Robles - UDEMY
Como programar y crear APIs RESTful con Golang

> Author @eduardoonetto

<img src="https://res.cloudinary.com/practicaldev/image/fetch/s--ZmWHP0Bg--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://i.postimg.cc/VLdgRJXF/Clipart-Key-2207878.png" width="60%" style="margin-left: 20%" alt="logo_golang" title="Logo Golang">
<hr/>

## 🛸 Cómo se usa     
Instalar Golang desde su sitio oficial:
<a href="https://go.dev/dl/" target="_blank">Golang Download Here!.</a>
<hr/>

## 🌎 Curso GoLang:
<hr/>

#### 📚 Introducción:
01. Conocer tipos de datos y comandos con Go.
02. Generar un tipo de datos personalizado segun su modelo.
03. Crear y llamar funciones con Go.
04. Conocer Arrays y Slices con Go.
05. Estructuras de control (if, bucles, switch) y recibir args desde la consola.

<hr/>

#### 📚 Ficheros:

01. Leer y Escribir en Ficheros (util para generar logs).

<hr/>

### 🌠 Servidor Web y APIs RESTful con Golang y MongoDB:



#### Requisitos:

📌 Instalar Gorilla mux: 
Dentro del directorio `Api` en la consola, ejecutar lo Siguiente:
* `go mod init 01-main.go`
* `go get -u github.com/gorilla/mux`
* `go mod tidy`

🧐 En caso de tener VSCode y de un error en la linea package main, se debe incluir lo siguiente en el archivo settings.json:
`"gopls": { "experimentalWorkspaceModule": true }`
Luego reiniciar VSCode.

01. Creacion de Cliente-Servidor.

🤞 Para hacer correr el servicio: 

`go run .`

🚀 Guardar una pelicula:

`curl -X POST "http://localhost:8000/pelicula" -d '{"name": "El lobo de wall street", "year": 2015, "director": "Martin"}' -i`

👀 Ver Pelicula:

<p>http://localhost:8000/peliculas</p>
<hr/>

## 🕵 Cómo contribuir 
Puedes crear un pull request al proyecto 😉.