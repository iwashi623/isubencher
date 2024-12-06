package teamsheet

import "context"

type TeamSheet interface {
	GetTeamNameByIP(ctx context.Context, ip string) (string, error)
}

type teamSheet struct {
	sheet TeamSheet
}

func NewTeamSheet(
	s TeamSheet,
) *teamSheet {
	return &teamSheet{
		sheet: s,
	}
}
