use ndarray::Array2;

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
    let mut max_indices = Array2::<usize>::zeros((len, len));
    for i in 0..len {
        max_indices[[i, i]] = i;
        for j in 1..len - i {
            if digits[i + j] > digits[max_indices[[i, i + j - 1]]] {
                max_indices[[i, i + j]] = i + j;
            } else {
                max_indices[[i, i + j]] = max_indices[[i, i + j - 1]];
            }
        }
    }

    let mut j = 0;
    let mut k = len - ndigits;
    let mut index = max_indices[[j, k]];
    let mut digit: u64 = digits[index].into();
    let mut result = digit;
    for i in 1..ndigits {
        j = index + 1;
        k = len - ndigits + i;
        index = max_indices[[j, k]];
        digit = digits[index].into();
        result = result * 10 + digit;
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
    fn sum_joltage_example_input_twelve() {
        let inputs = vec![
            "987654321111111",
            "811111111111119",
            "234234234234278",
            "818181911112111",
        ];
        assert_eq!(sum_joltage(inputs, 12), 3121910778619);
    }

    #[test]
    fn find_largest_joltage_two_small() {
        assert_eq!(find_largest_joltage("12", 2), 12);
    }

    #[test]
    fn find_largest_joltage_two_large() {
        assert_eq!(find_largest_joltage("818181911112111", 2), 92);
    }

    #[test]
    fn find_largest_joltage_twelve_small() {
        assert_eq!(find_largest_joltage("888911112111", 12), 888911112111);
    }

    #[test]
    fn find_largest_joltage_twelve_large() {
        assert_eq!(find_largest_joltage("818181911112111", 12), 888911112111);
    }
}
