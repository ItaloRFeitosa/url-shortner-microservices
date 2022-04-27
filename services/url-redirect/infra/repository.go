package infra

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-redirect/domain"
)

type DynamoDBRepository struct {
	d *dynamodb.Client
}

func NewDynamoDBRepository() *DynamoDBRepository {
	d := NewDynamoDB()
	return &DynamoDBRepository{d}
}

func (d *DynamoDBRepository) Insert(ctx context.Context, url *domain.ShortURL) error {
	PK := fmt.Sprint("SHORT_URL#", url.ID)
	_, err := d.d.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("url-redirect"),
		Item: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: PK,
			},
			"SK": &types.AttributeValueMemberS{
				Value: PK,
			},
			"ID": &types.AttributeValueMemberN{
				Value: strconv.Itoa(url.ID),
			},
			"RedirectTo": &types.AttributeValueMemberS{
				Value: url.RedirectTo,
			},
			"CreatedAt": &types.AttributeValueMemberS{
				Value: url.CreatedAt.Format(time.RFC3339),
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (d *DynamoDBRepository) FindRedirectURL(ctx context.Context, ID int) (string, error) {

	PK := fmt.Sprint("SHORT_URL#", ID)

	result, err := d.d.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("url-redirect"),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: PK,
			},
			"SK": &types.AttributeValueMemberS{
				Value: PK,
			},
		},
		ProjectionExpression: aws.String("RedirectTo"),
	})

	if err != nil {
		return "", err
	}

	attr, ok := result.Item["RedirectTo"]

	if !ok {
		return "", errors.New("redirect url not found")
	}
	url, ok := attr.(*types.AttributeValueMemberS)

	if !ok {
		return "", errors.New("redirect url not found")
	}
	return url.Value, nil
}
