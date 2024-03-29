package main

import (
    "fmt"
    "moduleI/greetings"
    "moduleI/nova"
    "moduleI/mlg"
)

func main ( ) {
    fmt.Println( "Main" )
    greetings.Hello ( )
    nova.Greet()
    mlg.My_log ( )
}

