# Newton SDK configuration


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "Newton",
            "slug": "newton",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
        "transport": "base",
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
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "abs",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "-1",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "arcco": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "arcco",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "1",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "arcsin": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "arcsin",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "0",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "arctan": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "arctan",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "0",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "area": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "area",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "2:4|x^3",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "cos": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "cos",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "pi",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "derive": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "derive",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "x^2+2x",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "factor": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "factor",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "x^2+2x",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "integrate": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "integrate",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "x^2+2x",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "log": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "log",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "2|8",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "simplify": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "simplify",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "2^2+2(2)",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "sin": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "sin",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "0",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "tan": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "tan",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "0",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "tangent": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "tangent",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "2|x^3",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "zero": {
        "fields": [
          {
            "name": "expression",
            "req": True,
            "short": "The mathematical expression that was processed",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "operation",
            "req": True,
            "short": "The mathematical operation that was performed",
            "type": "`$STRING`",
          },
          {
            "name": "result",
            "req": True,
            "short": "The result of the mathematical operation",
            "type": "`$STRING`",
          },
        ],
        "name": "zero",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "x^2+2x",
                      "kind": "param",
                      "name": "id",
                      "orig": "expression",
                      "reqd": True,
                      "type": "`$STRING`",
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
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
