const math = require('mathjs');

function defineFormula(mathExpression) {
    return math.parse(mathExpression).compile()
}

function evaluate(formula, scope) {
    return formula.evaluate(scope)
}

module.exports = {defineFormula, evaluate}