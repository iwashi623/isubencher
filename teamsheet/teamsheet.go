package teamsheet

type TeamSheet interface {
	GetTeamNameByIP(ip string) (string, error)
}
