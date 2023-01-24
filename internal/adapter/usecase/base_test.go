package usecase

import (
	"os"
	"testing"

	"jang-article/internal/port/mock"

	"github.com/golang/mock/gomock"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type MockHelper struct {
	ArticleUsecase *Usecases
	mockCtrl       *gomock.Controller
	mockPostgres   *mock.MockDatabaseRepository
	mockRedis      *mock.MockCacheRepository
	mockRedisearch *mock.MockSearchRepository
	mockValidation *mock.MockValidation
}

func usecaseTestHelper(t *testing.T) MockHelper {
	t.Helper()
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPostgres := mock.NewMockDatabaseRepository(mockCtrl)
	mockRedis := mock.NewMockCacheRepository(mockCtrl)
	mockRedisearch := mock.NewMockSearchRepository(mockCtrl)
	mockValidation := mock.NewMockValidation(mockCtrl)
	usecase := NewUsecases(mockValidation, mockPostgres, mockRedis, mockRedisearch)

	MockHelper := MockHelper{
		mockCtrl:       mockCtrl,
		mockPostgres:   mockPostgres,
		mockRedis:      mockRedis,
		mockRedisearch: mockRedisearch,
		mockValidation: mockValidation,
		ArticleUsecase: usecase,
	}

	return MockHelper
}
