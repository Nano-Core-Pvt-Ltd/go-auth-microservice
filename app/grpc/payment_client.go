// Package grpcclient provides a gRPC client for the Payment service
// (sajilowallet). The Auth service uses it to create an initial wallet for
// newly registered users.
package grpcclient

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-wallet-microservice/go-wallet-microservice/proto/wallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// paymentAddr returns the Payment service gRPC address.
func paymentAddr() string {
	addr := os.Getenv("PAYMENT_GRPC_ADDR")
	if addr == "" {
		addr = "localhost:50051"
	}
	return addr
}

// dial opens a blocking connection to the Payment service.
func dial(ctx context.Context) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, paymentAddr(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
}

// CreateWallet creates an initial wallet for the given user in the Payment
// service. It is safe to call multiple times for the same user (the Payment
// service is idempotent).
func CreateWallet(ctx context.Context, userID, currency string) error {
	if userID == "" {
		return fmt.Errorf("user id is required to create wallet")
	}

	dialCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	conn, err := dial(dialCtx)
	if err != nil {
		return fmt.Errorf("failed to connect to payment service at %s: %w", paymentAddr(), err)
	}
	defer conn.Close()

	client := wallet.NewWalletServiceClient(conn)

	if currency == "" {
		currency = "USD"
	}

	resp, err := client.CreateWallet(ctx, &wallet.CreateWalletRequest{
		UserId:   userID,
		Currency: currency,
	})
	if err != nil {
		return fmt.Errorf("payment service CreateWallet failed: %w", err)
	}

	if resp.GetUserId() != userID {
		return fmt.Errorf("wallet user id mismatch: requested %q, got %q", userID, resp.GetUserId())
	}

	return nil
}
