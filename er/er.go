package er

import ( 
    "fmt"
    "errors"
    "math/rand"
)

func My_err ( name string ) ( string, error ) {
    if name == "" {
        return "", errors.New ( "empty name" )
    }

    msg := fmt.Sprintf ( randomFormat(), name )

    return msg, nil
}

func randomFormat() string {
    formats := []string{
        "Привет, %v. Добро пожаловать!",
        "Рад тебя видеть, %v!",
        "HAIL, %v! СААААлаааам!",
    }

    return formats[rand.Intn(len(formats))]
}
