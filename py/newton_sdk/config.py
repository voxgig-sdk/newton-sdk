# Newton SDK configuration


def make_config():
    return {
        "main": {
            "name": "Newton",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
      },
        },
        "options": {
            "base": "https://newton.now.sh/api/v2",
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "abs": {},
                "arcco": {},
                "arcsin": {},
                "arctan": {},
                "area": {},
                "cos": {},
                "derive": {},
                "factor": {},
                "integrate": {},
                "log": {},
                "simplify": {},
                "sin": {},
                "tan": {},
                "tangent": {},
                "zero": {},
            },
        },
        "entity": {
      "abs": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "abs",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "-1",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/abs/{expression}",
                "parts": [
                  "abs",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "arcco": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "arcco",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "1",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/arccos/{expression}",
                "parts": [
                  "arccos",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "arcsin": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "arcsin",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "0",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/arcsin/{expression}",
                "parts": [
                  "arcsin",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "arctan": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "arctan",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "0",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/arctan/{expression}",
                "parts": [
                  "arctan",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "area": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "area",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "2:4|x^3",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/area/{expression}",
                "parts": [
                  "area",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "cos": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "cos",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "pi",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/cos/{expression}",
                "parts": [
                  "cos",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "derive": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "derive",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "x^2+2x",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/derive/{expression}",
                "parts": [
                  "derive",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "factor": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "factor",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "x^2+2x",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/factor/{expression}",
                "parts": [
                  "factor",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "integrate": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "integrate",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "x^2+2x",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/integrate/{expression}",
                "parts": [
                  "integrate",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "log": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "log",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "2|8",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/log/{expression}",
                "parts": [
                  "log",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "simplify": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "simplify",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "2^2+2(2)",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/simplify/{expression}",
                "parts": [
                  "simplify",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "sin": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "sin",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "0",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/sin/{expression}",
                "parts": [
                  "sin",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "tan": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "tan",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "0",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/tan/{expression}",
                "parts": [
                  "tan",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "tangent": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "tangent",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "2|x^3",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/tangent/{expression}",
                "parts": [
                  "tangent",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "zero": {
        "fields": [
          {
            "active": True,
            "name": "expression",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "operation",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "result",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
        ],
        "name": "zero",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "example": "x^2+2x",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/zeroes/{expression}",
                "parts": [
                  "zeroes",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "expression": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
