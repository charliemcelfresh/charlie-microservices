package twirp_server

import (
	"context"
	"github.com/charliemcelfresh/charlie-microservices/internal/config"
	pb "github.com/charliemcelfresh/charlie-microservices/rpc/charlie-microservices"
	"log/slog"
	"time"

	"github.com/twitchtv/twirp"
)

type Repository interface {
	GetTransactions(ctx context.Context, userID, cardNetwork string) ([]Transaction, error)
}

type provider struct {
	Repository Repository
	Logger     *slog.Logger
}

func NewProvider() provider {
	return provider{
		NewRepository(config.DB()),
		config.GetLogger(),
	}
}

func (p provider) GetAmexTransactions(ctx context.Context, req *pb.None) (*pb.Transactions, error) {
	time.Sleep(time.Second)
	return p.getTransactions(ctx, "Amex")
}

func (p provider) GetMasterCardTransactions(ctx context.Context, req *pb.None) (*pb.Transactions, error) {
	time.Sleep(time.Second)
	return p.getTransactions(ctx, "MasterCard")
}

func (p provider) GetVisaTransactions(ctx context.Context, req *pb.None) (*pb.Transactions, error) {
	time.Sleep(time.Second)
	return p.getTransactions(ctx, "Visa")
}

func (p provider) getTransactions(ctx context.Context, cardNetwork string) (*pb.Transactions, error) {
	userID := getUserIdFromContext(ctx)
	transactions, err := p.Repository.GetTransactions(ctx, userID, cardNetwork)
	if err != nil {
		return &pb.Transactions{}, twirp.NewError(twirp.Unauthenticated, "Unauthorized")
	}
	transactionsToReturn := make([]*pb.Transaction, 0, len(transactions))
	for _, t := range transactions {
		pbT := &pb.Transaction{
			Id:            t.ID,
			UserAccountId: t.UserAccountID,
			CardNetwork:   t.CardNetwork,
			MerchantId:    t.MerchantID,
			MerchantName:  t.MerchantName,
			Total:         t.Total,
			CreatedAt:     t.CreatedAt,
		}
		transactionsToReturn = append(transactionsToReturn, pbT)
	}
	return &pb.Transactions{Transactions: transactionsToReturn}, err
}
