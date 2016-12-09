package model

import (
	"fmt"
	"strings"
)

type CheckedItems []CheckedItem

const itemsReport = "%d items have metrics, which out of bounds"

func (items CheckedItems) Report() string {
	var reports []string

	for _, item := range items {
		reports = append(reports, item.Report())
	}

	report := fmt.Sprintf(itemsReport, len(items))

	return strings.Join(append([]string{report}, reports), "\n  ")

}

type CheckedItem struct {
	ID      string
	Metrics []CheckedMetric
}

const itemReport = "Item %s have %d are out of bounds metrics: "

func (c CheckedItem) Report() string {
	var reports []string

	for _, metric := range c.Metrics {
		reports = append(reports, metric.Report())
	}

	report := fmt.Sprintf(itemReport, c.ID, len(c.Metrics))

	return strings.Join(append([]string{report}, reports), "\n  ")
}

type CheckedMetric struct {
	Name           string
	Value          uint
	ComparingValue uint
	Upper          bool
}

const metricReport = "- Metric %s is out of bounds. Value %d %s than %s %d"

func (c CheckedMetric) Report() string {
	var sign string
	var extr string

	if c.Upper {
		sign = "greater"
		extr = "max"
	} else {
		sign = "less"
		extr = "min"
	}

	return fmt.Sprintf(metricReport, c.Name, c.Value, sign, extr, c.ComparingValue)
}
