package main

import (
    "fmt"
    "os"
    "net/http"
    "time"
    "bufio"
)

const monitoramentos = 3
const delay = 5

func main() {

    exibeIntroducao()
    leSitesDoArquivo()

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
    fmt.Println("")

    return comandoLido
}

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    sites := []string{"http://random-status-code.herokuapp.com/", "http://www.alura.com.br/", "http://www.caelum.com.br/"}

    for i:=0; i<monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }
        time.Sleep(delay * time.Second)
        fmt.Println("")
        
    }

    fmt.Println("")
    
}

func testaSite(site string) {
    resp, err := http.Get(site) // acessando o site
    
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    if resp.StatusCode == 200 {
        fmt.Println("Site", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
    }
}

func leSitesDoArquivo() []string {
    var sites []string
    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    linha, err := leitor.ReadString('\n')
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }
    fmt.Println(linha)
    return sites
}









