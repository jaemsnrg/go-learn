package db

import (
	"context"
	"testing"

	util "github.com/jaemsnrg/go-learn/utils/testing"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) (Transfer, error) {
	account1, err := CreateRandomAccount()
	require.NoError(t, err)

	account2, err := CreateRandomAccount()
	require.NoError(t, err)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)

	return transfer, nil
}

func TestCreateTransfer(t *testing.T) {
	account1, err := CreateRandomAccount()
	require.NoError(t, err)

	account2, err := CreateRandomAccount()
	require.NoError(t, err)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 3; i++ {
		_, err := CreateRandomTransfer(t)
		require.NoError(t, err)
	}

	arg := ListTransfersParams{
		Limit:  3,
		Offset: 0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 3)
}

func TestGetTransfer(t *testing.T) {
	transfer, err := CreateRandomTransfer(t)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	transfer, err := CreateRandomTransfer(t)
	require.NoError(t, err)

	err = testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
}
