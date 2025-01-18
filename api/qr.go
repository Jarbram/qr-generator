package api

import (
    "bytes"
    "encoding/base64"
    "fmt"
    "image/color"
    "net/http"

    "github.com/skip2/go-qrcode"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    // Configurar CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    
    // Obtener el texto del query parameter
    text := r.URL.Query().Get("text")
    if text == "" {
        http.Error(w, "El parámetro 'text' es requerido", http.StatusBadRequest)
        return
    }

    // Crear el código QR en memoria
    var buf bytes.Buffer
    qr, err := qrcode.New(text, qrcode.High)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error al generar QR: %v", err), http.StatusInternalServerError)
        return
    }

    // Configurar colores
    qr.BackgroundColor = color.Black
    qr.ForegroundColor = color.White

    // Generar PNG en memoria
    png, err := qr.PNG(1040)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error al generar PNG: %v", err), http.StatusInternalServerError)
        return
    }

    // Configurar headers para imagen PNG
    w.Header().Set("Content-Type", "image/png")
    w.Write(png)
} 