package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/Finnkell/codepix/codepix/domain/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModel_NewAccount(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)

	accountNumber := "406572783"
	ownerName := "Eric"
	account, arr := NewAccount(bank, accountNumber, ownerName)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(account.ID))
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.BankID, bank.ID)

	_, err = mode.NewAccount(bank, "", ownerName)
	require.NotNil(t, err)
}
