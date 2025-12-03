use std::ops::Range;

pub fn sum_invalid_ids<F>(input: &str, check: &F) -> i64
where
    F: Fn(i64) -> bool,
{
    let ranges = input.split(",").map(|x| {
        let bounds: Vec<i64> = x
            .split("-")
            .map(|y: &str| -> i64 { y.parse().unwrap() })
            .collect();
        bounds[0]..bounds[1] + 1
    });
    let mut invalid_ids: Vec<i64> = Vec::new();
    for range in ranges {
        invalid_ids.extend(find_invalid_ids(range, check));
    }
    invalid_ids.iter().sum()
}

fn find_invalid_ids<F>(r: Range<i64>, check: F) -> Vec<i64>
where
    F: Fn(i64) -> bool,
{
    let mut invalid_ids = Vec::new();
    for id in r {
        if check(id) {
            invalid_ids.push(id);
        }
    }
    invalid_ids
}

fn split_int(i: i64, n: u32) -> Result<Vec<i64>, &'static str> {
    let n_digits: u32 = i.ilog10() + 1;
    if n_digits % n != 0 {
        return Err("Cannot split, because n does not divide the number of digits");
    }
    let mut children = Vec::new();
    let mut dividend = i;
    let base: i64 = 10;
    let divisor = base.pow(n_digits / n);
    for _ in 0..n {
        children.push(dividend % divisor);
        dividend /= divisor;
    }
    Ok(children)
}

pub fn is_invalid(id: i64) -> bool {
    match split_int(id, 2) {
        Ok(parts) => {
            let first = parts[0];
            parts.iter().all(|x| *x == first)
        }
        Err(_) => false,
    }
}

pub fn is_invalid_all_repeats(id: i64) -> bool {
    let n_digits: u32 = id.ilog10() + 1;
    for n in 2..n_digits + 1 {
        if match split_int(id, n) {
            Ok(parts) => {
                let first = parts[0];
                parts.iter().all(|x| *x == first)
            }
            Err(_) => false,
        } {
            return true;
        }
    }
    return false;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sum_invalid_ids_example_input() {
        let input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";
        assert_eq!(sum_invalid_ids(input, &is_invalid), 1227775554);
    }

    #[test]
    fn sum_invalid_ids_all_repeats_example_input() {
        let input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";
        assert_eq!(sum_invalid_ids(input, &is_invalid_all_repeats), 4174379265);
    }

    #[test]
    fn is_invalid_small_even_digit() {
        assert!(is_invalid(11));
    }

    #[test]
    fn is_invalid_small_odd_digit() {
        assert!(!is_invalid(111));
    }

    #[test]
    fn is_invalid_large_even_digit() {
        assert!(is_invalid(1188511885));
    }

    #[test]
    fn is_invalid_large_odd_digit() {
        assert!(!is_invalid(11885118851));
    }

    #[test]
    fn split_int_half_small() {
        assert_eq!(split_int(11, 2).unwrap(), vec![1, 1])
    }

    #[test]
    fn split_int_half_large() {
        assert_eq!(split_int(1188511885, 2).unwrap(), vec![11885, 11885])
    }

    #[test]
    fn split_int_third_small() {
        assert_eq!(split_int(222, 3).unwrap(), vec![2, 2, 2])
    }

    #[test]
    fn split_int_third_large() {
        assert_eq!(split_int(824824824, 3).unwrap(), vec![824, 824, 824])
    }

    #[test]
    fn is_invalid_all_repeats_small_one() {
        assert!(is_invalid_all_repeats(111));
    }

    #[test]
    fn is_invalid_all_repeats_small_two() {
        assert!(is_invalid_all_repeats(1010));
    }

    #[test]
    fn is_invalid_all_repeats_large_one() {
        assert!(is_invalid_all_repeats(222222));
    }

    #[test]
    fn is_invalid_all_repeats_large_two() {
        assert!(is_invalid_all_repeats(565656));
    }
}
