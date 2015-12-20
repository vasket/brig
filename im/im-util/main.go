package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/disorganizer/brig/im"
	colorlog "github.com/disorganizer/brig/util/log"
	"github.com/tsuibin/goxmpp2/xmpp"
)

func init() {
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	// Log pretty text
	log.SetFormatter(&colorlog.ColorfulLogFormatter{})
}

func main() {
	sendFlag := flag.Bool("send", false, "Send otr query")

	flag.Parse()

	var jid xmpp.JID
	var partnerJid xmpp.JID
	var password string

	aliceJid := xmpp.JID("alice@jabber.nullcat.de/laptop")
	bobJid := xmpp.JID("bob@jabber.nullcat.de/desktop")

	if *sendFlag {
		jid, partnerJid, password = aliceJid, bobJid, "ThiuJ9wesh"
	} else {
		jid, partnerJid, password = bobJid, aliceJid, "eecot3oXan"
	}

	client, err := im.NewClient(&im.Config{
		Jid:                  jid,
		Password:             password,
		TLSConfig:            tls.Config{ServerName: jid.Domain()},
		KeyPath:              "/tmp/otr.key." + password,
		FingerprintStorePath: "/tmp/otr.buddies." + password,
	})

	if err != nil {
		log.Fatalf("Could not create client: %v", err)
		return
	}

	defer client.Close()

	log.Infof("Partner is Online: %v", client.IsOnline(partnerJid))

	if *sendFlag {
		cnv, err := client.Dial(partnerJid)
		if err != nil {
			log.Errorf("Dial: %v", err)
			return
		}

		for i := 0; !cnv.Ended() && i < 10; i++ {
			log.Infof("Alice: PING %d", i)
			cnv.Write([]byte(fmt.Sprintf("PING %d", i)))

			msg, err := cnv.ReadMessage()
			log.Infof("Alice: RECV %d: %s/%v", i, msg, err)
			time.Sleep(2 * time.Second)
		}

		cnv.Close()
	} else {
		for {
			cnv := client.Listen()

			log.Println("Dial to", cnv.Jid)
			go func() {
				for i := 0; !cnv.Ended() && i < 10; i++ {
					msg, err := cnv.ReadMessage()
					log.Infof("Bob: RECV %d: %s/%v", i, msg, err)
					log.Infof("Bob: PONG %d", i)
					cnv.Write([]byte(fmt.Sprintf("PONG %d", i)))
				}
			}()
		}
	}
}
