#[derive(Debug, Clone, PartialEq)]
enum Quadrent {
    I,
    II,
    III,
    IV,
}

impl Quadrent {
    fn from_point(p: (i64, i64)) -> Quadrent {
        if p.0 >= 0 {
            if p.1 >= 0 { Quadrent::I } else { Quadrent::IV }
        } else {
            if p.1 >= 0 {
                Quadrent::II
            } else {
                Quadrent::III
            }
        }
    }

    fn next(&self) -> Quadrent {
        match self {
            Quadrent::I => Quadrent::II,
            Quadrent::II => Quadrent::III,
            Quadrent::III => Quadrent::IV,
            Quadrent::IV => Quadrent::I,
        }
    }

    fn prev(&self) -> Quadrent {
        match self {
            Quadrent::I => Quadrent::IV,
            Quadrent::II => Quadrent::I,
            Quadrent::III => Quadrent::II,
            Quadrent::IV => Quadrent::III,
        }
    }
}

pub struct Polygon {
    vertices: Vec<(i64, i64)>,
    sides: Vec<((i64, i64), (i64, i64))>,
}

impl Polygon {
    pub fn from_strings(inputs: Vec<&str>) -> Polygon {
        let vertices: Vec<(i64, i64)> = inputs
            .iter()
            .map(|s| {
                let coords: Vec<i64> = s.split(",").map(|n| n.parse().unwrap()).collect();
                (coords[0], coords[1])
            })
            .collect();
        let mut sides: Vec<((i64, i64), (i64, i64))> = Vec::new();
        for i in 0..vertices.len() {
            sides.push((vertices[i], vertices[(i + 1) % vertices.len()]));
        }
        Polygon { vertices, sides }
    }

    pub fn vertices(&self) -> Vec<(i64, i64)> {
        self.vertices.to_vec()
    }

    pub fn contains(&self, p: (i64, i64)) -> bool {
        if self.on_boundary(p) {
            return true;
        }
        let (px, py) = p;
        let qs: Vec<Quadrent> = self
            .vertices
            .iter()
            .map(|&(vx, vy)| Quadrent::from_point((vx - px, vy - py)))
            .collect();
        let mut current;
        let mut stack = vec![qs[0].clone()];
        for q in &qs {
            current = stack.last().unwrap();
            if *q == current.next() {
                stack.push(q.clone());
            } else if *q == current.prev() {
                stack.pop();
                if stack.is_empty() {
                    stack.push(q.clone());
                }
            }
        }
        stack.len() >= 4
    }

    fn on_boundary(&self, p: (i64, i64)) -> bool {
        for &(start, end) in self.sides.iter() {
            if (start.0 == end.0 && p.0 == start.0)
                && ((start.1 < end.1 && p.1 >= start.1 && p.1 <= end.1)
                    || (start.1 > end.1 && p.1 <= start.1 && p.1 >= end.1))
            {
                return true;
            } else if (start.1 == end.1 && p.1 == start.1)
                && ((start.0 < end.0 && p.0 >= start.0 && p.0 <= end.0)
                    || (start.0 > end.0 && p.0 <= start.0 && p.0 >= end.0))
            {
                return true;
            }
        }
        false
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn from_strings_example() {
        let inputs = vec!["7,1", "11,1", "11,7", "9,7", "9,5", "2,5", "2,3", "7,3"];
        let p = Polygon::from_strings(inputs);
        assert_eq!(
            p.vertices,
            vec![
                (7, 1),
                (11, 1),
                (11, 7),
                (9, 7),
                (9, 5),
                (2, 5),
                (2, 3),
                (7, 3)
            ]
        );
        assert_eq!(
            p.sides,
            vec![
                ((7, 1), (11, 1)),
                ((11, 1), (11, 7)),
                ((11, 7), (9, 7)),
                ((9, 7), (9, 5)),
                ((9, 5), (2, 5)),
                ((2, 5), (2, 3)),
                ((2, 3), (7, 3)),
                ((7, 3), (7, 1)),
            ]
        )
    }

    #[test]
    fn on_boundary_example() {
        let inputs = vec!["7,1", "11,1", "11,7", "9,7", "9,5", "2,5", "2,3", "7,3"];
        let p = Polygon::from_strings(inputs);
        assert!(p.on_boundary((7, 1)));
        assert!(p.on_boundary((8, 1)));
        assert!(p.on_boundary((11, 1)));
        assert!(p.on_boundary((11, 2)));
        assert!(p.on_boundary((11, 7)));
        assert!(p.on_boundary((10, 7)));
        assert!(p.on_boundary((9, 7)));
        assert!(p.on_boundary((9, 6)));
        assert!(p.on_boundary((9, 5)));
        assert!(!p.on_boundary((12, 1)));
        assert!(!p.on_boundary((8, 2)));
    }

    #[test]
    fn contains_example() {
        let inputs = vec!["7,1", "11,1", "11,7", "9,7", "9,5", "2,5", "2,3", "7,3"];
        let p = Polygon::from_strings(inputs);
        assert!(p.contains((8, 2)));
        assert!(!p.contains((8, 0)));
    }
}
