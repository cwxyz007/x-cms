package service

import "github.com/cwxyz007/x-cms/model"

type PostCategoryService struct {
	BasicService
}

func (s *PostCategoryService) GetCategoriesByPostIds(postIds []string) ([]model.PostCategory, error) {
	var postCategories []model.PostCategory

	err := s.db.Where("post_id in ?", postIds).Find(&postCategories).Error

	return postCategories, err
}
