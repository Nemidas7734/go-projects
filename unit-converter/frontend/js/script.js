// Length conversion logic
function convertLength() {
    const value = parseFloat(document.getElementById('value').value);
    const fromUnit = document.getElementById('fromUnit').value;
    const toUnit = document.getElementById('toUnit').value;
    const resultDiv = document.getElementById('result');
    const UnitType = "length"

    // Conversion factors for length
    const lengthConversions = {
        "millimeter": 1,
        "centimeter": 10,
        "meter": 1000,
        "kilometer": 1000000,
        "inch": 25.4,
        "foot": 304.8,
        "yard": 914.4,
        "mile": 1609344
    };

    if (isNaN(value) || value === '') {
        resultDiv.innerHTML = 'Please enter a valid number!';
        return;
    }

    fetch('http://localhost:8090/convert', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            value: value,
            fromUnit: fromUnit,
            toUnit: toUnit,
            UnitType: UnitType
        })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('result').innerHTML = `Converted Value: ${data.result}`;
    })
    .catch(error => {
        document.getElementById('result').innerHTML = `Error: ${error.message}`;
    });

}

// Weight conversion logic
function convertWeight() {
    const value = parseFloat(document.getElementById('value').value);
    const fromUnit = document.getElementById('fromUnit').value;
    const toUnit = document.getElementById('toUnit').value;
    const resultDiv = document.getElementById('result');
    const UnitType = "weight"

    // Conversion factors for weight
    const weightConversions = {
        "milligram": 1,
        "gram": 1000,
        "kilogram": 1000000,
        "ounce": 28349.5,
        "pound": 453592
    };

    if (isNaN(value) || value === '') {
        resultDiv.innerHTML = 'Please enter a valid number!';
        return;
    }


    fetch('http://localhost:8090/convert', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            value: value,
            fromUnit: fromUnit,
            toUnit: toUnit,
            UnitType: UnitType
        })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('result').innerHTML = `Converted Value: ${data.result}`;
    })
    .catch(error => {
        document.getElementById('result').innerHTML = `Error: ${error.message}`;
    });

}

// Temperature conversion logic
function convertTemperature() {
    const value = parseFloat(document.getElementById('value').value);
    const fromUnit = document.getElementById('fromUnit').value;
    const toUnit = document.getElementById('toUnit').value;
    const resultDiv = document.getElementById('result');
    const UnitType = "temperature"

    if (isNaN(value) || value === '') {
        resultDiv.innerHTML = 'Please enter a valid number!';
        return;
    }

    let tempValue;

    // Convert from the input unit to Celsius first
    if (fromUnit === 'celsius') {
        tempValue = value;
    } else if (fromUnit === 'fahrenheit') {
        tempValue = (value - 32) * 5 / 9;
    } else if (fromUnit === 'kelvin') {
        tempValue = value - 273.15;
    }

    let convertedValue;
    if (toUnit === 'celsius') {
        convertedValue = tempValue;
    } else if (toUnit === 'fahrenheit') {
        convertedValue = (tempValue * 9 / 5) + 32;
    } else if (toUnit === 'kelvin') {
        convertedValue = tempValue + 273.15;
    }

    // resultDiv.innerHTML = `${value} ${fromUnit} = ${convertedValue} ${toUnit}`;
    fetch('http://localhost:8090/convert', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            value: value,
            fromUnit: fromUnit,
            toUnit: toUnit,
            UnitType: UnitType
        })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('result').innerHTML = `Converted Value: ${data.result}`;
    })
    .catch(error => {
        document.getElementById('result').innerHTML = `Error: ${error.message}`;
    });

}
