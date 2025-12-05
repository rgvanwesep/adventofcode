pub fn count_rolls(inputs: Vec<&str>) -> u32 {
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn count_rolls_example() {
        let inputs = vec![
            "..@@.@@@@.",
            "@@@.@.@.@@",
            "@@@@@.@.@@",
            "@.@@@@..@.",
            "@@.@@@@.@@",
            ".@@@@@@@.@",
            ".@.@.@.@@@",
            "@.@@@.@@@@",
            ".@@@@@@@@.",
            "@.@.@@@.@.",
        ];
        assert_eq!(count_rolls(inputs), 13);
    }
}
