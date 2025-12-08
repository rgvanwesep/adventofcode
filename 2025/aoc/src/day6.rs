use ndarray::{Array2, s};

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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sum_results_example() {
        let inputs = vec![
            "123 328  51 64 ",
            " 45 64  387 23 ",
            "  6 98  215 314",
            "*   +   *   + ",
        ];
        assert_eq!(sum_results(inputs), 4277556);
    }
}
