package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/Finnkell/codepix/codepix/domain/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModel_NewTransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "406523489"
	ownerName := "Eric"
	account, _ := model.NewAccount(bank, ownerName, accountNumber)

	accountNumberDestination := "4065123873"
	ownerName := "Lurdes"
	accountDestination,  _ := model.NewAccount(bank, ownerName, accountNumberDestination)

	kind := "cpf"
	key := "1234567891"
	pixKey, _ = model.NewPixKey(kind, key, accountDestination)
}
