const fs = require('fs');
const modformula = require('./modformula.js');

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
    let formula = modformula.defineFormula(formulaStr);
    let result = modformula.evaluate(formula, scope);
    console.log(`${scheme}: ${result}; profit: ${result - scope["A"]}`);
}

