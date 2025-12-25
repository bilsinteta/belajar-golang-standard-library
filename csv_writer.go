package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	writer.Write([]string{"eko", "kurniawan", "khannedy"})
	writer.Write([]string{"budi", "pratama", "diah"})
	writer.Write([]string{"joko", "morro", "diah"})

	writer.Flush()
}
