package er

import ( 
    "fmt"
    "errors"
)

func My_err ( name string ) ( string, error ) {
    if name == "" {
        return "", errors.New ( "empty name" )
    }

    msg := fmt.Sprintf ( "Привет, мое имя - %v.", name )

    return msg, nil
}
