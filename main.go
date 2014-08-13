package main

import (

    server "Server/server"
    
)

func main() {
    
    server.Init()
    server.ManageInit()
    server.Run()
// server.TestDatabase()
    
}