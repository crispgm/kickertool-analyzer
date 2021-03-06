package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/crispgm/kickertool-analyzer/elo"
	"github.com/pterm/pterm"
)

func main() {
	if len(os.Args) < 5 {
		pterm.Error.Println("Invalid params")
	}

	t1p1Score := convertToFloat(os.Args[1])
	t1p2Score := convertToFloat(os.Args[2])
	t2p1Score := convertToFloat(os.Args[3])
	t2p2Score := convertToFloat(os.Args[4])
	k := 10.0
	if len(os.Args) > 5 {
		k = convertToFloat(os.Args[5])
	}
	rhw := elo.Rate{
		T1P1Score: t1p1Score,
		T1P2Score: t1p2Score,
		T2P1Score: t2p1Score,
		T2P2Score: t2p2Score,
		HostWin:   true,
		K:         k,
	}
	rhw.CalcEloRating()
	raw := elo.Rate{
		T1P1Score: t1p1Score,
		T1P2Score: t1p2Score,
		T2P1Score: t2p1Score,
		T2P2Score: t2p2Score,
		HostWin:   false,
		K:         k,
	}
	raw.CalcEloRating()

	fmt.Printf("Estimated Elo rating (k=%d):\n", int(k))

	fmt.Println("- If host won:")
	hostWonData := [][]string{
		{"Team", "Player", "Cur Score", "Expect", "New Score", "Rating Change"},
		{"A", "1", fmt.Sprintf("%.f", t1p1Score), fmt.Sprintf("%.2f%%", rhw.T1P1Exp*100), fmt.Sprintf("%.f", rhw.T1P1Score), pterm.Green("+", rhw.T1P1Score-t1p1Score)},
		{"A", "2", fmt.Sprintf("%.f", t1p2Score), fmt.Sprintf("%.2f%%", rhw.T1P2Exp*100), fmt.Sprintf("%.f", rhw.T1P2Score), pterm.Green("+", rhw.T1P2Score-t1p2Score)},
		{"B", "1", fmt.Sprintf("%.f", t2p1Score), fmt.Sprintf("%.2f%%", rhw.T2P1Exp*100), fmt.Sprintf("%.f", rhw.T2P1Score), pterm.Red(rhw.T2P1Score - t2p1Score)},
		{"B", "2", fmt.Sprintf("%.f", t2p2Score), fmt.Sprintf("%.2f%%", rhw.T2P2Exp*100), fmt.Sprintf("%.f", rhw.T2P2Score), pterm.Red(rhw.T2P2Score - t2p2Score)},
	}
	pterm.DefaultTable.WithHasHeader().WithData(hostWonData).WithBoxed(true).Render()

	fmt.Println("- If away won:")
	awayWonData := [][]string{
		{"Team", "Player", "Cur Score", "Expect", "New Score", "Rating Change"},
		{"A", "1", fmt.Sprintf("%.f", t1p1Score), fmt.Sprintf("%.2f%%", raw.T1P1Exp*100), fmt.Sprintf("%.f", raw.T1P1Score), pterm.Red(raw.T1P1Score - t1p1Score)},
		{"A", "2", fmt.Sprintf("%.f", t1p2Score), fmt.Sprintf("%.2f%%", raw.T1P2Exp*100), fmt.Sprintf("%.f", raw.T1P2Score), pterm.Red(raw.T1P2Score - t1p2Score)},
		{"B", "1", fmt.Sprintf("%.f", t2p1Score), fmt.Sprintf("%.2f%%", raw.T2P1Exp*100), fmt.Sprintf("%.f", raw.T2P1Score), pterm.Green("+", raw.T2P1Score-t2p1Score)},
		{"B", "2", fmt.Sprintf("%.f", t2p2Score), fmt.Sprintf("%.2f%%", raw.T2P2Exp*100), fmt.Sprintf("%.f", raw.T2P2Score), pterm.Green("+", raw.T2P2Score-t2p2Score)},
	}
	pterm.DefaultTable.WithHasHeader().WithData(awayWonData).WithBoxed(true).Render()
}

func convertToFloat(in string) float64 {
	out, _ := strconv.Atoi(in)
	return float64(out)
}
