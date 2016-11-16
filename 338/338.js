/**
 * @param {number} num
 * @return {number[]}
 */

var countBits = function(num) {
    var exp = 0;
    var temp = num;
    while (temp > 0) {
        temp = temp >> 1;
        exp++;
    }
    return generate_sequence(exp).slice(0, num + 1);
};

var count_history = {};
var generate_sequence = function(exp) {
    if (exp === 0) {
        return [0, 1];
    }
    if (count_history[exp] !== undefined) {
        return count_history[exp];
    }
    sequence = generate_sequence(exp - 1);
    sequence = sequence.concat(increase(sequence));
    count_history[exp] = sequence;
    return count_history[exp];
}

var increase = function(list) {
    new_list = [];
    for (var key in list) {
        new_list[key] = list[key] + 1;
    }
    return new_list;
}
