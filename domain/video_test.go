package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()
	require.Error(t, err)

}

func TestVideoIdIsNotAUuid(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "not-uuid"
	video.ResourceID = "a988f820-7b1a-4953-af37-9fbaf2d70b0b"
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	err := video.Validate()
	require.Error(t, err)

}

func TestVideoValidateIsOkay(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "a988f820-7b1a-4953-af37-9fbaf2d70b0b"
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	err := video.Validate()
	require.Nil(t, err)

}
