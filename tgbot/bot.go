package main

import (
	"ethereum-wallet/pkg"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		Token:  "1672120660:AAEburjQTE_B7RLPB2SFfLF406208C-6IOw",
		Poller: &tb.LongPoller{Timeout: 60 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
		b.Send(m.Sender, "You can use command\n" +
			"/createAccount yourownpassword : to create your own wallet\n" +
			"/checkAccounts yourownpassword : to check accounts relative to your telegram id\n" +
			"You should keep enter the same password for any time password is required.")
	})

	b.Handle("/createAccount", func(m *tb.Message) {
		userId := m.Sender.ID
		passphrase := strings.Trim(strings.TrimPrefix(m.Text, "/checkAccounts"), " ")

		w, err := wallet.NewWallet(strconv.Itoa(userId), passphrase)
		if err != nil {
			log.Println(err)
		}

		err = w.GenerateAccount(passphrase)
		if err != nil {
			log.Println(err)
		}

		returnStr := "Created successfully! \n"
		accounts := w.GetAllAccounts()
		for i, account := range(accounts) {
			returnStr += "Address of Account " + strconv.Itoa(i) +":" + account.Address.String() + "\n"
		}

		b.Send(m.Sender, returnStr)

	})

	b.Handle("/checkAccounts", func(m *tb.Message){
		userId := m.Sender.ID
		passphrase := strings.Trim(strings.TrimPrefix(m.Text, "/checkAccounts"), " ")

		w, err := wallet.NewWallet(strconv.Itoa(userId), passphrase)
		if err != nil {
			log.Println(err)
		}

		returnStr := ""
		accounts := w.GetAllAccounts()
		for i, account := range(accounts) {
			returnStr += "Address of Account " + strconv.Itoa(i) +":" + account.Address.String() + "\n"
		}

		b.Send(m.Sender, returnStr)
	})

	b.Start()
}

