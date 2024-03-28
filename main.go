package main

import (
    "fmt"
    "moduleI/greetings"
    "moduleI/nova"
    //"moduleI/mlg"
    "moduleI/er"
    "log"
)

func main ( ) {
    fmt.Println( "Main" )
    greetings.Hello ( )
    nova.Greet()
    //mlg.my_log ( )
    log.SetPrefix ( "Err module: " )
    log.SetFlags ( 0 )

    msg, err := er.my_err ( "" )
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println ( msg )
}

