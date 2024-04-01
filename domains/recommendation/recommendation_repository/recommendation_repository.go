package recommendation_repository

import (
	"glossika_be_interview/db_client"
	"glossika_be_interview/domains/recommendation/recommendation_entities"
)

func FindAll() ([]recommendation_entities.Recommendation, error) {
	var recommendations []recommendation_entities.Recommendation
	rows, err := db_client.DB.Query("SELECT productId FROM recommendations")
    if err != nil {
		return recommendations, err
    }
    defer rows.Close()

    
    for rows.Next() {
		var recommendation recommendation_entities.Recommendation
        if err := rows.Scan(&recommendation.ProductId); err != nil {
            return recommendations, err
        }
		recommendations = append(recommendations, recommendation)
    }

    if err := rows.Err(); err != nil {
        return recommendations, err
    }
	return recommendations, nil
}