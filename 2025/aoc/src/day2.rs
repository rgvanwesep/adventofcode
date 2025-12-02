use std::ops::Range;

pub fn sum_invalid_ids(input: &str) -> i64 {
    let ranges = input.split(",").map(|x| {
        let bounds: Vec<i64> = x
            .split("-")
            .map(|y: &str| -> i64 { y.parse().unwrap() })
            .collect();
        bounds[0]..bounds[1] + 1
    });
    let mut invalid_ids: Vec<i64> = Vec::new();
    for range in ranges {
        invalid_ids.extend(find_invalid_ids(range));
    }
    invalid_ids.iter().sum()
}

fn find_invalid_ids(r: Range<i64>) -> Vec<i64> {
    let mut invalid_ids = Vec::new();
    for id in r {
        if is_invalid(id) {
            invalid_ids.push(id);
        }
    }
    invalid_ids
}

fn is_invalid(id: i64) -> bool {
    let n_digits: u32 = id.ilog10()+1;
    if n_digits % 2 == 0 {
        let base: i64 = 10;
        let divisor = base.pow(n_digits / 2);
        id / divisor == id % divisor
    } else {
        false
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sum_invalid_ids_example_input() {
        let input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";
        assert_eq!(sum_invalid_ids(input), 1227775554);
    }

    #[test]
    fn is_invalid_small_even_digit() {
        assert!(is_invalid(11));
    }

    #[test]
    fn is_invalid_small_odd_digit() {
        assert!(!is_invalid(111));
    }

    fn is_invalid_large_even_digit() {
        assert!(is_invalid(1188511885));
    }

    #[test]
    fn is_invalid_large_odd_digit() {
        assert!(!is_invalid(11885118851));
    }
}
