package logic

var m = model{
	{
		ID:        0,
		Depth:     0,
		Feature:   "b",
		Threshold: 162,
		Yes:       1,
		No:        2,
		Missing:   1,
		Children: []*node{
			{
				ID:        1,
				Depth:     1,
				Feature:   "r",
				Threshold: 86,
				Yes:       3,
				No:        4,
				Missing:   3,
				Children: []*node{
					{
						ID:        3,
						Depth:     2,
						Feature:   "g",
						Threshold: 98,
						Yes:       7,
						No:        8,
						Missing:   7,
						Children: []*node{
							{
								ID:   7,
								Leaf: -0,
							},
							{
								ID:   8,
								Leaf: -0.666666687,
							},
						},
					},
					{
						ID:        4,
						Depth:     2,
						Feature:   "b",
						Threshold: 112,
						Yes:       9,
						No:        10,
						Missing:   9,
						Children: []*node{
							{
								ID:   9,
								Leaf: 0.70588237,
							},
							{
								ID:   10,
								Leaf: 0.176470593,
							},
						},
					},
				},
			},
			{
				ID:        2,
				Depth:     1,
				Feature:   "r",
				Threshold: 247,
				Yes:       5,
				No:        6,
				Missing:   5,
				Children: []*node{
					{
						ID:   5,
						Leaf: -0.826086938,
					},
					{
						ID:        6,
						Depth:     2,
						Feature:   "g",
						Threshold: 221,
						Yes:       11,
						No:        12,
						Missing:   11,
						Children: []*node{
							{
								ID:   11,
								Leaf: -0.200000003,
							},
							{
								ID:   12,
								Leaf: 0.25,
							},
						},
					},
				},
			},
		},
	},
	{
		ID:        0,
		Depth:     0,
		Feature:   "b",
		Threshold: 152,
		Yes:       1,
		No:        2,
		Missing:   1,
		Children: []*node{
			{
				ID:        1,
				Depth:     1,
				Feature:   "r",
				Threshold: 92,
				Yes:       3,
				No:        4,
				Missing:   3,
				Children: []*node{
					{
						ID:   3,
						Leaf: -0.525906324,
					},
					{
						ID:        4,
						Depth:     2,
						Feature:   "b",
						Threshold: 75,
						Yes:       7,
						No:        8,
						Missing:   7,
						Children: []*node{
							{
								ID:   7,
								Leaf: 0.542077065,
							},
							{
								ID:   8,
								Leaf: 0.114652567,
							},
						},
					},
				},
			},
			{
				ID:        2,
				Depth:     1,
				Feature:   "g",
				Threshold: 236,
				Yes:       5,
				No:        6,
				Missing:   5,
				Children: []*node{
					{
						ID:   5,
						Leaf: -0.615649939,
					},
					{
						ID:   6,
						Leaf: -0.0373393893,
					},
				},
			},
		},
	},
	{
		ID:        0,
		Depth:     0,
		Feature:   "r",
		Threshold: 233,
		Yes:       1,
		No:        2,
		Missing:   1,
		Children: []*node{
			{
				ID:        1,
				Depth:     1,
				Feature:   "b",
				Threshold: 108,
				Yes:       3,
				No:        4,
				Missing:   3,
				Children: []*node{
					{
						ID:        3,
						Depth:     2,
						Feature:   "r",
						Threshold: 63,
						Yes:       7,
						No:        8,
						Missing:   7,
						Children: []*node{
							{
								ID:   7,
								Leaf: -0.383735001,
							},
							{
								ID:   8,
								Leaf: 0.147082433,
							},
						},
					},
					{
						ID:        4,
						Depth:     2,
						Feature:   "r",
						Threshold: 9,
						Yes:       9,
						No:        10,
						Missing:   9,
						Children: []*node{
							{
								ID:   9,
								Leaf: -0.115936428,
							},
							{
								ID:   10,
								Leaf: -0.505637467,
							},
						},
					},
				},
			},
			{
				ID:        2,
				Depth:     1,
				Feature:   "b",
				Threshold: 176,
				Yes:       5,
				No:        6,
				Missing:   5,
				Children: []*node{
					{
						ID:   5,
						Leaf: 0.597233415,
					},
					{
						ID:        6,
						Depth:     2,
						Feature:   "b",
						Threshold: 212,
						Yes:       11,
						No:        12,
						Missing:   11,
						Children: []*node{
							{
								ID:   11,
								Leaf: 0.0525054261,
							},
							{
								ID:   12,
								Leaf: -0.185408637,
							},
						},
					},
				},
			},
		},
	},
	{
		ID:        0,
		Depth:     0,
		Feature:   "r",
		Threshold: 194,
		Yes:       1,
		No:        2,
		Missing:   1,
		Children: []*node{
			{
				ID:        1,
				Depth:     1,
				Feature:   "b",
				Threshold: 74,
				Yes:       3,
				No:        4,
				Missing:   3,
				Children: []*node{
					{
						ID:        3,
						Depth:     2,
						Feature:   "r",
						Threshold: 62,
						Yes:       7,
						No:        8,
						Missing:   7,
						Children: []*node{
							{
								ID:   7,
								Leaf: -0.337070793,
							},
							{
								ID:   8,
								Leaf: 0.220472932,
							},
						},
					},
					{
						ID:        4,
						Depth:     2,
						Feature:   "g",
						Threshold: 149,
						Yes:       9,
						No:        10,
						Missing:   9,
						Children: []*node{
							{
								ID:   9,
								Leaf: -0.583193123,
							},
							{
								ID:   10,
								Leaf: -0.117968187,
							},
						},
					},
				},
			},
			{
				ID:        2,
				Depth:     1,
				Feature:   "g",
				Threshold: 208,
				Yes:       5,
				No:        6,
				Missing:   5,
				Children: []*node{
					{
						ID:   5,
						Leaf: 0.421763927,
					},
					{
						ID:        6,
						Depth:     2,
						Feature:   "r",
						Threshold: 246,
						Yes:       11,
						No:        12,
						Missing:   11,
						Children: []*node{
							{
								ID:   11,
								Leaf: -0.425008118,
							},
							{
								ID:   12,
								Leaf: 0.118131198,
							},
						},
					},
				},
			},
		},
	},
	{
		ID:        0,
		Depth:     0,
		Feature:   "b",
		Threshold: 176,
		Yes:       1,
		No:        2,
		Missing:   1,
		Children: []*node{
			{
				ID:        1,
				Depth:     1,
				Feature:   "r",
				Threshold: 151,
				Yes:       3,
				No:        4,
				Missing:   3,
				Children: []*node{
					{
						ID:        3,
						Depth:     2,
						Feature:   "b",
						Threshold: 74,
						Yes:       7,
						No:        8,
						Missing:   7,
						Children: []*node{
							{
								ID:   7,
								Leaf: 0.0376000479,
							},
							{
								ID:   8,
								Leaf: -0.327806205,
							},
						},
					},
					{
						ID:        4,
						Depth:     2,
						Feature:   "g",
						Threshold: 124,
						Yes:       9,
						No:        10,
						Missing:   9,
						Children: []*node{
							{
								ID:   9,
								Leaf: 0.116168551,
							},
							{
								ID:   10,
								Leaf: 0.445220917,
							},
						},
					},
				},
			},
			{
				ID:        2,
				Depth:     1,
				Feature:   "g",
				Threshold: 239,
				Yes:       5,
				No:        6,
				Missing:   5,
				Children: []*node{
					{
						ID:   5,
						Leaf: -0.479840338,
					},
					{
						ID:   6,
						Leaf: 0.132063642,
					},
				},
			},
		},
	},
	{
		ID:        0,
		Depth:     0,
		Feature:   "r",
		Threshold: 92,
		Yes:       1,
		No:        2,
		Missing:   1,
		Children: []*node{
			{
				ID:        1,
				Depth:     1,
				Feature:   "b",
				Threshold: 127,
				Yes:       3,
				No:        4,
				Missing:   3,
				Children: []*node{
					{
						ID:   3,
						Leaf: -0.367772937,
					},
					{
						ID:   4,
						Leaf: -0.0662978962,
					},
				},
			},
			{
				ID:        2,
				Depth:     1,
				Feature:   "b",
				Threshold: 214,
				Yes:       5,
				No:        6,
				Missing:   5,
				Children: []*node{
					{
						ID:        5,
						Depth:     2,
						Feature:   "g",
						Threshold: 119,
						Yes:       7,
						No:        8,
						Missing:   7,
						Children: []*node{
							{
								ID:   7,
								Leaf: -0.014199187,
							},
							{
								ID:   8,
								Leaf: 0.328373492,
							},
						},
					},
					{
						ID:   6,
						Leaf: -0.246607855,
					},
				},
			},
		},
	},
	{
		ID:        0,
		Depth:     0,
		Feature:   "g",
		Threshold: 207,
		Yes:       1,
		No:        2,
		Missing:   1,
		Children: []*node{
			{
				ID:        1,
				Depth:     1,
				Feature:   "r",
				Threshold: 155,
				Yes:       3,
				No:        4,
				Missing:   3,
				Children: []*node{
					{
						ID:        3,
						Depth:     2,
						Feature:   "g",
						Threshold: 93,
						Yes:       7,
						No:        8,
						Missing:   7,
						Children: []*node{
							{
								ID:   7,
								Leaf: 0.180075482,
							},
							{
								ID:   8,
								Leaf: -0.176351994,
							},
						},
					},
					{
						ID:        4,
						Depth:     2,
						Feature:   "g",
						Threshold: 119,
						Yes:       9,
						No:        10,
						Missing:   9,
						Children: []*node{
							{
								ID:   9,
								Leaf: 0.0896058604,
							},
							{
								ID:   10,
								Leaf: 0.41525653,
							},
						},
					},
				},
			},
			{
				ID:        2,
				Depth:     1,
				Feature:   "g",
				Threshold: 252,
				Yes:       5,
				No:        6,
				Missing:   5,
				Children: []*node{
					{
						ID:   5,
						Leaf: -0.321783125,
					},
					{
						ID:   6,
						Leaf: 0.0958191752,
					},
				},
			},
		},
	},
	{
		ID:        0,
		Depth:     0,
		Feature:   "r",
		Threshold: 233,
		Yes:       1,
		No:        2,
		Missing:   1,
		Children: []*node{
			{
				ID:        1,
				Depth:     1,
				Feature:   "r",
				Threshold: 210,
				Yes:       3,
				No:        4,
				Missing:   3,
				Children: []*node{
					{
						ID:        3,
						Depth:     2,
						Feature:   "r",
						Threshold: 151,
						Yes:       7,
						No:        8,
						Missing:   7,
						Children: []*node{
							{
								ID:   7,
								Leaf: -0.229265645,
							},
							{
								ID:   8,
								Leaf: 0.211816892,
							},
						},
					},
					{
						ID:   4,
						Leaf: -0.377990812,
					},
				},
			},
			{
				ID:        2,
				Depth:     1,
				Feature:   "g",
				Threshold: 207,
				Yes:       5,
				No:        6,
				Missing:   5,
				Children: []*node{
					{
						ID:   5,
						Leaf: 0.395183474,
					},
					{
						ID:        6,
						Depth:     2,
						Feature:   "g",
						Threshold: 230,
						Yes:       9,
						No:        10,
						Missing:   9,
						Children: []*node{
							{
								ID:   9,
								Leaf: -0.235886112,
							},
							{
								ID:   10,
								Leaf: 0.379882932,
							},
						},
					},
				},
			},
		},
	},
	{
		ID:        0,
		Depth:     0,
		Feature:   "r",
		Threshold: 233,
		Yes:       1,
		No:        2,
		Missing:   1,
		Children: []*node{
			{
				ID:        1,
				Depth:     1,
				Feature:   "b",
				Threshold: 162,
				Yes:       3,
				No:        4,
				Missing:   3,
				Children: []*node{
					{
						ID:        3,
						Depth:     2,
						Feature:   "r",
						Threshold: 210,
						Yes:       5,
						No:        6,
						Missing:   5,
						Children: []*node{
							{
								ID:   5,
								Leaf: 0.184997842,
							},
							{
								ID:   6,
								Leaf: -0.188720137,
							},
						},
					},
					{
						ID:   4,
						Leaf: -0.328900367,
					},
				},
			},
			{
				ID:   2,
				Leaf: 0.293951094,
			},
		},
	},
	{
		ID:        0,
		Depth:     0,
		Feature:   "b",
		Threshold: 124,
		Yes:       1,
		No:        2,
		Missing:   1,
		Children: []*node{
			{
				ID:        1,
				Depth:     1,
				Feature:   "b",
				Threshold: 75,
				Yes:       3,
				No:        4,
				Missing:   3,
				Children: []*node{
					{
						ID:        3,
						Depth:     2,
						Feature:   "g",
						Threshold: 136,
						Yes:       7,
						No:        8,
						Missing:   7,
						Children: []*node{
							{
								ID:   7,
								Leaf: 0.209352121,
							},
							{
								ID:   8,
								Leaf: -0.258644491,
							},
						},
					},
					{
						ID:        4,
						Depth:     2,
						Feature:   "r",
						Threshold: 182,
						Yes:       9,
						No:        10,
						Missing:   9,
						Children: []*node{
							{
								ID:   9,
								Leaf: -0.00322336657,
							},
							{
								ID:   10,
								Leaf: -0.538626313,
							},
						},
					},
				},
			},
			{
				ID:        2,
				Depth:     1,
				Feature:   "r",
				Threshold: 183,
				Yes:       5,
				No:        6,
				Missing:   5,
				Children: []*node{
					{
						ID:   5,
						Leaf: -0.173704311,
					},
					{
						ID:        6,
						Depth:     2,
						Feature:   "g",
						Threshold: 218,
						Yes:       11,
						No:        12,
						Missing:   11,
						Children: []*node{
							{
								ID:   11,
								Leaf: 0.0294989403,
							},
							{
								ID:   12,
								Leaf: 0.364735723,
							},
						},
					},
				},
			},
		},
	},
}
