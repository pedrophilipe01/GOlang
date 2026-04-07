package main

import (
	"fmt"
<<<<<<< HEAD
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

=======
	"net/http"
	"os"
)

>>>>>>> d482591be0f5c9949c95b33fd92dba17b212deb5
func main() {

	exibeIntroducao()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 3:
<<<<<<< HEAD
			fmt.Println("Saindo do programa")
=======
			fmt.Println("Saindo do programa...")
>>>>>>> d482591be0f5c9949c95b33fd92dba17b212deb5
			os.Exit(0)
		default:
			fmt.Println("Esse comando não existe")
			os.Exit(-1)
		}
	}
<<<<<<< HEAD

}

func exibeIntroducao() {
	versao := 1.1
	fmt.Println("")
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("")
=======
}

func exibeIntroducao() {
	nome := "Pedro"
	versao := 1.1
	fmt.Println("Olá, sr.", nome, "sua idade é")
	fmt.Println("Este programa está na versão", versao)
>>>>>>> d482591be0f5c9949c95b33fd92dba17b212deb5
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
<<<<<<< HEAD
	fmt.Println("")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("")
	fmt.Println("3 - Sair do Programa")
	fmt.Println("")
=======
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
>>>>>>> d482591be0f5c9949c95b33fd92dba17b212deb5
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
<<<<<<< HEAD
	fmt.Println("")
=======
>>>>>>> d482591be0f5c9949c95b33fd92dba17b212deb5

	return comandoLido
}

func iniciarMonitoramento() {
<<<<<<< HEAD
	fmt.Println("Monitorando...")
	//sites := []string{"https://www.youtube.com", "https://www.youtube.com", "https://www.youtube.com"}

	sites := leSitesDoArquivo()
=======
	fmt.Println("Monitorando...\n")

	sites := []string{"https://random-status-code.herokuapp.com/", "https://random-status-code.herokuapp.com/", "https://random-status-code.herokuapp.com/"}
>>>>>>> d482591be0f5c9949c95b33fd92dba17b212deb5

	for i := 0; i < 5; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
<<<<<<< HEAD
		time.Sleep(delay * time.Second)
		fmt.Println("")
=======
>>>>>>> d482591be0f5c9949c95b33fd92dba17b212deb5
	}
	fmt.Println("")
}

func testaSite(site string) {
<<<<<<< HEAD
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
=======
	resp, _ := http.Get(site)
>>>>>>> d482591be0f5c9949c95b33fd92dba17b212deb5

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
<<<<<<< HEAD
func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(arquivo)

	return sites
}

//Tratando erros
=======
>>>>>>> d482591be0f5c9949c95b33fd92dba17b212deb5
