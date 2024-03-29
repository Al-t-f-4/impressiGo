package mlg

import (
    "fmt"
    "log"
    "moduleI/er"
)

func My_log ( ) {

    log.SetPrefix ( "Err module: " )
    log.SetFlags ( 0 )

    msg, err := er.My_err ( "Master" )
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println ( msg )
}
