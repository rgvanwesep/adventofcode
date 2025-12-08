const SOURCE: u8 = b'S';
const SPLITTER: u8 = b'^';
const BEAM: u8 = b'|';

pub fn count_splits(inputs: Vec<&str>) -> u32 {
    let mut count = 0;

    let nrows = inputs.len();
    let ncols = inputs[0].bytes().count();
    let mut bytes = ndarray::Array2::<u8>::zeros((nrows, ncols));
    for (i, row) in inputs.iter().enumerate() {
        for (j, byte) in row.bytes().enumerate() {
            bytes[[i, j]] = byte
        }
    }

    for j in 0..ncols {
        if bytes[[0, j]] == SOURCE {
            bytes[[1, j]] = BEAM;
            break;
        }
    }

    for i in 2..nrows {
        for j in 0..ncols {
            if bytes[[i - 1, j]] == BEAM {
                if bytes[[i, j]] == SPLITTER {
                    count += 1;
                    bytes[[i, j - 1]] = BEAM;
                    bytes[[i, j + 1]] = BEAM;
                } else {
                    bytes[[i, j]] = BEAM;
                }
            }
        }
    }

    count
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn count_splits_example() {
        let inputs = vec![
            ".......S.......",
            "...............",
            ".......^.......",
            "...............",
            "......^.^......",
            "...............",
            ".....^.^.^.....",
            "...............",
            "....^.^...^....",
            "...............",
            "...^.^...^.^...",
            "...............",
            "..^...^.....^..",
            "...............",
            ".^.^.^.^.^...^.",
            "...............",
        ];
        assert_eq!(count_splits(inputs), 21);
    }
}
