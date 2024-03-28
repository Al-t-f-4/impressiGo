package logg

import (
    "fmt"
    "log"
    "moduleI/er"
)

func my_log ( ) {

    log.SetPrefix ( "Err module: " )
    log.SetFlags ( 0 )

    msg, err := er.my_err ( "" )
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println ( msg )
}
