package contenidos

import "fmt"

type Imagen struct {
	Titulo, Formato, Canales string
}
type Audio struct {
	Titulo, Formato, Duracion string
}
type Video struct {
	Titulo, Formato, Frames string
}

func (x *Imagen) Mostar() string {
	return "Titulo: " + x.Titulo + "\nFormato: " + x.Formato + "\nCanales: " + x.Canales+"\n"
}
func (x *Audio) Mostar() string {
	return "Titulo: " + x.Titulo + "\nFormato: " + x.Formato + "\nDuracion: " + x.Duracion+"\n"
}
func (x *Video) Mostar() string {
	return "Titulo: " + x.Titulo + "\nFormato: " + x.Formato + "\nFrames: " + x.Frames+"\n"
}

type Multimedia interface {
	Mostar() string
}
type ContenidoWeb struct {
	Archivos []Multimedia
}

func (c *ContenidoWeb) MostrarT() {
	for _, i := range c.Archivos {
		fmt.Println(i.Mostar())
	}
}
