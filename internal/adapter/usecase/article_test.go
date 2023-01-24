package usecase

import (
	"context"
	"errors"
	"fmt"
	"jang-article/internal/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	errMock = errors.New("error mock test")
)

func Test_SaveArticle(t *testing.T) {
	mockHelper := usecaseTestHelper(t)

	type args struct {
		ctx context.Context
		in  model.Article
	}

	type MockData struct {
		wantOut model.Article
		wantErr bool
	}

	tests := []struct {
		name     string
		args     args
		MockData MockData
		doMock   func(args args, mockData MockData)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				in: model.Article{
					Author: "author_jang",
					Title:  "title_test",
					Body:   "body_test",
				},
			},
			MockData: MockData{
				wantOut: model.Article{
					Author: "author_jang",
					Title:  "title_test",
					Body:   "body_test",
				},
				wantErr: false,
			},
			doMock: func(args args, mockData MockData) {
				mockHelper.mockValidation.EXPECT().ValidateRequest(args.ctx, args.in).Return(nil)
				mockHelper.mockPostgres.EXPECT().Save(args.in).Return(mockData.wantOut, nil)
				mockHelper.mockRedisearch.EXPECT().UpdateIndex(gomock.Any()).Return(nil)
			},
		},
		{
			name: "error when request required param is not valid",
			args: args{
				ctx: context.Background(),
				in: model.Article{
					Author: "",
					Title:  "",
					Body:   "body_test",
				},
			},
			MockData: MockData{
				wantErr: true,
			},
			doMock: func(args args, mockData MockData) {
				mockHelper.mockValidation.EXPECT().ValidateRequest(args.ctx, args.in).Return(errMock)
			},
		},
		{
			name: "error when trying to save article into database",
			args: args{
				ctx: context.Background(),
				in: model.Article{
					Author: "author_jang",
					Title:  "title_test",
					Body:   "body_test",
				},
			},
			MockData: MockData{
				wantOut: model.Article{
					Author: "",
					Title:  "",
					Body:   "",
				},
				wantErr: true,
			},
			doMock: func(args args, mockData MockData) {
				mockHelper.mockValidation.EXPECT().ValidateRequest(args.ctx, args.in).Return(nil)
				mockHelper.mockPostgres.EXPECT().Save(gomock.Any()).Return(mockData.wantOut, errMock)
			},
		},
		{
			name: "error when searching article in redisearch Index",
			args: args{
				ctx: context.Background(),
				in: model.Article{
					Author: "author_jang",
					Title:  "title_test",
					Body:   "body_test",
				},
			},
			MockData: MockData{
				wantOut: model.Article{
					Author: "",
					Title:  "",
					Body:   "",
				},
				wantErr: true,
			},
			doMock: func(args args, mockData MockData) {
				mockHelper.mockValidation.EXPECT().ValidateRequest(args.ctx, args.in).Return(nil)
				mockHelper.mockPostgres.EXPECT().Save(args.in).Return(mockData.wantOut, nil)
				mockHelper.mockRedisearch.EXPECT().UpdateIndex(gomock.Any()).Return(errMock)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args, tt.MockData)
			}
			gotout, err := mockHelper.ArticleUsecase.Article.SaveArticle(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.MockData.wantOut, gotout)
			assert.Equal(t, tt.MockData.wantErr, (err != nil))
		})
	}
}

func Test_GetArticle(t *testing.T) {
	mockHelper := usecaseTestHelper(t)

	type args struct {
		ctx context.Context
		in  string
	}

	type MockData struct {
		wantOut []model.Article
		wantErr bool
	}

	tests := []struct {
		name     string
		args     args
		MockData MockData
		doMock   func(args args, mockData MockData)
	}{
		{
			name: "success when there is cached items",
			args: args{
				ctx: context.Background(),
				in:  "article",
			},
			MockData: MockData{
				wantOut: []model.Article(nil),
				wantErr: true,
			},
			doMock: func(args args, mockData MockData) {
				mockHelper.mockRedis.EXPECT().Get(args.ctx, fmt.Sprintf("get_article:%v", args.in)).Return("", nil)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args, tt.MockData)
			}
			gotout, err := mockHelper.ArticleUsecase.Article.GetArticle(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.MockData.wantOut, gotout)
			assert.Equal(t, tt.MockData.wantErr, (err != nil))
		})
	}
}
