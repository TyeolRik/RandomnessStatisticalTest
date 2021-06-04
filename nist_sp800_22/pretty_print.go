package nist_sp800_22

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const NONE uint64 = ^uint64(0)

var t table.Writer
var testNumber uint16 = 1

func PrettyPrint_Add(testName string, P_value float64, isRandom bool) {
	var _isRandom string
	if isRandom {
		_isRandom = "Random"
	} else {
		_isRandom = "Non-Random"
	}
	t.AppendRows([]table.Row{
		{testNumber, testName, "-", _isRandom, fmt.Sprintf("%.11f", P_value)},
	})

	t.AppendSeparator()
	testNumber++
}

func PrettyPrint_Add_Array(testName string, P_value []float64, isRandom []bool) {
	var _isRandom string
	var countRandom, countNonRandom uint64
	for _, value := range isRandom {
		if value {
			countRandom++
		} else {
			countNonRandom++
		}
	}

	if countRandom > countNonRandom {
		_isRandom = "Random"
	} else {
		_isRandom = "Non-Random"
	}

	t.AppendRows([]table.Row{
		{testNumber, testName, fmt.Sprintf("%d / %d PASS", countRandom, len(isRandom)), _isRandom, "-"},
	})

	if len(P_value) > 5 {
		t.AppendRows([]table.Row{
			{"-", "", 1, _isRandom, fmt.Sprintf("%.11f", P_value[0])},
		})
		t.AppendRows([]table.Row{
			{"-", "", 2, _isRandom, fmt.Sprintf("%.11f", P_value[1])},
		})
		t.AppendRows([]table.Row{
			{"-", "", "..."},
		})
		t.AppendRows([]table.Row{
			{"-", "", len(P_value) - 1, _isRandom, fmt.Sprintf("%.11f", P_value[len(P_value)-2])},
		})
		t.AppendRows([]table.Row{
			{"-", "", len(P_value), _isRandom, fmt.Sprintf("%.11f", P_value[len(P_value)-1])},
		})
	} else {
		for i := range P_value {
			t.AppendRows([]table.Row{
				{"-", testName, i + 1, _isRandom, fmt.Sprintf("%.11f", P_value[i])},
			})
		}
	}

	t.AppendSeparator()
	testNumber++
}

func PrettyPrint_Init() {
	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Test Name", "Sub test Count", "Conclusion", "P-value"})
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 2, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 3, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 4, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
	})
}

func PrettyPrint_Render() {
	t.Render()
}
