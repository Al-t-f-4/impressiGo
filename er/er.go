package er

import ( 
    "fmt"
    "errors"
)

func my_err ( name string ) ( string, error ) {
    if name == "" {
        return "", errors.New ( "empty name" )
    }

    msg := fmt.Sprintf ( "Привет, %v. Welcome!", name )

    return msg, nil
}
