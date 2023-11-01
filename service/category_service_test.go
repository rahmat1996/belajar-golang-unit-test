package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil) // program mock
	category, err := categoryService.Get("1")               // execute category service get
	assert.NotNil(t, err)                                   // return error must have
	assert.Nil(t, category)                                 // return category must nil
}

func TestCategoryService_GetSuccess(t *testing.T) {
	// initial for return value mock
	data := entity.Category{
		Id:   "2",
		Name: "Computer",
	}
	categoryRepository.Mock.On("FindById", "2").Return(data) // program mock return value category
	category, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, data.Id, category.Id)
	assert.Equal(t, data.Name, category.Name)
}
