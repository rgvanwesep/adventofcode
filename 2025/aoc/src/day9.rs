mod image;
mod polygon;

use polygon::Polygon;

pub fn largest_area(inputs: Vec<&str>) -> usize {
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
    let poly = Polygon::from_strings(inputs);
    let red_tiles = poly.vertices();
    let mut largest = 0;
    let mut m1;
    let mut m2;
    let mut area;
    for (i, &ri) in red_tiles.iter().enumerate() {
        for &rj in &red_tiles[0..i] {
            area = ((ri.0 - rj.0).abs() + 1) * ((ri.1 - rj.1).abs() + 1);
            if area > largest {
                (m1, m2) = missing_corners(ri, rj);
                if poly.contains(m1) && poly.contains(m2) {
                    largest = area
                }
            }
        }
    }
    largest
}

fn missing_corners(p1: (i64, i64), p2: (i64, i64)) -> ((i64, i64), (i64, i64)) {
    ((p1.0, p2.1), (p2.0, p1.1))
}

#[derive(Debug, Clone, Hash, PartialEq, Eq)]
struct Point2d {
    x: usize,
    y: usize,
}

impl Point2d {
    fn from_str(s: &str) -> Result<Point2d, &str> {
        let coords: Vec<Result<usize, _>> = s.split(",").map(|s| s.parse()).collect();
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

    fn area(&self, other: &Point2d) -> usize {
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
