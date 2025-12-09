const SOURCE: u8 = b'S';
const SPLITTER: u8 = b'^';
const BEAM: u8 = b'|';

pub fn count_splits(inputs: Vec<&str>) -> u32 {
    Tree::build(inputs).splitter_count
}

pub fn count_paths(inputs: Vec<&str>) -> u32 {
    Tree::build(inputs).count_paths()
}

#[derive(Default)]
struct Node {
    value: u8,
    neighbors: Vec<(usize, usize)>,
}

struct Tree {
    head: Option<(usize, usize)>,
    nodes: ndarray::Array2<Node>,
    splitter_count: u32,
}

impl Tree {
    fn build(inputs: Vec<&str>) -> Tree {
        let mut splitter_count = 0;

        let nrows = inputs.len();
        let ncols = inputs[0].bytes().count();
        let mut nodes = ndarray::Array2::<Node>::default((nrows, ncols));
        for (i, row) in inputs.iter().enumerate() {
            for (j, byte) in row.bytes().enumerate() {
                nodes[[i, j]] = Node {
                    value: byte,
                    neighbors: Vec::new(),
                }
            }
        }

        let mut head: Option<(usize, usize)> = None;
        for j in 0..ncols {
            if nodes[[0, j]].value == SOURCE {
                head = Some((0, j));
                nodes[[0, j]].neighbors.push((1, j));
                nodes[[1, j]].value = BEAM;
                break;
            }
        }

        for i in 2..nrows {
            for j in 0..ncols {
                if nodes[[i - 1, j]].value == BEAM {
                    nodes[[i - 1, j]].neighbors.push((i, j));
                    if nodes[[i, j]].value == SPLITTER {
                        splitter_count += 1;
                        nodes[[i, j]].neighbors.extend([(i, j - 1), (i, j + 1)]);
                        nodes[[i, j - 1]].value = BEAM;
                        nodes[[i, j + 1]].value = BEAM;
                    } else {
                        nodes[[i, j]].value = BEAM;
                    }
                }
            }
        }

        Tree {
            head,
            nodes,
            splitter_count,
        }
    }

    fn count_paths(&self) -> u32 {
        let mut count = 0;
        let mut paths = Vec::<&Node>::new();
        let mut node: &Node;
        match self.head {
            Some((i, j)) => {
                paths.push(&self.nodes[[i, j]]);
                while !paths.is_empty() {
                    node = paths.pop().unwrap();
                    if node.neighbors.is_empty() {
                        count += 1;
                    } else {
                        for (i, j) in node.neighbors.iter() {
                            paths.push(&self.nodes[[*i, *j]]);
                        }
                    }
                }
                count
            }
            None => count,
        }
    }
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

    #[test]
    fn count_paths_example() {
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
        assert_eq!(count_paths(inputs), 40);
    }
}
