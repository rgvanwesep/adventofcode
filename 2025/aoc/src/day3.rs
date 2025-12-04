pub fn sum_joltage(inputs: Vec<&str>) -> i64 {
    let mut sum = 0;
    for input in inputs {
        sum += find_largest_joltage(input)
    }
    sum
}

fn find_largest_joltage(input: &str) -> i64 {
    let bytes = input.as_bytes();
    let len = bytes.len();
    let mut max_after = vec![0; len - 1];
    max_after[len - 2] = bytes[len - 1];
    let mut max = bytes[len - 2];
    let mut max_index = len - 2;
    for i in (0..len - 2).rev() {
        if bytes[i] > max {
            max = bytes[i];
            max_index = i;
        } else if bytes[i] == max {
            max_index = i;
        }
        if bytes[i + 1] > max_after[i + 1] {
            max_after[i] = bytes[i + 1];
        } else {
            max_after[i] = max_after[i + 1];
        }
    }
    max -= '0' as u8;
    (max * 10 + max_after[max_index] - '0' as u8).into()
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

    #[test]
    fn find_largest_joltage_small() {
        assert_eq!(find_largest_joltage("12"), 12);
    }
}
