package main

import (	
	"os"

	server "monitoring/internal/server"
    "monitoring/internal/utils"
)

func main() {
    if err := server.Run(); err != nil {
        utils.Log.Errorf("Ошибка: %v", err)
        os.Exit(1)
    }
}