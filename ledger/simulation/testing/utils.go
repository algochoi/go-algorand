package simulationtesting

import (
	"testing"

	"github.com/algorand/go-algorand/config"
	"github.com/algorand/go-algorand/crypto"
	"github.com/algorand/go-algorand/data"
	"github.com/algorand/go-algorand/data/basics"
	"github.com/algorand/go-algorand/data/transactions"
	"github.com/algorand/go-algorand/ledger"
	ledgertesting "github.com/algorand/go-algorand/ledger/testing"
	"github.com/algorand/go-algorand/logging"
	"github.com/algorand/go-algorand/protocol"
	"github.com/stretchr/testify/require"
)

// ==============================
// > Simulation Test Helpers
// ==============================

type Account struct {
	Addr     basics.Address
	Sk       *crypto.SignatureSecrets
	AcctData basics.AccountData
}

func PrepareSimulatorTest(t *testing.T) (l *data.Ledger, accounts []Account, makeTxnHeader func(sender basics.Address) transactions.Header) {
	genesisInitState, keys := ledgertesting.GenerateInitState(t, protocol.ConsensusCurrentVersion, 100)

	// Prepare ledger
	const inMem = true
	cfg := config.GetDefaultLocal()
	cfg.Archival = true
	log := logging.TestingLog(t)
	log.SetLevel(logging.Warn)
	realLedger, err := ledger.OpenLedger(log, t.Name(), inMem, genesisInitState, cfg)
	require.NoError(t, err, "could not open ledger")

	l = &data.Ledger{Ledger: realLedger}
	require.NotNil(t, l)

	// Reformat accounts
	accounts = make([]Account, len(keys)-2) // -2 for pool and sink accounts
	i := 0
	for addr, key := range keys {
		if addr == ledgertesting.PoolAddr() || addr == ledgertesting.SinkAddr() {
			continue
		}

		acctData := genesisInitState.Accounts[addr]
		accounts[i] = Account{
			Addr:     addr,
			Sk:       key,
			AcctData: acctData,
		}
		i++
	}

	// txn header generator
	hdr, err := l.BlockHdr(l.Latest())
	require.NoError(t, err)
	makeTxnHeader = func(sender basics.Address) transactions.Header {
		return transactions.Header{
			Fee:         basics.MicroAlgos{Raw: 1000},
			FirstValid:  hdr.Round,
			GenesisID:   hdr.GenesisID,
			GenesisHash: hdr.GenesisHash,
			LastValid:   hdr.Round + basics.Round(1000),
			Note:        []byte{240, 134, 38, 55, 197, 14, 142, 132},
			Sender:      sender,
		}
	}

	return
}
