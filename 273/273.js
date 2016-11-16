/**
 * @param {number} num
 * @return {string}
 */
var numberToWords = function(num) {
    if (num === 0) {
        return 'Zero';
    }
    if (num < 0) {
        return 'Negative ' + chunks_to_string(split_into_chunks(-num));
    }
    return chunks_to_string(split_into_chunks(num));
};

var split_into_chunks = function(num) {
    var chunks = [];
    while (num > 0) {
        chunks = [num % 1000].concat(chunks);
        num = parseInt(num / 1000);
    }
    return chunks;
};

var chunks_to_string = function(chunks) {
    var scales = [' Quadrillion', ' Trillion', ' Billion', ' Million', ' Thousand', ''];
    var digits = {0: 'Zero', 1: 'One', 2: 'Two', 3: 'Three',
                  4: 'Four', 5: 'Five', 6: 'Six',
                  7: 'Seven', 8: 'Eight', 9: 'Nine',
                  10: 'Ten', 11: 'Eleven', 12: 'Twelve',
                  13: 'Thirteen', 14: 'Fourteen', 15: 'Fifteen',
                  16: 'Sixteen', 17: 'Seventeen', 18: 'Eighteen',
                  19: 'Nineteen'};
    var tens = {2: 'Twenty', 3: 'Thirty', 4: 'Forty', 5: 'Fifty',
                6: 'Sixty', 7: 'Seventy', 8: 'Eighty', 9: 'Ninety'};

    var strs = [];
    for (var i in chunks) {
        strs[i] = '';
        if (chunks[i] >= 100) {
            strs[i] = digits[parseInt(chunks[i]/100)] + ' Hundred';
            if (chunks[i] % 100 === 0) {
                continue;
            } else {
                strs[i] += ' ';
            }
        }
        chunks[i] = chunks[i] % 100;
        if (chunks[i] >= 20) {
            strs[i] += tens[parseInt(chunks[i]/10)];
            if (chunks[i] % 10 === 0) {
                continue;
            } else {
                chunks[i] = chunks[i] % 10;
                strs[i] += ' ';
            }
        }
        strs[i] += digits[chunks[i]];
    }


    var temp_array = [];
    while (strs.length > 0 && scales.length > 0) {
        str = strs.pop();
        scale = scales.pop();
        if (str == 'Zero') {
            continue;
        }
        temp_array = [].concat(str + scale, temp_array);
    }
    return temp_array.join(' ');
};
