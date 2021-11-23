package counter

import (
	"app/internal/app/services/post"
	mock_post "app/test/gomock/mocks/postmock"
	// "errors"
	// "fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostCount(t *testing.T) {
	req := require.New(t)
	fakeslice := []post.Post{{ID: 3, Title: "lol", Body: "kek", UserID: 2}, {ID: 4, Title: "lol", Body: "kek", UserID: 3}}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockClient := mock_post.NewMockClient(mockCtrl)

	MockClient.EXPECT().GetList().Return(fakeslice, nil).Times(1)
	// MockClient.EXPECT().GetList().Return(fakeslice, errors.New("alarm")).Times(1)

	value, err := PostCount(MockClient)
	req.NoError(err)
	req.Equal(value, 2)
}
