pub fn sum_joltage(inputs: Vec<&str>) -> i64 {
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sum_joltage_example_input() {
        let inputs = vec![
            "987654321111111",
            "811111111111119",
            "234234234234278",
            "818181911112111",
        ];
        assert_eq!(sum_joltage(inputs), 357);
    }
}
