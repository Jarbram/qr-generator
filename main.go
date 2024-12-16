package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	fmt.Println("Bienvenido al generador de códigos QR")
	fmt.Println("Por favor, ingresa la URL que deseas convertir en un código QR:")

	// Leer la URL desde la entrada del usuario
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("URL: ")
	url, _ := reader.ReadString('\n')

	// Limpiar espacios en blanco y saltos de línea
	url = url[:len(url)-1]

	// Configuración del archivo QR
	outputFile := "qr-code.png" // Nombre del archivo generado
	size := 1040                // Tamaño del código QR

	// Crear el código QR con fondo negro y código blanco
	err := qrcode.WriteColorFile(
		url,
		qrcode.High,
		size,
		color.Black,      // Fondo negro
		color.White,      // Código blanco
		outputFile,       // Archivo de salida
	)

	if err != nil {
		fmt.Println("Error al generar el código QR:", err)
		return
	}

	fmt.Printf("¡Código QR generado exitosamente! Archivo: %s\n", outputFile)
}
