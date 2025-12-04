pub fn sum_joltage(inputs: Vec<&str>, ndigits: usize) -> u64 {
    let mut sum = 0;
    for input in inputs {
        sum += find_largest_joltage(input, ndigits)
    }
    sum
}

fn find_largest_joltage(input: &str, ndigits: usize) -> u64 {
    let digits: Vec<u8> = input.as_bytes().iter().map(|x| x - '0' as u8).collect();
    let len = digits.len();
    let mut maxes = vec![0; len];
    let mut next_max_indices = vec![len; len - 1];
    let mut current_max_index = len - 1;

    for i in 0..ndigits {
        maxes[len - i - 1] = digits[len - i - 1];
    }

    for i in (0..len - 1).rev() {
        next_max_indices[i] = current_max_index;
        if digits[i] >= maxes[i + 1] {
            maxes[i] = digits[i];
            current_max_index = i;
        }
    }
    let mut result = maxes[0] as u64;
    let mut i = 0;
    for _ in 0..ndigits - 1 {
        result = result * 10 + maxes[next_max_indices[i]] as u64;
        i = next_max_indices[i];
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
