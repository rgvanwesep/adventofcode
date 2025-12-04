pub fn sum_joltage(inputs: Vec<&str>, digits: usize) -> u64 {
    let mut sum = 0;
    for input in inputs {
        sum += find_largest_joltage(input, digits)
    }
    sum
}

fn find_largest_joltage(input: &str, digits: usize) -> u64 {
    let bytes = input.as_bytes();
    let len = bytes.len();
    let mut max_after = vec![0; len - 1];
    let mut max_indices = vec![len; len - 1];
    for i in 0..digits - 1 {
        max_after[len - i - 2] = bytes[len - i - 1];
        max_indices[len - i - 2] = len - i - 1;
    }
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
            max_indices[i] = i + 1;
        } else {
            max_after[i] = max_after[i + 1];
            max_indices[i] = max_indices[i + 1];
        }
    }
    let mut result = (max - '0' as u8).into();
    let mut next_max: u64 = (max_after[max_index] - '0' as u8).into();
    result = result * 10 + next_max;
    for i in 0..digits - 2 {
        next_max = (max_after[max_indices[max_index]] - '0' as u8).into();
        result = result * 10 + next_max;
        max_index = max_indices[max_index];
    }
    result
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sum_joltage_example_input_two() {
        let inputs = vec![
            "987654321111111",
            "811111111111119",
            "234234234234278",
            "818181911112111",
        ];
        assert_eq!(sum_joltage(inputs, 2), 357);
    }

    #[test]
    fn find_largest_joltage_small_two() {
        assert_eq!(find_largest_joltage("12", 2), 12);
    }

    #[test]
    fn find_largest_joltage_small_three() {
        assert_eq!(find_largest_joltage("123", 3), 123);
    }
}
