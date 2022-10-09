package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()

	//fmt.Println("O tipo da variável idade é", reflect.TypeOf(versao))
	fmt.Println("")

	for {

		exibeNomes()

		fmt.Println("")

		exibeMenu()

		fmt.Println("")

		comando := leComando()

		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo Logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do programa...")
		// } else {
		// 	fmt.Println("Não conheço este comando")
		// }

		fmt.Println("")

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func iniciarMonitoramento() {

	fmt.Println("Monitorando...")

	// var sites [4]string
	// sites[0] = "https://random-status-code.herokuapp.com/"
	// sites[1] = "https://www.alura.com.br"
	// sites[2] = "https://www.caelum.com.br"
	//fmt.Println(reflect.TypeOf(sites))

	// sites := []string{"https://random-status-code.herokuapp.com/",
	// 	"https://www.alura.com.br", "https://www.caelum.com.br"}

	sites := leSitesDoArquivo()

	// for i := 0; i < len(sites); i++ {
	//     fmt.Println(sites[i])
	// }

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func exibeIntroducao() {
	//formas de declaração de variaveis
	var nome string = "João Paulo"
	var idade = 25
	versao := 1.1

	fmt.Println("Olá Mundo!!")
	fmt.Println("Olá ", nome)
	fmt.Println("Sua idade é ", idade)
	fmt.Println("A versão é ", versao)
}

func exibeMenu() {
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

// slice = list
func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida")
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	//string(arquivo) -> convertendo os byts para string
	//fmt.Println(string(arquivo))

	leitor := bufio.NewReader(arquivo)

	for {

		linha, err := leitor.ReadString('\n')

		sites = append(sites, strings.TrimSpace(linha))

		if err == io.EOF {
			break
		} else if err != nil && err != io.EOF {
			fmt.Println("Ocorreu um erro:", err)
		}
	}

	defer arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString("Data: " + time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	defer arquivo.Close()
}

func imprimeLogs() {
	fmt.Println("Exibindo Logs...")

	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
