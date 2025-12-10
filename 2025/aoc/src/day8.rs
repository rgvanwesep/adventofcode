use petgraph::{
    algo::tarjan_scc,
    graph::{NodeIndex, UnGraph},
};

pub fn multiply_circuit_sizes(inputs: Vec<&str>, num_connections: usize) -> usize {
    let mut points = Vec::<Point3d>::new();
    let mut distances = Vec::<(usize, usize, u64)>::new();
    for input in inputs {
        points.push(Point3d::build(
            input.split(",").map(|s| s.parse().unwrap()).collect(),
        ));
    }
    for (i, pi) in points.iter().enumerate() {
        for (j, pj) in points[0..i].iter().enumerate() {
            distances.push((i, j, pi.distance_squared(pj)));
        }
    }
    distances.sort_by_key(|&x| x.2);
    let graph = UnGraph::<u32, ()>::from_edges(
        distances[0..num_connections]
            .iter()
            .map(|&(i, j, _)| (i as u32, j as u32)),
    );
    let mut comps_sizes: Vec<usize> = tarjan_scc(&graph).iter().map(|comp| comp.len()).collect();
    comps_sizes.sort();
    comps_sizes.reverse();
    comps_sizes[0..3].iter().product()
}

pub fn multiply_final_x_coords(inputs: Vec<&str>) -> u64 {
    let mut points = Vec::<Point3d>::new();
    let mut distances = Vec::<(usize, usize, u64)>::new();
    for input in inputs {
        points.push(Point3d::build(
            input.split(",").map(|s| s.parse().unwrap()).collect(),
        ));
    }
    for (i, pi) in points.iter().enumerate() {
        for (j, pj) in points[0..i].iter().enumerate() {
            distances.push((i, j, pi.distance_squared(pj)));
        }
    }
    distances.sort_by_key(|&x| x.2);
    let mut graph = UnGraph::<usize, ()>::new_undirected();
    let node_idxs: Vec<NodeIndex> = points
        .iter()
        .enumerate()
        .map(|(i, _)| graph.add_node(i))
        .collect();
    for (i, j, _) in distances {
        graph.add_edge(node_idxs[i], node_idxs[j], ());
        if tarjan_scc(&graph).len() == 1 {
            return points[i].x * points[j].x;
        }
    }
    0
}

struct Point3d {
    x: u64,
    y: u64,
    z: u64,
}

impl Point3d {
    fn build(coords: Vec<u64>) -> Point3d {
        Point3d {
            x: coords[0],
            y: coords[1],
            z: coords[2],
        }
    }

    fn distance_squared(&self, other: &Point3d) -> u64 {
        let mut d2 = 0;
        for pair in [(self.x, other.x), (self.y, other.y), (self.z, other.z)] {
            if pair.0 > pair.1 {
                d2 += (pair.0 - pair.1).pow(2);
            } else {
                d2 += (pair.1 - pair.0).pow(2);
            }
        }
        d2
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn multiply_circuit_sizes_example() {
        let inputs = vec![
            "162,817,812",
            "57,618,57",
            "906,360,560",
            "592,479,940",
            "352,342,300",
            "466,668,158",
            "542,29,236",
            "431,825,988",
            "739,650,466",
            "52,470,668",
            "216,146,977",
            "819,987,18",
            "117,168,530",
            "805,96,715",
            "346,949,466",
            "970,615,88",
            "941,993,340",
            "862,61,35",
            "984,92,344",
            "425,690,689",
        ];
        assert_eq!(multiply_circuit_sizes(inputs, 10), 40);
    }

    #[test]
    fn multiply_final_x_coords_example() {
        let inputs = vec![
            "162,817,812",
            "57,618,57",
            "906,360,560",
            "592,479,940",
            "352,342,300",
            "466,668,158",
            "542,29,236",
            "431,825,988",
            "739,650,466",
            "52,470,668",
            "216,146,977",
            "819,987,18",
            "117,168,530",
            "805,96,715",
            "346,949,466",
            "970,615,88",
            "941,993,340",
            "862,61,35",
            "984,92,344",
            "425,690,689",
        ];
        assert_eq!(multiply_final_x_coords(inputs), 25272);
    }
}
