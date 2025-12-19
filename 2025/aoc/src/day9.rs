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
