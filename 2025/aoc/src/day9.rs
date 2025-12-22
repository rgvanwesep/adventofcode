use std::cmp;

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

    let mut to_expand = vec![(interior_point.clone(), interior_point.clone())];
    let mut block;
    while !to_expand.is_empty() {
        block = to_expand.pop().unwrap();
        block.0.y = expand_up(&block, &red_tiles, &green_rows).unwrap_or(block.0.y);
        block.1.y = expand_down(&block, &red_tiles, &green_rows).unwrap_or(block.1.y);
        block.0.x = expand_left(&block, &red_tiles, &green_cols).unwrap_or(block.0.x);
        block.1.x = expand_right(&block, &red_tiles, &green_cols).unwrap_or(block.1.x);
    }

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

fn expand_up(
    block: &(Point2d, Point2d),
    red_tiles: &Vec<Point2d>,
    green_rows: &Vec<(Point2d, Point2d)>,
) -> Option<i64> {
    let closest_intersecting_red_tile = red_tiles
        .iter()
        .filter(|point| point.y < block.0.y && point.x == block.0.x)
        .max_by_key(|point| point.y);
    let closest_intersecting_green_row = green_rows
        .iter()
        .filter(|(start, end)| start.y < block.0.y && start.x <= block.0.x && end.x >= block.0.x)
        .max_by_key(|(start, _)| start.y);
    match (
        closest_intersecting_red_tile,
        closest_intersecting_green_row,
    ) {
        (Some(red_tile), Some(green_row)) => Some(cmp::max(red_tile.y, green_row.0.y)),
        (Some(red_tile), None) => Some(red_tile.y),
        (None, Some(green_row)) => Some(green_row.0.y),
        (None, None) => None,
    }
}

fn expand_down(
    block: &(Point2d, Point2d),
    red_tiles: &Vec<Point2d>,
    green_rows: &Vec<(Point2d, Point2d)>,
) -> Option<i64> {
    let closest_intersecting_red_tile = red_tiles
        .iter()
        .filter(|point| point.y > block.1.y && point.x == block.1.x)
        .min_by_key(|point| point.y);
    let closest_intersecting_green_row = green_rows
        .iter()
        .filter(|(start, end)| start.y > block.1.y && start.x <= block.1.x && end.x >= block.1.x)
        .min_by_key(|(start, _)| start.y);
    match (
        closest_intersecting_red_tile,
        closest_intersecting_green_row,
    ) {
        (Some(red_tile), Some(green_row)) => Some(cmp::min(red_tile.y, green_row.0.y)),
        (Some(red_tile), None) => Some(red_tile.y),
        (None, Some(green_row)) => Some(green_row.0.y),
        (None, None) => None,
    }
}

fn expand_left(
    block: &(Point2d, Point2d),
    red_tiles: &Vec<Point2d>,
    green_cols: &Vec<(Point2d, Point2d)>,
) -> Option<i64> {
    let closest_intersecting_red_tile = red_tiles
        .iter()
        .filter(|point| point.x < block.0.x && point.y >= block.0.y && point.y <= block.1.y)
        .max_by_key(|point| point.x);
    let closest_intersecting_green_col = green_cols
        .iter()
        .filter(|(start, end)| start.x < block.0.x && start.y <= block.1.y && end.y >= block.0.y)
        .max_by_key(|(start, _)| start.x);
    match (
        closest_intersecting_red_tile,
        closest_intersecting_green_col,
    ) {
        (Some(red_tile), Some(green_col)) => Some(cmp::max(red_tile.x, green_col.0.x)),
        (Some(red_tile), None) => Some(red_tile.x),
        (None, Some(green_col)) => Some(green_col.0.x),
        (None, None) => None,
    }
}

fn expand_right(
    block: &(Point2d, Point2d),
    red_tiles: &Vec<Point2d>,
    green_cols: &Vec<(Point2d, Point2d)>,
) -> Option<i64> {
    let closest_intersecting_red_tile = red_tiles
        .iter()
        .filter(|point| point.x > block.1.x && point.y >= block.0.y && point.y <= block.1.y)
        .min_by_key(|point| point.x);
    let closest_intersecting_green_col = green_cols
        .iter()
        .filter(|(start, end)| start.x > block.0.x && start.y <= block.1.y && end.y >= block.0.y)
        .min_by_key(|(start, _)| start.x);
    match (
        closest_intersecting_red_tile,
        closest_intersecting_green_col,
    ) {
        (Some(red_tile), Some(green_col)) => Some(cmp::min(red_tile.x, green_col.0.x)),
        (Some(red_tile), None) => Some(red_tile.x),
        (None, Some(green_col)) => Some(green_col.0.x),
        (None, None) => None,
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
