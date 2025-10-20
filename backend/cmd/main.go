package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"stocks-challenge/internal/api"
)

func main() {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No se encontró el archivo .env, usando valores por defecto.")
	}

	// Leer el puerto desde la variable de entorno o usar 8081 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Inicializar el router de Gin
	router := gin.Default()

	// Registrar las rutas
	api.SetupRoutes(router)

	// Iniciar el servidor
	log.Printf("🚀 Servidor ejecutándose en http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Error al iniciar el servidor: %v", err)
	}
}
