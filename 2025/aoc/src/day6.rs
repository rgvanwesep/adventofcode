pub fn sum_results(inputs: Vec<&str>) -> u64 {
    0
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
