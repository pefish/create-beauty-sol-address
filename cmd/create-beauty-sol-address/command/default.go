package command

import (
	"strings"
	"sync"

	"github.com/pefish/create-beauty-sol-address/pkg/global"

	"github.com/gagliardetto/solana-go"
	"github.com/pefish/go-commander"
)

type PersistenceDataType struct {
}

type DefaultCommand struct {
	persistenceData *PersistenceDataType
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{
		persistenceData: &PersistenceDataType{},
	}
}

func (dc *DefaultCommand) Config() interface{} {
	return &global.GlobalConfig
}

func (dc *DefaultCommand) Data() interface{} {
	return dc.persistenceData
}

func (dc *DefaultCommand) Init(command *commander.Commander) error {
	return nil
}

func (dc *DefaultCommand) OnExited(command *commander.Commander) error {
	return nil
}

func (dc *DefaultCommand) Start(command *commander.Commander) error {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-command.Ctx.Done():
					return
				default:
					account := solana.NewWallet()
					address := account.PublicKey().String()
					if isBeauty(address) {
						command.CancelFunc()
						command.Logger.InfoF("address: %s, priv: %s", address, account.PrivateKey.String())
						return
					}
				}
			}

		}()
	}

	wg.Wait()

	return nil
}

func isBeauty(address string) bool {
	prefix6 := address[:6]
	if strings.EqualFold(prefix6, "solsol") {
		return true
	}
	if prefix6 == "666666" || prefix6 == "888888" {
		return true
	}

	prefix5 := address[:5]
	if prefix5 == "trade" || prefix5 == "TRADE" {
		return true
	}
	// if prefix5 == "66666" || prefix5 == "88888" {
	// 	return true
	// }

	// prefix4 := address[:4]
	// if prefix4 == "8888" {
	// 	return true
	// }
	// if prefix4 == "6666" {
	// 	return true
	// }

	prefix3 := address[:3]
	if prefix3 == "SOL" || prefix3 == "sol" {
		return true
	}

	return false
}
