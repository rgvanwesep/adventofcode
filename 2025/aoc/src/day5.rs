pub fn count_fresh_ids(inputs: Vec<&str>) -> u64 {
    let mut inputs_iter = inputs.iter();
    let mut fresh_intervals = IntervalSet::new();
    let mut ids: Vec<u64> = Vec::new();
    let mut input: &str;
    let mut split_result: Vec<u64>;
    let mut count = 0;
    loop {
        input = match inputs_iter.next() {
            Some(s) => *s,
            None => "",
        };
        if input == "" {
            break;
        }
        split_result = input.split("-").map(|x| x.parse().unwrap()).collect();
        fresh_intervals.add((split_result[0], split_result[1]));
    }
    loop {
        input = match inputs_iter.next() {
            Some(s) => *s,
            None => "",
        };
        if input == "" {
            break;
        }
        ids.push(input.parse().unwrap());
    }
    for id in ids {
        if fresh_intervals.contains(id) {
            count += 1;
        }
    }
    count
}

pub fn count_total_fresh_ids(inputs: Vec<&str>) -> u64 {
    let mut fresh_intervals = IntervalSet::new();
    let mut split_result: Vec<u64>;
    for input in inputs {
        if input == "" {
            break;
        }
        split_result = input.split("-").map(|x| x.parse().unwrap()).collect();
        fresh_intervals.add((split_result[0], split_result[1]));
    }
    fresh_intervals.size()
}

#[derive(Debug)]
struct IntervalSet {
    intervals: Vec<(u64, u64)>,
}

impl IntervalSet {
    fn new() -> IntervalSet {
        let intervals = Vec::new();
        IntervalSet { intervals }
    }

    fn add(&mut self, interval: (u64, u64)) {
        let mut new_intervals: Vec<(u64, u64)> = Vec::new();
        let mut overlapping: Vec<(u64, u64)> = Vec::new();
        let mut min = interval.0;
        let mut max = interval.1;
        for existing in &self.intervals {
            if existing.1 < min || existing.0 > max {
                new_intervals.push(*existing);
            } else {
                overlapping.push(*existing);
            }
        }
        for ov in overlapping {
            if ov.0 < min {
                min = ov.0
            }
            if ov.1 > max {
                max = ov.1
            }
        }
        new_intervals.push((min, max));
        self.intervals = new_intervals;
    }

    fn size(&self) -> u64 {
        let mut count = 0;
        for interval in &self.intervals {
            count += interval.1 - interval.0 + 1;
        }
        count
    }

    fn contains(&self, i: u64) -> bool {
        for interval in &self.intervals {
            if i >= interval.0 && i <= interval.1 {
                return true;
            }
        }
        false
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn count_fresh_ids_example() {
        let inputs = vec![
            "3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32",
        ];
        assert_eq!(count_fresh_ids(inputs), 3);
    }

    #[test]
    fn count_total_fresh_ids_example() {
        let inputs = vec![
            "3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32",
        ];
        assert_eq!(count_total_fresh_ids(inputs), 14);
    }
}
