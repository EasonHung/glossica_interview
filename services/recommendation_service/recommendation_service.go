package recommendation_service

import (
	"context"
	"encoding/json"
	"glossika_be_interview/db_client"
	"glossika_be_interview/domains/recommendation/recommendation_entities"
	"glossika_be_interview/domains/recommendation/recommendation_repository"
	"time"
)

func GetAll() ([]recommendation_entities.Recommendation, error) {
	var recommendations []recommendation_entities.Recommendation

	recommendationListStr, err := db_client.Rdb.Get(context.TODO(), "recommendation-list").Result()
	if err == nil {
		err = json.Unmarshal([]byte(recommendationListStr), &recommendations)
		if err != nil {
			return nil, err
		}
		return recommendations, nil
	}

	recommendations, err = recommendation_repository.FindAll()
	if err != nil {
		return nil, err
	}

	recommendationListByte, err := json.Marshal(recommendations)
	err = db_client.Rdb.Set(context.TODO(), "recommendation-list", string(recommendationListByte), time.Minute*time.Duration(10)).Err()
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}
