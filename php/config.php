<?php
declare(strict_types=1);

// Newton SDK configuration

class NewtonConfig
{
    /** @var array<string,mixed>|null */
    private static ?array $shared_config = null;

    /**
     * Return the process-wide config, built once on first use. The SDK reads
     * the config on every request and never writes to it, so one instance is
     * shared by every client rather than rebuilt per client.
     *
     * PHP arrays are copy-on-write, so callers that do mutate the result get
     * their own copy and cannot disturb the shared one.
     */
    public static function shared_config(): array
    {
        if (self::$shared_config === null) {
            self::$shared_config = self::make_config();
        }
        return self::$shared_config;
    }

    /**
     * Build a fresh, fully materialised config array. Every call rebuilds the
     * whole structure, so prefer shared_config unless you need a private copy.
     */
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "Newton",
                "slug" => "newton",
                "version" => "0.0.1",
                "target" => "php",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
        ],
            ],
            "options" => [
                "base" => "https://newton.now.sh/api/v2",
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "abs" => [],
                    "arcco" => [],
                    "arcsin" => [],
                    "arctan" => [],
                    "area" => [],
                    "cos" => [],
                    "derive" => [],
                    "factor" => [],
                    "integrate" => [],
                    "log" => [],
                    "simplify" => [],
                    "sin" => [],
                    "tan" => [],
                    "tangent" => [],
                    "zero" => [],
                ],
            ],
            "entity" => [
        'abs' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'abs',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '-1',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/abs/{expression}',
                  'parts' => [
                    'abs',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'arcco' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'arcco',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '1',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/arccos/{expression}',
                  'parts' => [
                    'arccos',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'arcsin' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'arcsin',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '0',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/arcsin/{expression}',
                  'parts' => [
                    'arcsin',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'arctan' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'arctan',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '0',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/arctan/{expression}',
                  'parts' => [
                    'arctan',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'area' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'area',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '2:4|x^3',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/area/{expression}',
                  'parts' => [
                    'area',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'cos' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'cos',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'pi',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/cos/{expression}',
                  'parts' => [
                    'cos',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'derive' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'derive',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'x^2+2x',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/derive/{expression}',
                  'parts' => [
                    'derive',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'factor' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'factor',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'x^2+2x',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/factor/{expression}',
                  'parts' => [
                    'factor',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'integrate' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'integrate',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'x^2+2x',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/integrate/{expression}',
                  'parts' => [
                    'integrate',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'log' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'log',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '2|8',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/log/{expression}',
                  'parts' => [
                    'log',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'simplify' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'simplify',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '2^2+2(2)',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/simplify/{expression}',
                  'parts' => [
                    'simplify',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'sin' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'sin',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '0',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/sin/{expression}',
                  'parts' => [
                    'sin',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'tan' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'tan',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '0',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/tan/{expression}',
                  'parts' => [
                    'tan',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'tangent' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'tangent',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '2|x^3',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/tangent/{expression}',
                  'parts' => [
                    'tangent',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'zero' => [
          'fields' => [
            [
              'name' => 'expression',
              'req' => true,
              'short' => 'The mathematical expression that was processed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'operation',
              'req' => true,
              'short' => 'The mathematical operation that was performed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'result',
              'req' => true,
              'short' => 'The result of the mathematical operation',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'zero',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'x^2+2x',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'expression',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/zeroes/{expression}',
                  'parts' => [
                    'zeroes',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'expression' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return NewtonFeatures::make_feature($name);
    }
}
