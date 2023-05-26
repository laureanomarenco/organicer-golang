package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Crea un nuevo enrutador Gin
    r := gin.Default()

    // Define tus rutas y controladores
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "¡Hola, mundo!"})
    })

    // Inicia el servidor en el puerto 8080
    r.Run(":8080")
}