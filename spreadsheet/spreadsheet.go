package spreadsheet

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/iwashi623/kinben/teamsheet"
)

type Spreadsheet struct {
}

var _ teamsheet.TeamSheet = (*Spreadsheet)(nil)

func NewSpreadsheet() *Spreadsheet {
	return &Spreadsheet{}
}

func (s *Spreadsheet) GetTeamNameByIP(ip string) (string, error) {
	unixTime := time.Now().Unix()
	filename := fmt.Sprintf("%s-%d.csv", ip, unixTime)
	err := s.donwloadCSV(filename)
	if err != nil {
		return "", fmt.Errorf("failed to download CSV: %w", err)
	}

	records, err := s.getContent(filename)
	if err != nil {
		return "", fmt.Errorf("failed to get content: %w", err)
	}

	ipToTeam := make(map[string]string, len(records)*3)
	for _, row := range records[1:] {
		teamName := row[0]
		for _, ip := range row[1:] {
			ipToTeam[ip] = teamName
		}
	}

	teamName, ok := ipToTeam[ip]
	if !ok {
		return "", fmt.Errorf("no team name found for IP: %s", ip)
	}

	return teamName, nil
}

func (s *Spreadsheet) donwloadCSV(filename string) error {
	url := "https://docs.google.com/spreadsheets/d/%s/export?format=csv"

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to get spreadsheet: %w", err)
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (s *Spreadsheet) getContent(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if err := os.Remove(filename); err != nil {
		return nil, fmt.Errorf("failed to remove file: %w", err)
	}

	return records, nil
}
