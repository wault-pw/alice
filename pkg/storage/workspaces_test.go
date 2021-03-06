package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wault-pw/alice/pkg/domain"
)

func TestStorage_CreateWorkspace(t *testing.T) {
	store, savepoint := MustNewStore(t)
	t.Cleanup(savepoint.Flush)

	t.Run("it works", func(t *testing.T) {
		ctx := context.Background()
		workspace := mustBuildWorkspace(t, store, &domain.Workspace{})
		userWorkspace := mustBuildUserWorkspace(t, store, &domain.UserWorkspace{})

		err := store.CreateWorkspace(ctx, userWorkspace, workspace)
		require.NoError(t, err)

		require.NotEmpty(t, workspace.ID.String)
		require.NotEmpty(t, userWorkspace.ID.String)
	})
}

func TestStorage_DeleteWorkspace(t *testing.T) {
	store, savepoint := MustNewStore(t)
	t.Cleanup(savepoint.Flush)

	t.Run("it works", func(t *testing.T) {
		ctx := context.Background()

		workspace01 := mustCreateWorkspace(t, store, &domain.Workspace{})
		workspace02 := mustCreateWorkspace(t, store, &domain.Workspace{})

		err := store.DeleteWorkspace(ctx, workspace01.ID.String)
		require.NoError(t, err)

		_, err11 := store.FindWorkspace(ctx, workspace01.ID.String)
		_, err12 := store.FindWorkspace(ctx, workspace02.ID.String)

		require.ErrorIs(t, err11, ErrNotFound)
		require.NoError(t, err12)
	})
}

func TestStorage_UpdateWorkspace(t *testing.T) {
	store, savepoint := MustNewStore(t)
	t.Cleanup(savepoint.Flush)

	t.Run("it works", func(t *testing.T) {
		ctx := context.Background()

		workspace01 := mustCreateWorkspace(t, store, &domain.Workspace{})
		workspace02 := mustCreateWorkspace(t, store, &domain.Workspace{})
		title := []byte{1}

		err := store.UpdateWorkspace(ctx, workspace01.ID.String, title)
		require.NoError(t, err)

		workspace11, err11 := store.FindWorkspace(ctx, workspace01.ID.String)
		workspace12, err12 := store.FindWorkspace(ctx, workspace02.ID.String)
		require.NoError(t, err11)
		require.NoError(t, err12)

		require.Equal(t, title, workspace11.TitleEnc.Bytea)
		require.Equal(t, workspace02.TitleEnc.Bytea, workspace12.TitleEnc.Bytea)
	})
}

func mustBuildWorkspace(t *testing.T, storage *Storage, input *domain.Workspace) *domain.Workspace {
	out := &domain.Workspace{
		TitleEnc: domain.NewEmptyBytes([]byte("TitleEnc")),
	}

	if input.TitleEnc.Valid {
		out.TitleEnc = input.TitleEnc
	}

	return out
}

func mustCreateWorkspace(t *testing.T, storage *Storage, input *domain.Workspace) *domain.Workspace {
	ctx := context.Background()
	output := mustBuildWorkspace(t, storage, input)
	if err := storage.insertWorkspace(ctx, storage.db, output); err != nil {
		t.Errorf("cant create factory workspace: %s", err)
		t.FailNow()
	}
	return output
}
