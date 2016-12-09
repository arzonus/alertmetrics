package checker

import (
	"errors"
	"fmt"
	m "github.com/arzonus/alertmetrics/pkg/model"
)

type Checker struct {
	LowerBounds map[string]uint
	UpperBounds map[string]uint
}

func New(LowerBounds map[string]uint, UpperBounds map[string]uint) *Checker {
	return &Checker{
		LowerBounds: LowerBounds,
		UpperBounds: UpperBounds,
	}
}

func (c Checker) Validate() error {
	if len(c.LowerBounds) != len(c.UpperBounds) {
		return errors.New("Sizes of bounds doesn't match ")
	}

	for k, ub := range c.UpperBounds {

		lb, ok := c.LowerBounds[k]
		if !ok {
			return errors.New(fmt.Sprint("Does not exist key in lower bounds: ", k))
		}

		if lb > ub {
			return errors.New(fmt.Sprint("Lower bounds of ", k, " ", lb, " greater than ", ub))
		}
	}

	return nil
}

func (c Checker) CheckItems(items m.Items) *m.CheckedItems {
	checkedItems := new(m.CheckedItems)

	for i := range items {
		checkedItem := c.CheckItem(items[i])
		if checkedItem != nil {
			*checkedItems = append(*checkedItems, *(c.CheckItem(items[i])))
		}
	}

	if len(*checkedItems) == 0 {
		return nil
	}

	return checkedItems
}

func (c Checker) CheckItem(item m.Item) *m.CheckedItem {
	var metrics []m.CheckedMetric

	for name, value := range item.Metrics {
		checkedMetric := c.CheckMetric(name, value)
		if checkedMetric != nil {
			metrics = append(metrics, *checkedMetric)
		}
	}

	if len(metrics) == 0 {
		return nil
	}

	checkedItem := new(m.CheckedItem)
	checkedItem.ID = item.ID
	checkedItem.Metrics = metrics

	return checkedItem
}

func (c Checker) CheckMetric(name string, value uint) *m.CheckedMetric {

	ub, ok := c.UpperBounds[name]
	if !ok {
		return nil
	}

	if ub < value {
		return &m.CheckedMetric{
			Name:           name,
			Value:          value,
			ComparingValue: ub,
			Upper:          true,
		}
	}

	lb, ok := c.LowerBounds[name]
	if !ok {
		return nil
	}

	if lb > value {
		return &m.CheckedMetric{
			Name:           name,
			Value:          value,
			ComparingValue: lb,
			Upper:          false,
		}
	}

	return nil
}
