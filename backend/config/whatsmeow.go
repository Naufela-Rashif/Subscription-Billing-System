package config

import (
	"context"
	"fmt"
	"os"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
)

func Whatsmeow() {
	// Inisialisasi logger untuk database
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Inisialisasi sqlstore menggunakan SQLite (file: whatsmeow.db)
	container, err := sqlstore.New("sqlite3", "file:whatsmeow.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}

	// Ambil device store (otomatis membuat tabel jika belum ada)
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}

	// Inisialisasi logger untuk client
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	// Buat instance Whatsmeow client
	client := whatsmeow.NewClient(deviceStore, clientLog)

	// Proses login/pairing
	if client.Store.ID == nil {
		// Jika belum ada sesi, dapatkan channel QR code
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		// Mendengarkan channel untuk event QR code
		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("Silakan scan QR code berikut:")
				// Mengenerate QR code dan tampilkan di terminal
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Event pairing:", evt.Event)
			}
		}
	} else {
		// Jika sudah ada sesi, langsung reconnect
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}

	// Aplikasi akan tetap berjalan untuk menerima event hingga dihentikan
	select {}
}
