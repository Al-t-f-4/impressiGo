package main

import (
    "fmt"
    "moduleI/greetings"
    "moduleI/nova"
)

func main ( ) {
    fmt.Println( "Main" )
    greetings.Hello ( )
    nova.Greet()
}
