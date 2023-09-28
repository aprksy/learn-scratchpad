// Import the math.js library
const math = require('mathjs');
const fs = require('fs');

function defineFormula(mathExpression) {
    return math.parse(mathExpression).compile()
}

function evaluate(formula, scope) {
    return formula.evaluate(scope)
}

// Get the command-line arguments
const args = process.argv.slice(2);

let rowId = args[0];

if (isNaN(rowId) || rowId < 0) {
  rowId = 0;
}

let dataFile = "data.json";
let pricingFile = "pricing-scheme.json";
let columnMappingFile = "column-mapping.json";

let dataContent;
try {
    dataContent = fs.readFileSync(dataFile, 'utf8')
} catch(error) {
    console.error(error)
}

let pricingContent;
try {
    pricingContent = fs.readFileSync(pricingFile, 'utf8')
} catch(error) {
    console.error(error)
}

let columnMappingContent;
try {
    columnMappingContent = fs.readFileSync(columnMappingFile, 'utf8')
} catch(error) {
    console.error(error)
}

let data = JSON.parse(dataContent)["data"][rowId];
let columnMapping = JSON.parse(columnMappingContent);

let schemes = [
    "normal_price",
    "regular_discount",
    "flash_sale",
    "super_sale",
];

let formulaStrs = JSON.parse(pricingContent);

// create scope -   assign value to column mappings
let scope = {}
for (const [variable, column] of Object.entries(columnMapping)) {
    scope[variable] = data[column]
}

// evaluate formulas
console.log(`item name: ${data.name}`)
for (const [scheme, formulaStr] of Object.entries(formulaStrs)) {
    let formula = defineFormula(formulaStr);
    let result = evaluate(formula, scope);
    console.log(`${scheme}: ${result}; profit: ${result - scope["A"]}`);
}

