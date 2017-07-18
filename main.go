/*
Copyright (c) 2017 Dino Sauro
*/
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/text/encoding/charmap"
)

var config Config
var elvira Elvira

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:8081" will only work if you are accessing the server from your local machine.
	http.Redirect(w, r, "https://127.0.0.1:443"+r.RequestURI, http.StatusMovedPermanently)
}

//handlerStatus salva status garagem
func handlerStatus(w http.ResponseWriter, r *http.Request) {

	icon := " \U0001f618"

	status := ""
	var err error
	operacao := r.FormValue("operacao")
	if operacao != "" {
		if operacao == "abrirGaragemH4ck3r" {
			status = "aberta"
		} else if operacao == "fecharGaragemH4ck3r" {
			status = "fechada"
			icon = " \U0001f614"
		}
		d1 := []byte(status)
		err = ioutil.WriteFile("status.txt", d1, 0644)
		if err != nil {
			fmt.Println("não foi possivel gravar o status")
		}
		elvira.sendMsg(config.GaragemChatId, "Queridos a Garagem Hacker está <b>"+status+"</b>"+icon)
	} else {
		status, err = getStatusGaragem()
		if err != nil {
			status = "Error GetStatus"
		}
	}

	io.WriteString(w, status)

}

func getStatusGaragem() (string, error) {

	path := "status.txt"
	contents := ""

	file, err := os.Open(path)
	if err != nil {
		return string(contents), err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) > 0 {
		contents = lines[0]
	} else {
		err = errors.New("Erro ao pesquisar status")
	}

	return string(contents), err
}
func getTempo() (string, error) {
	response, err := http.Get("http://www.cptec.inpe.br/cidades/tempo/227")
	var dados string

	doc, err := html.Parse(charmap.ISO8859_1.NewDecoder().Reader(response.Body))
	if err != nil {
		return dados, err
	}

	var ff func(*html.Node)
	ff = func(n *html.Node) {
		if n.Type == html.TextNode {

			dd := strings.Replace(strings.Replace(strings.Replace(n.Data, "\t", "", -1), "\n", "", -1), "  ", "", -1)

			n := len(dd)

			if n > 1 {

				if n > 12 {
					dd = "\n" + dd
				} else {
					dd = " " + dd
				}
				dados = dados + dd
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			ff(c)
		}
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" {
					if a.Val == "cond deg_azul" {

						ff(n)
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return strings.Title(strings.ToLower(dados)), err
}
func getDolar() (string, string, error) {
	response, err := http.Get("https://ptax.bcb.gov.br/ptax_internet/consultarUltimaCotacaoDolar.do")
	var contents []byte
	if err == nil {
		defer response.Body.Close()
		contents, err = ioutil.ReadAll(response.Body)
	}

	tratamento := strings.Split(string(contents), "</tr")
	tratamento2 := strings.Split(tratamento[1], "<td")

	compra := strings.Replace(strings.Replace(strings.Replace(tratamento2[2], "</td>", "", -1), "align=\"right\">", "", -1), " ", "", -1)
	venda := strings.Replace(strings.Replace(strings.Replace(tratamento2[3], "</td>", "", -1), "align=\"right\">", "", -1), " ", "", -1)
	compra = strings.Replace(strings.Replace(strings.Replace(compra, "\t", "", -1), "\n", "", -1), "  ", "", -1)
	venda = strings.Replace(strings.Replace(strings.Replace(venda, "\t", "", -1), "\n", "", -1), "  ", "", -1)
	return compra, venda, err
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
	defer r.Body.Close()
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	update := Update{}
	err = json.Unmarshal(contents, &update)
	if err != nil {
		fmt.Println(err)
	}
	msg := update.Message

	if msg.LeftChatMember.ID > 0 {
		elvira.sendMsg(msg.Chat.ID, "Querid@ "+msg.LeftChatMember.FirstName+" porque saiu do grupo? \U0001f614")
	} else if msg.NewChatMember.ID > 0 {
		elvira.sendMsg(msg.Chat.ID, "Querid@ "+msg.NewChatMember.FirstName+" bem vind@ ao grupo. \U0001f618. Para saber se a garagemhacker está aberta envie /status ")
	} else if strings.Contains(msg.Text, "oi") {
		elvira.sendMsg(msg.Chat.ID, "Ola querido "+msg.From.FirstName+", o que você quer.")
	} else if strings.EqualFold(msg.Text, "@ElviraBot") {
		elvira.sendMsg(msg.Chat.ID, "Querido "+msg.From.FirstName+" pode falar com a rainha das trevas.")
	} else if strings.Contains(msg.Text, "como vai") {
		elvira.sendMsg(msg.Chat.ID, msg.From.FirstName+", vou Bem e você.")
	} else if strings.Contains(msg.Text, "garagemhacker") {
		elvira.sendMsg(msg.Chat.ID, "A <b>GaragemHacker</b> é um espaço físico aberto para qualquer pessoa frequentar, não existe nenhum pré-requisito ou conhecimento para se integrar ao espaço. Basta ser ativo e ter interesse. As atividades são gratuítas, mas tambem podem existir atividades ou workshops que são pagos para auxiliar o espaço a se manter e crescer. \n http://garagemhacker.org \n\n <b>Endereço:</b> Rua Antônio Chella, 431 - Bom Retiro, Curitiba - PR, 80520-460")
	} else if strings.EqualFold(msg.Text, "/garagem") || strings.EqualFold(msg.Text, "/garagem@ElviraBot") {
		elvira.sendMsg(msg.Chat.ID, "A <b>GaragemHacker</b> é um espaço físico aberto para qualquer pessoa frequentar, não existe nenhum pré-requisito ou conhecimento para se integrar ao espaço. Basta ser ativo e ter interesse. As atividades são gratuítas, mas tambem podem existir atividades ou workshops que são pagos para auxiliar o espaço a se manter e crescer. \n http://garagemhacker.org \n\n <b>Endereço:</b> Rua Antônio Chella, 431 - Bom Retiro, Curitiba - PR, 80520-460")
	} else if strings.EqualFold(msg.Text, "/endereco") || strings.EqualFold(msg.Text, "/endereco@ElviraBot") {
		elvira.sendMsg(msg.Chat.ID, "<b>Endereço:</b> Rua Antônio Chella, 431 - Bom Retiro, Curitiba - PR, 80520-460")
	} else if strings.EqualFold(msg.Text, "/status") || strings.EqualFold(msg.Text, "/status@ElviraBot") {
		status, err := getStatusGaragem()

		if err != nil {
			elvira.sendMsg(msg.Chat.ID, "Queridos Erro ao buscar <b>o status da garagem</b>")
		} else {
			if strings.Compare("aberta", status) == 0 {
				status = "aberta  \U0001f618"
				elvira.sendMsg(msg.Chat.ID, "Queridos a Garagem Hacker está <b>"+status+"</b>")
			} else if strings.Compare("fechada", status) == 0 {
				status = "fechada  \U0001f614"
				elvira.sendMsg(msg.Chat.ID, "Queridos a Garagem Hacker está <b>"+status+"</b>")
			} else {
				elvira.sendMsg(msg.Chat.ID, "Queridos não consegui checar o status da <b>Garagem Hacker</b> ")
			}
		}
	} else if strings.EqualFold(msg.Text, "/tempo") || strings.EqualFold(msg.Text, "/tempo@ElviraBot") {
		tempo, err := getTempo()
		if err != nil {
			elvira.sendMsg(msg.Chat.ID, "Queridos Erro ao buscar <b>o Tempo</b>")
		} else {
			elvira.sendMsg(msg.Chat.ID, tempo)
		}
	} else if strings.EqualFold(msg.Text, "/dolar") || strings.EqualFold(msg.Text, "/dolar@ElviraBot") {
		compra, venda, err := getDolar()
		if err == nil {
			elvira.sendMsg(msg.Chat.ID, "O dolár está cotado a: \n <b>R$:"+compra+"</b> para compra e \n <b>R$:"+venda+"</b> para venda.")
		}
	} else if strings.EqualFold(msg.Text, "/help") || strings.EqualFold(msg.Text, "/help@ElviraBot") {
		elvira.sendMsg(msg.Chat.ID, "Elvira rainha das trevas \n\n/garagem - Mais informações sobre a GaragemHacker\n/status - Saiba se a Garagem Hacker está aberta ou fechada\n/dolar - Veja o tamanho da facada antes de comprar a sua plaquinha, módulo, shield, etc...\n/endereco - Veja o endereço da Garagem Hacker\n/talk - Diga um olá para a rainha das trevas\n\ngaragemhacker.com")
	} else if strings.EqualFold(msg.Text, "@ElviraBot nudes") {
		elvira.sendMsg(msg.Chat.ID, msg.From.FirstName+" vou tirar uma especial para você.")
		//elvira.sendPhoto(msg.Chat.ID, "giphy.gif")

	} else {
		elvira.sendMsg(msg.Chat.ID, "Querid@ "+msg.From.FirstName+" não entendi sobre oque você quer falar.")
	}

}

func main() {
	fmt.Println("start bot")
	if err := readConfig(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	elvira.Token = config.BotToken

	go func() {
		http.HandleFunc(config.ReceiveStatusPath, handlerStatus)
		err := http.ListenAndServe(":8090", nil)

		if err != nil {
			fmt.Println(err)
		}
	}()

	http.HandleFunc(config.SecretPath, handler)
	// Start the HTTPS server in a goroutine
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	// Start the HTTP server and redirect all incoming connections to HTTPS
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(err)
}

// read a config.json file and popule return
func readConfig() error {
	raw, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, &config)
}
