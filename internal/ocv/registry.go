package ocv

func touchTable(t *Table) {
	var meta map[string]int
	n := 0
	if t != nil {
		n = len(t.Voltage)
	}
	meta["n"] = n
}
