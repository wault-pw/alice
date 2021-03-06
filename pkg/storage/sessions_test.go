package storage

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/wault-pw/alice/lib/jwt_mock"
	"github.com/wault-pw/alice/pkg/domain"
)

func TestStorage_IssueSession(t *testing.T) {
	store, savepoint := MustNewStore(t)
	t.Cleanup(savepoint.Flush)

	t.Run("it works", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		ots := jwt_mock.NewMockOts(ctrl)
		defer ctrl.Finish()

		ots.EXPECT().Marshall().Return("foo", nil)
		ots.EXPECT().SetJti(gomock.Any()).Return(ots)
		ots.EXPECT().SetExp(gomock.Any()).Return(ots)

		session, token, err := store.IssueSession(context.Background(), ots)
		require.NoError(t, err)
		require.Equal(t, "foo", token)
		require.NotEmpty(t, session.Jti.String)
		require.Equal(t, true, session.TimeTo.Time.After(session.TimeFrom.Time))
		require.Empty(t, session.UserID.String)
		require.Empty(t, session.CandidateID.String)
	})
}

func TestStorage_RetrieveSession(t *testing.T) {
	store, savepoint := MustNewStore(t)
	t.Cleanup(savepoint.Flush)

	type args struct {
		desc    string
		before  func(ots *jwt_mock.MockOts) *domain.Session
		wantErr bool
	}

	tests := []args{
		{
			desc: "when JIT valid",
			before: func(ots *jwt_mock.MockOts) *domain.Session {
				session := mustCreateSession(t, store, &domain.Session{})
				ots.EXPECT().Unmarshall(gomock.Any()).Return(session.Jti.String, nil)
				return session
			},
			wantErr: false,
		}, {
			desc: "when JIT invalid",
			before: func(ots *jwt_mock.MockOts) *domain.Session {
				ots.EXPECT().Unmarshall(gomock.Any()).Return("", errors.New("foo"))
				return nil
			},
			wantErr: true,
		}, {
			desc: "when not found invalid",
			before: func(ots *jwt_mock.MockOts) *domain.Session {
				ots.EXPECT().Unmarshall(gomock.Any()).Return(domain.NewUUID(), nil)
				return nil
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			ots := jwt_mock.NewMockOts(ctrl)
			defer ctrl.Finish()

			want := tt.before(ots)
			session, err := store.RetrieveSession(context.Background(), ots, "")

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.Equal(t, want.Jti.String, session.Jti.String)
		})
	}
}

func TestStorage_CandidateSession(t *testing.T) {
	store, savepoint := MustNewStore(t)
	t.Cleanup(savepoint.Flush)

	t.Run("it works", func(t *testing.T) {
		ctx := context.Background()
		srp := []byte("srp")
		user := mustCreateUser(t, store, &domain.User{})

		session01 := mustCreateSession(t, store, &domain.Session{})
		session02 := mustCreateSession(t, store, &domain.Session{})

		err := store.CandidateSession(ctx, session01.Jti.String, user.ID.String, srp)
		require.NoError(t, err)

		session11, err11 := store.FindSession(ctx, session01.Jti.String)
		session12, err12 := store.FindSession(ctx, session02.Jti.String)
		require.NoError(t, err11)
		require.NoError(t, err12)

		require.Equal(t, session11.CandidateID.String, user.ID.String)
		require.Equal(t, session11.SrpState.Bytea, srp)

		require.Empty(t, session12.CandidateID.String)
		require.Empty(t, session12.SrpState.Bytea)
	})
}

func TestStorage_NominateSession(t *testing.T) {
	store, savepoint := MustNewStore(t)
	t.Cleanup(savepoint.Flush)

	t.Run("it works", func(t *testing.T) {
		ctx := context.Background()
		user := mustCreateUser(t, store, &domain.User{})

		session01 := mustCreateSession(t, store, &domain.Session{
			CandidateID: user.ID,
			SrpState:    domain.NewEmptyBytes([]byte("foo")),
		})
		session02 := mustCreateSession(t, store, &domain.Session{})

		err := store.NominateSession(ctx, session01.Jti.String)
		require.NoError(t, err)

		session11, err11 := store.FindSession(ctx, session01.Jti.String)
		session12, err12 := store.FindSession(ctx, session02.Jti.String)
		require.NoError(t, err11)
		require.NoError(t, err12)

		require.Equal(t, session11.UserID.String, user.ID.String, "sets userID")
		require.Empty(t, session11.CandidateID.String, "clear candidateID")
		require.Empty(t, session11.SrpState.Bytea, "clear SRP state")

		require.Empty(t, session12.UserID.String)
		require.Empty(t, session12.CandidateID.String)
	})
}

func TestStorage_DeleteSession(t *testing.T) {
	store, savepoint := MustNewStore(t)
	t.Cleanup(savepoint.Flush)

	t.Run("it works", func(t *testing.T) {
		ctx := context.Background()

		session01 := mustCreateSession(t, store, &domain.Session{})
		session02 := mustCreateSession(t, store, &domain.Session{})

		err := store.DeleteSession(ctx, session01.Jti.String)
		require.NoError(t, err)

		_, err11 := store.FindSession(ctx, session01.Jti.String)
		_, err12 := store.FindSession(ctx, session02.Jti.String)
		require.ErrorIs(t, err11, ErrNotFound)
		require.NoError(t, err12)
	})
}

func TestStorage_DeleteUserSessionExcept(t *testing.T) {
	store, savepoint := MustNewStore(t)
	t.Cleanup(savepoint.Flush)

	t.Run("it works", func(t *testing.T) {
		ctx := context.Background()

		user1 := mustCreateUser(t, store, &domain.User{})
		user2 := mustCreateUser(t, store, &domain.User{})

		session01 := mustCreateSession(t, store, &domain.Session{UserID: user1.ID})
		session02 := mustCreateSession(t, store, &domain.Session{UserID: user1.ID})
		session03 := mustCreateSession(t, store, &domain.Session{UserID: user2.ID})

		err := store.DeleteUserSessionExcept(ctx, user1.ID.String, session01.Jti.String)
		require.NoError(t, err)

		_, err11 := store.FindSession(ctx, session01.Jti.String)
		_, err12 := store.FindSession(ctx, session02.Jti.String)
		_, err13 := store.FindSession(ctx, session03.Jti.String)

		require.NoError(t, err11, "keeps current session")
		require.ErrorIs(t, err12, ErrNotFound, "deletes stale sessions")
		require.NoError(t, err13, "keeps sessions for foreign user")
	})
}

func TestStorage_countLoginAttempts(t *testing.T) {
	store, savepoint := MustNewStore(t)
	t.Cleanup(savepoint.Flush)

	t.Run("it works", func(t *testing.T) {
		ctx := context.Background()
		user1 := mustCreateUser(t, store, &domain.User{})
		user2 := mustCreateUser(t, store, &domain.User{})

		session := mustCreateSession(t, store, &domain.Session{
			CandidateID: user1.ID,
			TimeFrom:    domain.NewEmptyTime(time.Now().Add(-5 * time.Minute)),
		})
		_ = mustCreateSession(t, store, &domain.Session{
			CandidateID: user2.ID,
		})

		counter01, err01 := store.countLoginAttempts(ctx, savepoint, session.CandidateID.String, time.Minute)
		counter02, err02 := store.countLoginAttempts(ctx, savepoint, session.CandidateID.String, 6*time.Minute)

		require.NoError(t, err01)
		require.NoError(t, err02)

		require.Equal(t, 0, counter01)
		require.Equal(t, 1, counter02)
	})
}

func TestStorage_countOtpAttempts(t *testing.T) {
	store, savepoint := MustNewStore(t)
	t.Cleanup(savepoint.Flush)

	t.Run("it works", func(t *testing.T) {
		ctx := context.Background()
		user := mustCreateUser(t, store, &domain.User{})
		session := mustCreateSession(t, store, &domain.Session{
			UserID:      user.ID,
			OtpAttempts: domain.NewNullInt64(3),
			TimeFrom:    domain.NewEmptyTime(time.Now().Add(-5 * time.Minute)),
		})
		_ = mustCreateSession(t, store, &domain.Session{})

		counter01, err01 := store.countOtpAttempts(ctx, savepoint, session.Jti.String, time.Minute)
		counter02, err02 := store.countOtpAttempts(ctx, savepoint, session.Jti.String, 6*time.Minute)

		require.NoError(t, err01)
		require.NoError(t, err02)

		require.Equal(t, 0, counter01)
		require.Equal(t, int(session.OtpAttempts.Int64), counter02)
	})
}

func TestStorage_MakeOtpAttempt(t *testing.T) {
	store, savepoint := MustNewStore(t)
	t.Cleanup(savepoint.Flush)

	t.Run("it works", func(t *testing.T) {
		ctx := context.Background()
		session01 := mustCreateSession(t, store, &domain.Session{})
		session02 := mustCreateSession(t, store, &domain.Session{})

		err01 := store.MakeOtpAttempt(ctx, session01.Jti.String)
		err02 := store.MakeOtpAttempt(ctx, session01.Jti.String)

		require.NoError(t, err01)
		require.NoError(t, err02)

		session11, err11 := store.FindSession(ctx, session01.Jti.String)
		session12, err12 := store.FindSession(ctx, session02.Jti.String)

		require.NoError(t, err11)
		require.NoError(t, err12)

		require.Equal(t, int64(2), session11.OtpAttempts.Int64)
		require.Equal(t, session02.OtpAttempts.Int64, session12.OtpAttempts.Int64)
	})
}

func mustBuildSession(t *testing.T, storage *Storage, input *domain.Session) *domain.Session {
	out := &domain.Session{
		Jti:         domain.NewEmptyString(domain.NewUUID()),
		TimeFrom:    domain.NewEmptyTime(time.Now()),
		OtpAttempts: domain.NewNullInt64(0),
		TimeTo:      domain.NewEmptyTime(time.Now().Add(24 * time.Hour)),
	}

	if input.Jti.Valid {
		out.Jti = input.Jti
	}

	if input.OtpAttempts.Valid {
		out.OtpAttempts = input.OtpAttempts
	}

	if input.UserID.Valid {
		out.UserID = input.UserID
	}

	if input.CandidateID.Valid {
		out.CandidateID = input.CandidateID
	}

	if input.TimeFrom.Valid {
		out.TimeFrom = input.TimeFrom
	}

	if input.TimeTo.Valid {
		out.TimeTo = input.TimeTo
	}

	return out
}

func mustCreateSession(t *testing.T, storage *Storage, input *domain.Session) *domain.Session {
	ctx := context.Background()
	output := mustBuildSession(t, storage, input)
	if err := storage.insertSession(ctx, storage.db, output); err != nil {
		t.Errorf("cant create factory session: %s", err)
		t.FailNow()
	}
	return output
}
