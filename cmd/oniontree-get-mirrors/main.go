package main

import (
	"errors"
	"flag"
	"fmt"
	goomg "github.com/onionltd/go-omg"
	"github.com/onionltd/oniontree-omg/pkg/utils"
	"github.com/onionltd/oniontree-tools/pkg/oniontree"
	"golang.org/x/crypto/openpgp"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func exitError(msg string) {
	fmt.Printf("%s: %s\n", os.Args[0], msg)
	os.Exit(1)
}

const warningMessage = "WARNING: Message signature verification DISABLED!"

func main() {
	id := flag.String("id", "", "Onion service ID")
	replace := flag.Bool("replace", false, "Replace existing URLs")
	timeout := flag.Duration("timeout", 30*time.Second, "HTTP request timeout")
	verifySig := flag.Bool("verify-signature", true, "Enable signature verification")
	flag.Parse()

	if *id == "" {
		exitError("id not specified")
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	onionTree, err := oniontree.Open(wd)
	if err != nil {
		panic(err)
	}

	s, err := onionTree.Get(*id)
	if err != nil {
		if err == oniontree.ErrIdNotExists {
			exitError(err.Error())
		}
		panic(err)
	}

	entities := openpgp.EntityList{}
	if *verifySig {
		var err error
		entities, err = utils.KeysToKeyRing(s.PublicKeys)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println(strings.Repeat("=", len(warningMessage)))
		fmt.Println(warningMessage)
		fmt.Println(strings.Repeat("=", len(warningMessage)))
	}

	c := goomg.NewClient(&http.Client{
		Timeout: *timeout,
	})

	changed := false
	for _, url := range s.URLs {
		var urls []string
		msg, err := c.GetMirrorsMessage(url)
		if err != nil {
			goto printlog
		}
		if *verifySig {
			if _, err = msg.VerifySignature(entities); err != nil {
				goto printlog
			}
		}
		urls, err = msg.List()
		if err != nil {
			goto printlog
		}
		if len(urls) == 0 {
			err = errors.New("no mirrors found")
			goto printlog
		}
		changed = true
		if *replace {
			s.SetURLs(urls...)
		} else {
			s.AddURLs(urls...)
		}
	printlog:
		if err == nil {
			err = errors.New("ok")
		}
		log.Printf("%s: %s", url, err)
	}

	if changed {
		if err := onionTree.Edit(*id, s); err != nil {
			panic(err)
		}
	}
}
