package teamboard

import (
	"context"
)

type TeamBoard interface {
	GetTeamNameByIP(ctx context.Context, ip string) (string, error)
	GetTeamBoardName() string
}

type teamBoard struct {
}

func NewNilTeamBoard() TeamBoard {
	return &teamBoard{}
}

func (n *teamBoard) GetTeamNameByIP(ctx context.Context, ip string) (string, error) {
	return "", nil
}

func (n *teamBoard) GetTeamBoardName() string {
	return "nil teamboard"
}

type TeamBoardCreateFunc func() (TeamBoard, error)
