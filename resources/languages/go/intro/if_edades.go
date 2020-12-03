/*
1. Mostrar según la tabla el siguiente mensaje
Edad  Mensaje
1-5   Bebe
6-12  Niño
13-16 Adolescente
17-23 Joven
24-40 Adulto joven
41-59 Adulto
>60   Adulto Mayor
*/
/*&& -> and
|| -> or
!  -> not*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func clasiEdad(edad int) string {
	if edad >= 1 && edad <= 5 {
		return "bebe"
	} else if edad >= 6 && edad <= 12 {
		return "niño"
	} else if edad >= 6 && edad <= 12 {
		return "niño"
	} else if edad >= 13 && edad <= 16 {
		return "Adolescente"
	} else if edad >= 17 && edad <= 23 {
		return "Joven"
	} else if edad >= 24 && edad <= 40 {
		return "Adulto joven"
	} else if edad >= 41 && edad <= 59 {
		return "Adulto"
	} else {
		return "Adulto Mayor"
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("-" + r.URL.Path[11:] + "-")
	edad, err := strconv.Atoi(r.URL.Path[11:])
	if err == nil {
		fmt.Fprintf(w, "Tu eres un %s", clasiEdad(edad))
	} else {
		fmt.Fprintf(w, "Edad inválida")
	}
}

func main() {
	http.HandleFunc("/clasiedad/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
