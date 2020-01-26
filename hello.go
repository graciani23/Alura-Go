package main

import (
    "fmt"
    "os"
    "net/http"
)

func main() {

    exibeIntroducao()

    for { // loop infinito
        exibeMenu()
        comando := leComando()
        
        switch comando {
        case 1:
            iniciarMonitoramento()
        case 2:
            fmt.Println("Exibindo Logs...")
        case 0:
            fmt.Println("Saindo do programa...")
            os.Exit(0)
        default:
            fmt.Println("Não conheço este comando!")
            os.Exit(-1)
       }

    }
}

func exibeIntroducao() {
    nome := "Graciani"
    versao := 1.1
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}

func exibeMenu()  {
    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}
func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O comando escolhido foi", comandoLido)

    return comandoLido
}

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    sites := []string{"http://random-status-code.herokuapp.com/", "http://www.alura.com.br/", "http://www.caelum.com.br/"}

    fmt.Println(sites)
    
    site := "http://www.alura.com.br/"
    resp, _ := http.Get(site) // acessando o site
    
    if resp.StatusCode == 200 {
        fmt.Println("Site", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
    }
}











