package main

import(
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	// endpoint de prueba para verificar la conexión con el servidor
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Aigis Gateway Operativo")
	})

	port:= ":8001"
	log.Printf("Iniciando AI Hub en el puerto %s...", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Fallo crítico en el servidor: %v", err)
	}
}
