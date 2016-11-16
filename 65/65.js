a/**
 * @param {string} s
 * @return {boolean}
 */

var isNumber = function(s) {
    s = s.trim();
    if (s.match(/^[0-9\.+\-e]+$/) === null) {
        return false;
    }
    if (isInteger(s)) {
        return true;
    }
    if (isDecimal(s)) {
        return true;
    }
    if (isScientificNotation(s)) {
        return true;
    }
    return false;
};

var isUnsignedInt = function(s) {
    if (s.match(/^[0-9]+$/)) {
        return true;
    }
    return false;
};

var isInteger = function(s) {
    if (s.match(/^[+\-]?[0-9]+$/) !== null) {
        return true;
    }
    return false;
};

var isDecimal = function(s) {
    if (isInteger(s)) {
        return true;
    }
    if (s.match(/^[+\-]?[0-9]*\.[0-9]+$/) !== null ||
        s.match(/^[+\-]?[0-9]+\.[0-9]*$/) !== null) {
        return true;
    }
    return false;
};

var isScientificNotation = function(s) {
    slices = s.split('e');
    if (slices.length != 2) {
        return false;
    }
    if (!isDecimal(slices[0]) || !isInteger(slices[1])) {
        console.log(3);
        return false;
    }
    return true;
};
