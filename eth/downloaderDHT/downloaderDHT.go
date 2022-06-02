package downloaderdht

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type DownloaderDHT struct {
	headerHashes chan common.Hash		// hashs des header à dl 
}

// Crée notre nouvelle objet dlDHT
func New(headerChan chan common.Hash) *DownloaderDHT {
	dlDHT := &DownloaderDHT{
		headerHashes: headerChan,
	}

	go dlDHT.run()

	return dlDHT
}

// Observe le channel et quand nouveau hash exec
func (d *DownloaderDHT) run() {

	for {
		select {
		case hash := <-d.headerHashes:
			go d.main(hash)
		}
	}

}

func (d *DownloaderDHT) main(hash common.Hash) {
	log.Info("Oui", "hash", hash)
}