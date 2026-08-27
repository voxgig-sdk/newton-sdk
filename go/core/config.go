package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Newton",
			"slug": "newton",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"transport": "base",
			},
		},
		"options": map[string]any{
			"base": "https://newton.now.sh/api/v2",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"abs": map[string]any{},
				"arcco": map[string]any{},
				"arcsin": map[string]any{},
				"arctan": map[string]any{},
				"area": map[string]any{},
				"cos": map[string]any{},
				"derive": map[string]any{},
				"factor": map[string]any{},
				"integrate": map[string]any{},
				"log": map[string]any{},
				"simplify": map[string]any{},
				"sin": map[string]any{},
				"tan": map[string]any{},
				"tangent": map[string]any{},
				"zero": map[string]any{},
			},
		},
		"entity": map[string]any{
			"abs": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "abs",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "-1",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/abs/{expression}",
								"parts": []any{
									"abs",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"arcco": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "arcco",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "1",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/arccos/{expression}",
								"parts": []any{
									"arccos",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"arcsin": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "arcsin",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "0",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/arcsin/{expression}",
								"parts": []any{
									"arcsin",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"arctan": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "arctan",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "0",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/arctan/{expression}",
								"parts": []any{
									"arctan",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"area": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "area",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "2:4|x^3",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/area/{expression}",
								"parts": []any{
									"area",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"cos": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "cos",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "pi",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/cos/{expression}",
								"parts": []any{
									"cos",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"derive": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "derive",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "x^2+2x",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/derive/{expression}",
								"parts": []any{
									"derive",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"factor": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "factor",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "x^2+2x",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/factor/{expression}",
								"parts": []any{
									"factor",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"integrate": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "integrate",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "x^2+2x",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/integrate/{expression}",
								"parts": []any{
									"integrate",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"log": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "log",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "2|8",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/log/{expression}",
								"parts": []any{
									"log",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"simplify": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "simplify",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "2^2+2(2)",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/simplify/{expression}",
								"parts": []any{
									"simplify",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"sin": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "sin",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "0",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sin/{expression}",
								"parts": []any{
									"sin",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"tan": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "tan",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "0",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/tan/{expression}",
								"parts": []any{
									"tan",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"tangent": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "tangent",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "2|x^3",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/tangent/{expression}",
								"parts": []any{
									"tangent",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"zero": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "expression",
						"req": true,
						"short": "The mathematical expression that was processed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"short": "The mathematical operation that was performed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "The result of the mathematical operation",
						"type": "`$STRING`",
					},
				},
				"name": "zero",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "x^2+2x",
											"kind": "param",
											"name": "id",
											"orig": "expression",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/zeroes/{expression}",
								"parts": []any{
									"zeroes",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
