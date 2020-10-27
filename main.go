package main

import (
	"bufio"
	"fmt"
	"os"

	"./contenidos"
)
func main() {
	var opc int
	
	//img := contenidos.Imagen{Titulo: "5", Formato: "5", Canales: "tem2"}
	//aud := contenidos.Audio{Titulo: "la Caida de Oscar", Formato: ".mp3", Duracion: "200 s"}
	scanner := bufio.NewScanner(os.Stdin)
	contenedor:=contenidos.ContenidoWeb{}
	cont := 1
	for cont != 0 {
		opc = 0
		fmt.Println("1.- Registro Imagen\n2.- Registro Audio\n3.- Registro Video")
		fmt.Println("4.- Mostar Registros\n0.- Salir")
		fmt.Scan(&opc)
		if opc == 1 {
			var tem, tem1, tem2 string
			fmt.Println("Registro de Imagen")
			fmt.Println("Titulo")
			scanner.Scan()
			scanner.Scan()
			tem = scanner.Text()
			fmt.Println("Formato")
			scanner.Scan()
			tem1 = scanner.Text()
			fmt.Println("Canal")
			scanner.Scan()
			tem2 = scanner.Text()
			img := contenidos.Imagen{Titulo: tem, Formato: tem1, Canales: tem2}
			contenedor.Archivos = append(contenedor.Archivos,&img)
		} else if opc == 2 {
			var tem, tem1, tem2 string
			fmt.Println("Registro de Audio")
			fmt.Println("Titulo")
			scanner.Scan()
			scanner.Scan()
			tem = scanner.Text()
			fmt.Println("Formato")
			fmt.Scan(&tem1)
			fmt.Println("Duracion")
			fmt.Scan(&tem2)
			aud := contenidos.Audio{Titulo: tem, Formato: tem1, Duracion: tem2}
			contenedor.Archivos = append(contenedor.Archivos,&aud)
		}else if opc ==3{
			var tem, tem1, tem2 string
			fmt.Println("Registro de Video")
			fmt.Println("Titulo")
			scanner.Scan()
			scanner.Scan()
			tem = scanner.Text()
			fmt.Println("Formato")
			fmt.Scan(&tem1)
			fmt.Println("Frames")
			fmt.Scan(&tem2)
			video := contenidos.Video{Titulo: tem, Formato: tem1, Frames: tem2}
			contenedor.Archivos = append(contenedor.Archivos,&video)
		}else if opc ==4{
			contenedor.MostrarT()
		} else if opc == 0 {
			cont = 0
		}
	}

}
