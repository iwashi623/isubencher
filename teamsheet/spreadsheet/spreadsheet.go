package spreadsheet

type Spreadsheet struct {
}

func NewSpreadsheet() *Spreadsheet {
	return &Spreadsheet{}
}

func (s *Spreadsheet) GetTeamNameByIP(ip string) string {
	return "team1"
}
