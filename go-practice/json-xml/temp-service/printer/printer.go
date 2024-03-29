package printer

import (
	"fmt"
	"os"
	"temp-service/models"
	"text/tabwriter"
)

type Printer struct {
	w *tabwriter.Writer
}

func New() *Printer {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
	return &Printer{
		w: w,
	}
}

// CityHeader prints out the table header for the city table
func (p *Printer) CityHeader() {
	fmt.Fprintln(p.w, "Name\tTempC\tTempF\tBeach ready?\tSki ready?")
}

// CityDetails prints out the given city
func (p *Printer) CityDetails(c models.CityTemp) {
	fmt.Fprintf(p.w, "%v\t%v\t%v\t%v\t%v\n", c.Name(), c.TempC(), c.TempF(), c.BeachVacationReady(), c.SkiVacationReady())
}

// Cleanup closes up the priter instance
func (p *Printer) Cleanup() {
	p.w.Flush()
}
