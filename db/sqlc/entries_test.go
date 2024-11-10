package db

import (
	"context"
	"testing"

	util "github.com/jaemsnrg/go-learn/utils/testing"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccountAndEntry() (Entry, error) {
	account, err := CreateRandomAccount()

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	return entry, err
}

func TestCreateEntry(t *testing.T) {

	entry, err := CreateRandomAccountAndEntry()

	require.NoError(t, err)
	require.NotEmpty(t, entry)
}

func TestGetEntry(t *testing.T) {
	entry, err := CreateRandomAccountAndEntry()
	require.NoError(t, err)

	selectedEntry, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	require.NotEmpty(t, selectedEntry)
}

func TestListEntries(t *testing.T) {

	var createdEntries []Entry
	for i := 0; i < 3; i++ {
		entry, err := CreateRandomAccountAndEntry()
		require.NoError(t, err)
		createdEntries = append(createdEntries, entry)
	}

	args := ListEntriesParams{
		Limit:  3,
		Offset: 0,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, entries, 3)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestUpdateEntry(t *testing.T) {
	entry, err := CreateRandomAccountAndEntry()
	require.NoError(t, err)

	params := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}

	updatedEntry, err := testQueries.UpdateEntry(context.Background(), params)
	require.NoError(t, err)

	require.NotEmpty(t, updatedEntry)
}

func TestDeleteEntry(t *testing.T) {
	entry, err := CreateRandomAccountAndEntry()
	require.NoError(t, err)

	err = testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)
}
