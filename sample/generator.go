package sample

import "klever-api/pb"

func NewExchange() *pb.ExchangeRequest {
	exchange := &pb.ExchangeRequest{
		CoinName: randomCoinName(),
		Value: randomFloat32(),
	}

	return exchange
}

func NewCoinExchange() *pb.CoinExchangeRequest {
	coinExchange := &pb.CoinExchangeRequest{
		CoinName: randomCoinName(),
	}

	return coinExchange
}
