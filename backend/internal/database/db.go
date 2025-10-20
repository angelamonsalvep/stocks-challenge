package database

import "fmt"

// DB es una referencia global (por ahora simulada)
var DB interface{}

// Connect inicializa la conexión a la base de datos (placeholder)
func Connect() error {
	fmt.Println("ℹ️  Base de datos no configurada aún. Usando modo sin conexión.")
	return nil
}
