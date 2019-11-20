package utils

import (
	"github.com/onionltd/oniontree-tools/pkg/types/service"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
	"strings"
)

func KeysToKeyRing(keys []service.PublicKey) (openpgp.EntityList, error) {
	el := openpgp.EntityList{}
	for idx, _ := range keys {
		block, err := armor.Decode(strings.NewReader(keys[idx].Value))
		if err != nil {
			return nil, err
		}
		e, err := openpgp.ReadEntity(packet.NewReader(block.Body))
		if err != nil {
			return nil, err
		}
		el = append(el, e)
	}
	return el, nil
}
