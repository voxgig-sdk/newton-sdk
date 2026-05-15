package core

func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Newton",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://newton.now.sh/api/v2",
			"auth": map[string]any{
				"prefix": "Bearer",
			},
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "abs",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "arcco",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "arcsin",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "arctan",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "area",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "cos",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "derive",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "factor",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "integrate",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "log",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "simplify",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "sin",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "tan",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "tangent",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "operation",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "zero",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
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
