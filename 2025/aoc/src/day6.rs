use ndarray::{Array2, s};

const ADD: u8 = b'+';
const MULTIPLY: u8 = b'*';
const EMPTY: u8 = b' ';
const ZERO: u8 = b'0';

pub fn sum_results(inputs: Vec<&str>) -> u64 {
    let nrows = inputs.len() - 1;
    let ncols = inputs[0].split_whitespace().count();
    let mut numbers = Array2::<u64>::zeros((nrows, ncols));
    let mut sum = 0;
    for (i, row) in inputs[..nrows].iter().enumerate() {
        for (j, value) in row
            .split_whitespace()
            .map(|s| s.parse().unwrap())
            .enumerate()
        {
            numbers[[i, j]] = value;
        }
    }
    for (i, op) in inputs[nrows].split_whitespace().enumerate() {
        sum += match op {
            "+" => numbers.slice(s![.., i]).sum(),
            "*" => numbers.slice(s![.., i]).product(),
            _ => panic!("Unknown operation!"),
        }
    }
    sum
}

pub fn sum_results_cepha(inputs: Vec<&str>) -> u64 {
    let mut sum = 0;

    let nrows = inputs.len();
    let ncols = inputs[0].bytes().count();
    let mut bytes = Array2::<u8>::zeros((nrows, ncols));
    for (i, row) in inputs.iter().enumerate() {
        for (j, byte) in row.bytes().enumerate() {
            bytes[[i, j]] = byte;
        }
    }

    let mut operands = Vec::<u64>::new();
    let mut operand: u64;
    let mut operator: u8;
    for col in bytes.columns().into_iter().rev() {
        operand = 0;
        for &byte in col {
            if byte > ZERO {
                operand = operand * 10 + (byte - ZERO) as u64;
            }
        }
        if operand != 0 {
            operands.push(operand);
        }
        operator = *col.last().unwrap();
        if operator != EMPTY {
            sum += match operator {
                ADD => operands.iter().sum::<u64>(),
                MULTIPLY => operands.iter().product(),
                _ => panic!("Unknown operator!"),
            };
            operands.clear();
        }
    }
    sum
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sum_results_example() {
        let inputs = vec![
            "123 328  51 64 ",
            " 45 64  387 23 ",
            "  6 98  215 314",
            "*   +   *   +  ",
        ];
        assert_eq!(sum_results(inputs), 4277556);
    }

    #[test]
    fn sum_results_cepha_example() {
        let inputs = vec![
            "123 328  51 64 ",
            " 45 64  387 23 ",
            "  6 98  215 314",
            "*   +   *   +  ",
        ];
        assert_eq!(sum_results_cepha(inputs), 3263827);
    }
}
