package main

import (
    "fmt"
    "math/rand"
)

func main() {

    caratteri := "abcdefghijklmnopqrstuvwzyxABCDEFGHIJKLMNOPQRSTUVWZYX"

    caratterispeciali := "@[]#.,?^|'ì°*éò_:;ç"

    var rispostacaratterispeciali string

    fmt.Println("Vuoi includere caratteri speciali nella tua password (s/n)?")
    fmt.Scanln(&rispostacaratterispeciali)

    if rispostacaratterispeciali == "s" {
        fmt.Println("Ok, includerò caratteri speciali")
        caratteri = caratteri + caratterispeciali
    }

    if rispostacaratterispeciali == "n" {
        fmt.Println("Ok, non includerò caratteri speciali")
    }

    lunghezzalista := len(caratteri)

    password := ""
    
    var lunghezzaPassword int

    fmt.Print("Scegli la lunghezza della tua password [esempio: 8]")

    fmt.Scanln(&lunghezzaPassword)
    

    for i := 0; i < lunghezzaPassword; i++ {
        indicecasuale := rand.Intn(lunghezzalista)
        caratterescelto := string(caratteri[indicecasuale])
        password = password + caratterescelto
    }
fmt.Println("La tua password è", password)
}

