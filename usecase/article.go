package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/snappy"
	"news/entity"
)

func CreateArticle(param entity.Article) error {
	db, err := entity.Init()
	if err != nil {
		return err
	}
	result := db.Create(&param)
	if result.Error != nil {
		return errors.New("Failed to insert DB")
	}
	return nil
}

func GetArticle(param entity.ArticleParam) ([]entity.Article, error) {
	var articles []entity.Article

	// get form redis
	rawKey, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("rediskey:%s", rawKey)
	redisResult, err := getRedisArticle(key)
	if err != nil {
		return nil, err
	}
	if len(redisResult) != 0 {
		return redisResult, nil
	}

	db, err := entity.Init()
	if err != nil {
		return nil, err
	}
	db.Model(&articles)
	if param.Author != "" {
		db = db.Where("author like ?", "%"+param.Author+"%")
	}
	if param.Query != "" {
		db = db.Where("body like ? or title like ?", "%"+param.Query+"%", "%"+param.Query+"%")
	}

	err = db.Find(&articles).Error
	if err != nil {
		return nil, err
	}

	err = setRedisArticle(param, articles)
	if err != nil {
		return nil, nil
	}

	return articles, nil
}

func getRedisArticle(key string) ([]entity.Article, error) {
	var article []entity.Article
	var decRedis []byte
	result, err := entity.GetKey(key)
	if err != nil {
		return nil, nil
	}

	decRedis, err = snappy.Decode(decRedis, result)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(decRedis, &article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func setRedisArticle(param entity.ArticleParam, value []entity.Article) error {
	rawKey, err := json.Marshal(param)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("rediskey:%s", rawKey)

	rawJSON, err := json.Marshal(value)
	if err != nil {
		return err
	}
	var encJSON []byte
	encJSON = snappy.Encode(encJSON, rawJSON)

	err = entity.SetKey(key, encJSON)
	if err != nil {
		return err
	}
	return nil

}
