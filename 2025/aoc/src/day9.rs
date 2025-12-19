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
    let red_tiles: Vec<Point2d> = inputs
        .iter()
        .map(|s| Point2d::from_str(s).unwrap())
        .collect();
    let (green_rows, green_cols) = build_green_rows_cols(&red_tiles);
    let interior_point = find_interior_point(&red_tiles, &green_cols);

    let mut green_blocks = vec![(interior_point.clone(), interior_point.clone())];
    let mut expansion_direction = Direction::Up;
    // TODO: Implement green tile fill

    let mut largest = 0;
    let mut current;
    for (i, pi) in red_tiles.iter().enumerate() {
        for pj in &red_tiles[0..i] {
            current = pi.area(pj);
            if current > largest {
                largest = current
            }
        }
    }
    largest
}

fn build_green_rows_cols(
    red_tiles: &Vec<Point2d>,
) -> (Vec<(Point2d, Point2d)>, Vec<(Point2d, Point2d)>) {
    let num_red_tiles = red_tiles.len();
    let mut green_rows = Vec::new();
    let mut green_cols = Vec::new();
    let mut j;
    let mut red_j;
    for (i, red_i) in red_tiles.iter().enumerate() {
        j = (i + 1) % num_red_tiles;
        red_j = &red_tiles[j];
        if red_i.x == red_j.x {
            if red_i.y < red_j.y {
                green_cols.push((
                    Point2d {
                        x: red_i.x,
                        y: red_i.y + 1,
                    },
                    Point2d {
                        x: red_i.x,
                        y: red_j.y - 1,
                    },
                ));
            } else if red_i.y > red_j.y {
                green_cols.push((
                    Point2d {
                        x: red_i.x,
                        y: red_j.y + 1,
                    },
                    Point2d {
                        x: red_i.x,
                        y: red_i.y - 1,
                    },
                ));
            } else {
                panic!("Invalid red tile pair");
            }
        } else if red_i.y == red_j.y {
            if red_i.x < red_j.x {
                green_rows.push((
                    Point2d {
                        x: red_i.x + 1,
                        y: red_i.y,
                    },
                    Point2d {
                        x: red_j.x - 1,
                        y: red_i.y,
                    },
                ));
            } else if red_i.x > red_j.x {
                green_cols.push((
                    Point2d {
                        x: red_j.x + 1,
                        y: red_i.y,
                    },
                    Point2d {
                        x: red_i.x - 1,
                        y: red_i.y,
                    },
                ));
            } else {
                panic!("Invalid red tile pair");
            }
        } else {
            panic!("Invalid red tile pair");
        }
    }
    (green_rows, green_cols)
}

fn find_interior_point(red_tiles: &Vec<Point2d>, green_cols: &Vec<(Point2d, Point2d)>) -> Point2d {
    let mut max_x = 0;
    let mut max_y = 0;
    for point in red_tiles.iter() {
        if point.x > max_x {
            max_x = point.x
        }
        if point.y > max_y {
            max_y = point.y
        }
    }
    let mut rand_x: i64;
    let mut rand_y: i64;
    let interior_point;
    loop {
        rand_x = rand::random_range(0..=max_x);
        rand_y = rand::random_range(0..=max_y);
        if green_cols
            .iter()
            .filter(|(pi, pj)| rand_x < pi.x && rand_y >= pi.y && rand_y <= pj.y)
            .count()
            % 2
            == 1
        {
            interior_point = Point2d {
                x: rand_x,
                y: rand_y,
            };
            break;
        }
    }
    interior_point
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

enum Direction {
    Up,
    Right,
    Down,
    Left,
}

impl Direction {
    fn next(&self) -> Self {
        match self {
            Self::Up => Self::Right,
            Self::Right => Self::Down,
            Self::Down => Self::Left,
            Self::Left => Self::Up,
        }
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
