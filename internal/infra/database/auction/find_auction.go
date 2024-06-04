package auction

import (
	"context"
	"time"

	"github.com/natansa/go-leilao/configuration/logger"
	"github.com/natansa/go-leilao/internal/entity/auction_entity"
	"github.com/natansa/go-leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ar *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{"_id": id}
	var auctionEntityMongo AuctionEntityMongo
	if err := ar.Collection.FindOne(ctx, filter).Decode(&auctionEntityMongo); err != nil {
		logger.Error("Error finding auction in database", err)
		return nil, internal_error.NewInternalServerError("Error finding auction in database")
	}

	return &auction_entity.Auction{
		Id:          auctionEntityMongo.Id,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition:   auctionEntityMongo.Condition,
		Status:      auctionEntityMongo.Status,
		Timestamp:   time.Unix(auctionEntityMongo.Timestamp, 0),
	}, nil
}

func (ar *AuctionRepository) FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, category string, productName string) ([]auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["productName"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}

	cursor, err := ar.Collection.Find(ctx, filter)

	if err != nil {
		logger.Error("Error finding auctions in database", err)
		return nil, internal_error.NewInternalServerError("Error finding auctions in database")
	}

	defer cursor.Close(ctx)

	var auctionEntityMongo []AuctionEntityMongo
	if err := cursor.All(ctx, auctionEntityMongo); err != nil {
		logger.Error("Error finding auctions in database", err)
		return nil, internal_error.NewInternalServerError("Error finding auctions in database")
	}

	var auctionEntity []auction_entity.Auction
	for _, auctionMongo := range auctionEntityMongo {
		auctionEntity = append(auctionEntity, auction_entity.Auction{
			Id:          auctionMongo.Id,
			ProductName: auctionMongo.ProductName,
			Category:    auctionMongo.Category,
			Description: auctionMongo.Description,
			Condition:   auctionMongo.Condition,
			Status:      auctionMongo.Status,
			Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
		})
	}

	return auctionEntity, nil
}
