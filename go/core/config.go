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
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "abs",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"abs",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "arccos",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"arccos",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "arcsin",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"arcsin",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "arctan",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"arctan",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "area",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"area",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "cos",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"cos",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "derive",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"derive",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "factor",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"factor",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "integrate",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"integrate",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "log",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"log",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "simplify",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"simplify",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "sin",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"sin",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "tan",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"tan",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "tangent",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"tangent",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"expression": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "zeroes",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"zeroes",
									"{id}",
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

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
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
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
