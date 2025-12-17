use std::{collections::HashMap, vec};

use petgraph::graph::UnGraph;

const NEUTRAL: u8 = b'.';
const RED: u8 = b'#';
const GREEN: u8 = b'X';

pub fn largest_area(inputs: Vec<&str>) -> i64 {
    let points: Vec<Point2d> = inputs
        .iter()
        .map(|s| Point2d::from_str(s).unwrap())
        .collect();
    let mut largest = 0;
    let mut current;
    for (i, pi) in points.iter().enumerate() {
        for pj in &points[0..i] {
            current = pi.area(pj);
            if current > largest {
                largest = current
            }
        }
    }
    largest
}

pub fn largest_area_green_red(inputs: Vec<&str>) -> i64 {
    let points: Vec<Point2d> = inputs
        .iter()
        .map(|s| Point2d::from_str(s).unwrap())
        .collect();
    let num_points = points.len();

    let mut max_x = 0;
    let mut max_y = 0;
    for point in points.iter() {
        if point.x > max_x {
            max_x = point.x
        }
        if point.y > max_y {
            max_y = point.y
        }
    }

    let mut floor = Floor::new(max_x + 1, max_y + 1);
    for point in points.iter() {
        floor.add_red_tile(point)
    }

    let mut j;
    let mut point_j;
    for (i, point_i) in points.iter().enumerate() {
        j = (i + 1) % num_points;
        point_j = &points[j];
        floor.color_green_tiles((point_i, point_j));
    }

    let mut largest = 0;
    let mut current;
    for (i, pi) in points.iter().enumerate() {
        for pj in &points[0..i] {
            current = pi.area(pj);
            if current > largest {
                largest = current
            }
        }
    }
    largest
}

struct Floor {
    blocks: Vec<Block>,
}

impl Floor {
    fn new(max_x: i64, max_y: i64) -> Floor {
        let corners = vec![
            Point2d { x: 0, y: 0 },
            Point2d { x: max_x, y: 0 },
            Point2d { x: 0, y: max_y },
            Point2d { x: max_x, y: max_y },
        ];
        let value = NEUTRAL;
        let blocks = vec![Block { corners, value }];
        Floor { blocks }
    }

    fn add_red_tile(&mut self, point: &Point2d) {
        self.blocks = self
            .blocks
            .iter()
            .flat_map(|block| block.subtract(point))
            .collect();
        let corners = vec![point.clone()];
        let value = RED;
        self.blocks.push(Block { corners, value });
    }

    fn color_green_tiles(&mut self, pair: (&Point2d, &Point2d)) {
        let orientation = if pair.0.x == pair.1.x {
            Orientation::Vertical
        } else {
            Orientation::Horizontal
        };
        let (start, end) = match orientation {
            Orientation::Horizontal => {
                if pair.0.x <= pair.1.x {
                    (pair.0, pair.1)
                } else {
                    (pair.1, pair.0)
                }
            }
            Orientation::Vertical => {
                if pair.0.y <= pair.1.y {
                    (pair.0, pair.1)
                } else {
                    (pair.1, pair.0)
                }
            }
        };
        for block in self.blocks.iter_mut().filter(|block| {
            if block.corners.len() == 1 {
                let point = &block.corners[0];
                match orientation {
                    Orientation::Horizontal => point.x == start.x + 1 && point.x == end.x - 1,
                    Orientation::Vertical => point.y == start.y + 1 && point.y == end.y - 1,
                }
            } else if block.corners.len() == 2 {
                let point1 = &block.corners[0];
                let point2 = &block.corners[1];
                match orientation {
                    Orientation::Horizontal => point1.x == start.x + 1 && point2.x == end.x - 1,
                    Orientation::Vertical => point1.y == start.y + 1 && point2.y == end.y - 1,
                }
            } else {
                false
            }
        }) {
            block.value = GREEN
        }
    }
}

enum Orientation {
    Horizontal,
    Vertical,
}

struct Block {
    corners: Vec<Point2d>,
    value: u8,
}

impl Block {
    fn subtract(&self, point: &Point2d) -> Vec<Block> {
        if self.corners.iter().any(|corner| corner == point) {
            Vec::new()
        } else if self.corners.iter().any(|corner| corner.x == point.x) {
            Vec::new()
        } else if self.corners.iter().any(|corner| corner.y == point.y) {
            Vec::new()
        } else {
            Vec::new()
        }
    }
}

#[derive(Debug, Clone, Hash, PartialEq, Eq)]
struct Point2d {
    x: i64,
    y: i64,
}

impl Point2d {
    fn from_str(s: &str) -> Result<Point2d, &str> {
        let coords: Vec<Result<i64, _>> = s.split(",").map(|s| s.parse()).collect();
        if coords.len() != 2 {
            return Err("There must be two coordinates");
        }
        let &x = match &coords[0] {
            Ok(n) => n,
            Err(_) => {
                return Err("Parse error");
            }
        };
        let &y = match &coords[1] {
            Ok(n) => n,
            Err(_) => {
                return Err("Parse error");
            }
        };
        Ok(Point2d { x, y })
    }

    fn area(&self, other: &Point2d) -> i64 {
        let width = if self.x > other.x {
            self.x - other.x + 1
        } else {
            other.x - self.x + 1
        };
        let height = if self.y > other.y {
            self.y - other.y + 1
        } else {
            other.y - self.y + 1
        };
        width * height
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn largest_area_example() {
        let inputs = vec!["7,1", "11,1", "11,7", "9,7", "9,5", "2,5", "2,3", "7,3"];
        assert_eq!(largest_area(inputs), 50)
    }

    #[test]
    fn largest_area_green_red_example() {
        let inputs = vec!["7,1", "11,1", "11,7", "9,7", "9,5", "2,5", "2,3", "7,3"];
        assert_eq!(largest_area_green_red(inputs), 24)
    }
}
