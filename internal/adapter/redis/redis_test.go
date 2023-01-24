package redis

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func redisTestHelper(t *testing.T) (redismock.ClientMock, *Client) {
	t.Helper()
	t.Parallel()

	db, mock := redismock.NewClientMock()
	redisUsecase := New(db)

	return mock, redisUsecase
}

func Test_SetIfNotExist(t *testing.T) {
	mock, rc := redisTestHelper(t)

	type args struct {
		key  string
		data interface{}
		ttl  time.Duration
	}
	tests := []struct {
		name    string
		args    args
		doMock  func(args args)
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Error",
			args: args{
				key:  "123456789",
				data: "Success",
				ttl:  30 * time.Second,
			},
			wantErr: true,
			doMock: func(args args) {
				mock.ExpectSetNX(args.key, args.data, args.ttl).SetErr(redis.ErrClosed)
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args)
			}

			got, err := rc.SetIfNotExist(context.TODO(), tt.args.key, tt.args.data, tt.args.ttl)

			assert.Equal(t, tt.wantErr, (err != nil))
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_Get(t *testing.T) {
	mock, rc := redisTestHelper(t)

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		doMock  func(args args)
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			args: args{
				key: "123456789",
			},
			wantErr: false,
			want:    "test",
			doMock: func(args args) {
				mock.ExpectGet(args.key).SetVal("test")
			},
		},
		{
			name: "Error",
			args: args{
				key: "123456789",
			},
			wantErr: true,
			want:    "",
			doMock: func(args args) {
				mock.ExpectGet(args.key).SetErr(redis.ErrClosed)
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args)
			}

			got, err := rc.Get(context.TODO(), tt.args.key)

			assert.Equal(t, tt.wantErr, (err != nil))
			assert.Equal(t, tt.want, got)
		})
	}
}
