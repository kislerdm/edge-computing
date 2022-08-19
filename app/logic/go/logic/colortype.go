package logic

import (
	"errors"
	"math"
)

type node struct {
	ID        int     `json:"nodeid"`
	Depth     int     `json:"depth,omitempty"`
	Feature   string  `json:"split,omitempty"`
	Threshold float64 `json:"split_condition,omitempty"`
	Yes       int     `json:"yes,omitempty"`
	No        int     `json:"no,omitempty"`
	Missing   int     `json:"missing,omitempty"`
	Leaf      float64 `json:"leaf,omitempty"`
	Children  []*node `json:"children,omitempty"`
}

func (t *node) getNodeByID(id int) *node {
	if t.ID == id {
		return t
	}
	for _, nodeChild := range t.Children {
		if nodeChild.ID == id {
			return nodeChild
		}
		if n := nodeChild.getNodeByID(id); n != nil {
			return n
		}
	}
	return nil
}

type model []node

func (m model) predict(v map[string]float64) (float64, error) {
	var o float64
	for _, tree := range m {
		idx := 0
		for {
			n := tree.getNodeByID(idx)
			if n == nil {
				return 0, errors.New("node not found")
			}

			if len(n.Children) == 0 {
				o += n.Leaf
				break
			}

			f, ok := v[n.Feature]
			if !ok {
				idx = n.Missing
				continue
			}

			idx = n.Yes
			if f >= n.Threshold {
				idx = n.No
			}
		}
	}

	return o, nil
}

func Type(r, g, b float64) (bool, error) {
	if m == nil {
		return false, errors.New("wrong model generated")
	}

	v, err := m.predict(
		map[string]float64{
			"r": r,
			"g": g,
			"b": b,
		},
	)
	if err != nil {
		return false, err
	}

	return math.Exp(-v) <= 1., nil
}
