use ndarray::Array2;

const ROLL: u8 = b'@';
const EMPTY: u8 = b'.';

pub fn count_rolls(inputs: Vec<&str>) -> usize {
    let (nrows, ncols, grid) = build_grid(inputs);
    find_removable(nrows, ncols, &grid).len()
}

pub fn count_removable_rolls(inputs: Vec<&str>) -> usize {
    let (nrows, ncols, mut grid) = build_grid(inputs);
    let mut count = 0;
    loop {
        let pairs = find_removable(nrows, ncols, &grid);
        if pairs.len() == 0 {
            break;
        }
        count += pairs.len();
        for (i, j) in pairs {
            grid[[i, j]] = EMPTY;
        }
    }
    count
}

fn find_removable(nrows: usize, ncols: usize, grid: &Array2<u8>) -> Vec<(usize, usize)> {
    let mut pairs = Vec::new();
    let mut nneighbors: u8;
    for i in 0..nrows {
        for j in 0..ncols {
            if grid[[i + 1, j + 1]] == ROLL {
                nneighbors = 0;
                for k in 0..3 {
                    for l in 0..3 {
                        if !(k == 1 && l == 1) && grid[[i + k, j + l]] == ROLL {
                            nneighbors += 1;
                        }
                    }
                }
                if nneighbors < 4 {
                    pairs.push((i + 1, j + 1))
                }
            }
        }
    }
    pairs
}

fn build_grid(inputs: Vec<&str>) -> (usize, usize, Array2<u8>) {
    let nrows = inputs.len();
    let ncols = inputs[0].len();
    let mut grid = Array2::from_elem((nrows + 2, ncols + 2), EMPTY);
    let mut row: &[u8];
    for i in 0..nrows {
        row = inputs[i].as_bytes();
        for j in 0..ncols {
            if row[j] == ROLL {
                grid[[i + 1, j + 1]] = ROLL;
            }
        }
    }
    (nrows, ncols, grid)
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

    #[test]
    fn count_removable_rolls_example() {
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
        assert_eq!(count_removable_rolls(inputs), 43);
    }

    #[test]
    fn count_rolls_small_zero() {
        let inputs = vec!["...", ".@.", "..."];
        assert_eq!(count_rolls(inputs), 1);
    }

    #[test]
    fn count_rolls_small_two() {
        let inputs = vec!["@..", ".@.", "..."];
        assert_eq!(count_rolls(inputs), 2);
    }

    #[test]
    fn count_rolls_small_three() {
        let inputs = vec!["@@.", ".@.", "..."];
        assert_eq!(count_rolls(inputs), 3);
    }

    #[test]
    fn count_rolls_small_four() {
        let inputs = vec!["@@@", ".@.", "..."];
        assert_eq!(count_rolls(inputs), 4);
    }

    #[test]
    fn count_rolls_small_five() {
        let inputs = vec!["@@@", ".@@", "..."];
        assert_eq!(count_rolls(inputs), 3);
    }
    #[test]
    fn count_rolls_small_six() {
        let inputs = vec!["@@@", ".@@", "..@"];
        assert_eq!(count_rolls(inputs), 3);
    }

    #[test]
    fn count_rolls_small_seven() {
        let inputs = vec!["@@@", ".@@", ".@@"];
        assert_eq!(count_rolls(inputs), 4);
    }
}
