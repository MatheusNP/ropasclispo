package service

import (
	"context"
	"math/rand"

	"github.com/MatheusNP/ropasclispo/internal/constant"
	"github.com/MatheusNP/ropasclispo/internal/domain"
)

type Play interface {
	Result(ctx context.Context, request domain.FightRequest) (domain.FightResponse, error)
}

type playService struct{}

func NewPlayService() Play {
	return &playService{}
}

func (p playService) Result(ctx context.Context, request domain.FightRequest) (domain.FightResponse, error) {
	choices := constant.CHOICES()

	enemyChoice := domain.Choose(uint8(rand.Intn(len(choices))), choices)
	myChoice := domain.Choose(request.Value, choices)

	var point int8
	var result string

	for _, v := range enemyChoice.WinsFrom {
		if myChoice.Value == v.Value {
			point = -1
			result = v.Message
			break
		}
	}

	if point != -1 {
		for _, v := range myChoice.WinsFrom {
			if enemyChoice.Value == v.Value {
				point = 1
				result = v.Message
				break
			}
		}
	}

	return domain.FightResponse{
		Enemy:  enemyChoice,
		Point:  point,
		Result: result,
	}, nil
}
