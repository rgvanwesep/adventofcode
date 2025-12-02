use std::ops::Range;

pub fn sum_invalid_ids(input: &str) -> i64 {
    let ranges = input.split(",").map(|x| {
        let mut bounds: Vec<i64> = x.split("-").map(|y: &str| -> i64 { y.parse().unwrap() }).collect();
        bounds[0]..bounds[1] + 1
    });
    let mut invalid_ids: Vec<i64> = Vec::new();
    for range in ranges {
        invalid_ids.extend(find_invalid_ids(&range));
    }
    invalid_ids.iter().sum()
}

fn find_invalid_ids(r: &Range<i64>) -> Vec<i64> {
    Vec::new()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sum_invalid_ids_example_input() {
        let input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";
        assert_eq!(sum_invalid_ids(input), 1227775554);
    }
}
