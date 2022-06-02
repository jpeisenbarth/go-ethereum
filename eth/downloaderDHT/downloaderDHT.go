package downloaderdht

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

type Config struct {
	Database       ethdb.Database            // Database for direct sync insertions
	P2pServer			 *p2p.Server
}

type DownloaderDHT struct {
	headerHashes 	chan common.Hash		// hashs des header à dl 
	config				Config
}

// Crée notre nouvelle objet dlDHT
func New(headerChan chan common.Hash, conf *Config) *DownloaderDHT {
	dlDHT := &DownloaderDHT{
		headerHashes: headerChan,
		config: *conf,
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
	d.getClosestPeers(hash)

	// demander le bodies

	// ajouter au stockage
}

func (d *DownloaderDHT) getClosestPeers(hash common.Hash) {
	peer := d.findPeer(hash)

	log.Info("peer", "id", peer.ID(), "hash", hash.Hex())
	
	// ce connecter au pair
	// le retrouner
}

func (d *DownloaderDHT) findPeer(hash common.Hash) *enode.Node {
	server := d.config.P2pServer
	// rechercher dans les connections actuelle
	for _, peer := range(server.Peers()) {
		if bytes.Equal(peer.ID().Bytes(), hash[:]) {
			return peer.Node()
		}
	}
	// rechercher dans la table p2p
	for _, node := range(server.DiscV5.AllNodes()) {
		if bytes.Equal(node.ID().Bytes(), hash[:]) {
			return node
		}
	}
	// faire un lookup
	return server.DiscV5.Lookup(hash)[0]
}