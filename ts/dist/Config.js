"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.FEATURE_PLUGINS = exports.config = void 0;
const TestFeature_1 = require("./feature/test/TestFeature");
const FEATURE_CLASS = {
    test: TestFeature_1.TestFeature,
};
// Per-feature plugin DEFINITIONS (voxgig/plugin `Definition` values), from
// the model's active plugin groups. A feature that takes a `plugins` option
// (secrets over sekreto) reads its own entry; a feature with no plugins has
// none. Named imports above make each definition statically reachable, so
// an SDK carries exactly the plugin modules its model selects — the same
// leanness the old side-effect registry imports bought, without a registry.
const FEATURE_PLUGINS = {};
exports.FEATURE_PLUGINS = FEATURE_PLUGINS;
class Config {
    makeFeature(fn) {
        const fc = FEATURE_CLASS[fn];
        const fi = new fc();
        // TODO: errors etc
        return fi;
    }
    // False for a feature added at runtime via options.extend (station's
    // adopt path) - the constructor uses this to skip makeFeature for names
    // no generated class backs.
    hasFeature(fn) {
        return null != FEATURE_CLASS[fn];
    }
    main = {
        name: 'Newton',
        slug: "newton",
        version: "0.0.1",
        target: "ts",
    };
    feature = {
        test: {
            "options": {
                "active": false
            },
            "transport": "base"
        },
    };
    options = {
        base: "https://newton.now.sh/api/v2",
        headers: {
            "content-type": "application/json"
        },
        entity: {
            abs: {},
            arcco: {},
            arcsin: {},
            arctan: {},
            area: {},
            cos: {},
            derive: {},
            factor: {},
            integrate: {},
            log: {},
            simplify: {},
            sin: {},
            tan: {},
            tangent: {},
            zero: {},
        }
    };
    entity = {
        "abs": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/abs/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "abs"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "abs",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "arcco": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/arccos/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "arccos"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "arccos",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "arcsin": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/arcsin/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "arcsin"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "arcsin",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "arctan": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/arctan/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "arctan"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "arctan",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "area": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/area/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "area"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "area",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "cos": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/cos/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "cos"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "cos",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "derive": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/derive/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "derive"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "derive",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "factor": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/factor/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "factor"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "factor",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "integrate": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/integrate/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "integrate"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "integrate",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "log": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/log/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "log"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "log",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "simplify": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/simplify/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "simplify"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "simplify",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "sin": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/sin/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "sin"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "sin",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "tan": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/tan/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "tan"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "tan",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "tangent": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/tangent/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "tangent"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "tangent",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "zero": {
            "fields": [
                {
                    "name": "expression",
                    "req": true,
                    "short": "The mathematical expression that was processed",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "name": "operation",
                    "req": true,
                    "short": "The mathematical operation that was performed",
                    "type": "`$STRING`"
                },
                {
                    "name": "result",
                    "req": true,
                    "short": "The result of the mathematical operation",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
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
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/zeroes/{expression}",
                            "rename": {
                                "param": {
                                    "expression": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "zeroes"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "zeroes",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        }
    };
}
const config = new Config();
exports.config = config;
//# sourceMappingURL=Config.js.map