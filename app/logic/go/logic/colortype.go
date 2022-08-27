package logic

import (
	"errors"
	"math"
)

type node struct {
	ID        uint8   `json:"nodeid"`
	Depth     uint8   `json:"depth,omitempty"`
	Feature   string  `json:"split,omitempty"`
	Threshold uint8   `json:"split_condition,omitempty"`
	Yes       uint8   `json:"yes,omitempty"`
	No        uint8   `json:"no,omitempty"`
	Missing   uint8   `json:"missing,omitempty"`
	Leaf      float64 `json:"leaf,omitempty"`
	Children  []*node `json:"children,omitempty"`
}

func (t *node) getNodeByID(id uint8) *node {
	if t.ID == id {
		return t
	}
	for _, nodeChild := range t.Children {
		if n := nodeChild.getNodeByID(id); n != nil {
			return n
		}
	}
	return nil
}

type model []node

func (m model) predict(v map[string]uint8) (float64, error) {
	var o float64
	for _, tree := range m {
		var idx uint8 = 0
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

func Type(r, g, b uint8) (bool, error) {
	if colorTypeModel == nil {
		return false, errors.New("wrong model generated")
	}

	v, err := colorTypeModel.predict(
		map[string]uint8{
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
