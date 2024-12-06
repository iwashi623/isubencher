package teamboard

import (
	"context"
)

type TeamBoard interface {
	GetTeamNameByIP(ctx context.Context, ip string) (string, error)
}

type NilTeamBoard struct {
}

func NewNilTeamBoard() *NilTeamBoard {
	return &NilTeamBoard{}
}

func (n *NilTeamBoard) GetTeamNameByIP(ctx context.Context, ip string) (string, error) {
	return "", nil
}

type TeamBoardCreateFunc func() (TeamBoard, error)
