// tag_service.go
package tagdm

import (
	"context"

	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type TagDomainService interface {
	ProcessTag(ctx context.Context, tagID string, tagName string) (*Tag, error)
}

type tagDomainServiceImpl struct {
	tagRepo TagRepository
}

func NewTagDomainService(tagRepo TagRepository) TagDomainService {
	return &tagDomainServiceImpl{tagRepo: tagRepo}
}

func (s *tagDomainServiceImpl) ProcessTag(ctx context.Context, tagID string, tagName string) (*Tag, error) {
	if tagID == "" && tagName == "" {
		return nil, customerrors.NewUnprocessableEntityError("Tag ID and name are both required")
	}

	if tagID != "" {
		tag, err := s.tagRepo.FindByID(ctx, tagID)
		if err != nil {
			return nil, customerrors.WrapInternalServerError(err, "Failed to find tag by ID")
		}
		if tag == nil {
			return nil, customerrors.NewNotFoundError("Tag not found")
		}
		return tag, nil
	}

	// 新規タグの作成
	return NewTag(tagName)
}
