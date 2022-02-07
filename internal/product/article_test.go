package product

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/unicod3/horreum/pkg/dbclient"
	"github.com/unicod3/horreum/pkg/dbclient/mocks"
	"testing"
)

func TestArticleServiceImplementsArticleRepositoryInterface(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*ArticleRepository)(nil), new(ArticleService))
}

func TestArticleService_GetAll(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	articleService := &ArticleService{
		dataTable: &dataTable,
	}

	articles := []Article{
		Article{ID: 1, Name: "test"},
	}

	var w []Article
	dataTable.On("FindAll", &w).Run(func(args mock.Arguments) {
		w = articles
	}).Return(nil).Once()
	_, err := articleService.GetAll()
	assert.Nil(err)
	assert.Equal(articles, w)
}

func TestArticleService_GetById(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	articleService := &ArticleService{
		dataTable: &dataTable,
	}

	article := Article{ID: 1, Name: "test"}

	var w Article
	dataTable.On("FindOne", dbclient.Condition{"id": article.ID}, &w).Run(func(args mock.Arguments) {
		w = article
	}).Return(nil).Once()
	_, err := articleService.GetById(article.ID)
	assert.Nil(err)
	assert.Equal(article, w)
}

func TestArticleService_Create(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	articleService := &ArticleService{
		dataTable: &dataTable,
	}

	article := Article{ID: 1, Name: "test"}

	var w Article
	dataTable.On("InsertReturning", &article).Run(func(args mock.Arguments) {
		w = article
	}).Return(nil).Once()
	err := articleService.Create(&article)
	assert.Nil(err)
	assert.Equal(article, w)
}

func TestArticleService_Update(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	articleService := &ArticleService{
		dataTable: &dataTable,
	}

	article := Article{ID: 1, Name: "test"}

	var w Article
	dataTable.On("UpdateReturning", &article).Run(func(args mock.Arguments) {
		w = article
	}).Return(nil).Once()
	err := articleService.Update(&article)
	assert.Nil(err)
	assert.Equal(article, w)
}

func TestArticleService_Delete(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	articleService := &ArticleService{
		dataTable: &dataTable,
	}

	article := Article{ID: 1, Name: "test"}
	dataTable.On("Delete", dbclient.Condition{"id": article.ID}).Return(nil).Once()
	err := articleService.Delete(&article)
	assert.Nil(err)
}
