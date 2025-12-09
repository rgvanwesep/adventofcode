use ndarray::Array2;

const SOURCE: u8 = b'S';
const SPLITTER: u8 = b'^';
const BEAM: u8 = b'|';

pub fn count_splits(inputs: Vec<&str>) -> u32 {
    Graph::build(inputs).splitter_count
}

pub fn count_paths(inputs: Vec<&str>) -> u64 {
    Graph::build(inputs).count_paths()
}

#[derive(Default)]
struct Node {
    value: u8,
    neighbors: Vec<(usize, usize)>,
}

struct Graph {
    head: Option<(usize, usize)>,
    nodes: Array2<Node>,
    splitter_count: u32,
}

impl Graph {
    fn build(inputs: Vec<&str>) -> Graph {
        let mut splitter_count = 0;

        let nrows = inputs.len();
        let ncols = inputs[0].bytes().count();
        let mut nodes = Array2::<Node>::default((nrows, ncols));
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

        Graph {
            head,
            nodes,
            splitter_count,
        }
    }

    fn count_paths(&self) -> u64 {
        let mut stack = Vec::<(usize, usize)>::new();
        let mut path_counts = Array2::<u64>::zeros((self.nodes.nrows(), self.nodes.ncols()));
        let mut i;
        let mut j;
        let mut node: &Node;
        match self.head {
            Some(pair) => {
                stack.push(pair);
                while !stack.is_empty() {
                    (i, j) = stack.pop().unwrap();
                    node = &self.nodes[[i, j]];
                    if node.neighbors.is_empty() {
                        path_counts[[i, j]] = 1;
                    } else {
                        if node.neighbors.iter().all(|&(k, l)| path_counts[[k, l]] > 0) {
                            path_counts[[i, j]] = node
                                .neighbors
                                .iter()
                                .fold(0, |s, &(k, l)| s + path_counts[[k, l]])
                        } else {
                            stack.push((i, j));
                            stack.extend(
                                node.neighbors
                                    .iter()
                                    .filter(|&&(k, l)| path_counts[[k, l]] == 0),
                            )
                        }
                    }
                }
                path_counts[[pair.0, pair.1]]
            }
            None => 0,
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
