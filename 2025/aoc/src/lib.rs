use regex::Regex;

#[derive(Debug, PartialEq)]
enum Direction {
    Left,
    Right,
}

impl Direction {
    fn build(input: &str) -> Result<Direction, &'static str> {
        match input {
            "L" => Ok(Direction::Left),
            "R" => Ok(Direction::Right),
            _ => Err("Invalid input for Direction"),
        }
    }
}

#[derive(Debug, PartialEq)]
struct Rotation {
    direction: Direction,
    distance: i32,
}

impl Rotation {
    fn build(input: &str) -> Result<Rotation, &'static str> {
        let re = Regex::new(r"(?<direction>L|R)(?<distance>[0-9]+)").unwrap();
        let Some(caps) = re.captures(input) else {
            return Err("Invalid input for Rotation");
        };
        let direction = Direction::build(&caps["direction"])?;
        let distance: i32 = String::from(&caps["distance"]).parse().unwrap();
        Ok(Rotation {
            direction,
            distance,
        })
    }
}

#[derive(Debug, PartialEq)]
struct Dial {
    position: i32,
    zero_landed_count: i32,
    zero_passed_count: i32,
}

impl Dial {
    fn new(start: i32) -> Dial {
        Dial {
            position: start,
            zero_landed_count: 0,
            zero_passed_count: 0,
        }
    }

    fn rotate(&mut self, r: Rotation) {
        match r {
            Rotation {
                direction: Direction::Left,
                distance,
            } => self.rotate_left(distance),
            Rotation {
                direction: Direction::Right,
                distance,
            } => self.rotate_right(distance),
        }
    }

    fn rotate_left(&mut self, distance: i32) {
        if self.position == 0 {
            if distance % 100 == 0 {
                self.position = 0;
                self.zero_landed_count += 1;
                self.zero_passed_count += distance / 100 - 1;
            } else {
                self.position = 100 - distance % 100;
                self.zero_passed_count += distance / 100;
            }
        } else if distance > self.position {
            let difference = distance - self.position;
            if difference % 100 == 0 {
                self.position = 0;
                self.zero_landed_count += 1;
                self.zero_passed_count += difference / 100;
            } else {
                self.position = 100 - difference % 100;
                self.zero_passed_count += difference / 100 + 1;
            }
        } else if distance == self.position {
            self.position = 0;
            self.zero_landed_count += 1;
        } else {
            self.position -= distance;
        }
    }

    fn rotate_right(&mut self, distance: i32) {
        if self.position == 0 {
            if distance % 100 == 0 {
                self.position = 0;
                self.zero_landed_count += 1;
                self.zero_passed_count += distance / 100 - 1;
            } else {
                self.position = distance % 100;
                self.zero_passed_count += distance / 100;
            }
        } else if distance > 100 - self.position {
            let difference = distance - 100 + self.position;
            if difference % 100 == 0 {
                self.position = 0;
                self.zero_landed_count += 1;
                self.zero_passed_count += difference / 100;
            } else {
                self.position = difference % 100;
                self.zero_passed_count += difference / 100 + 1;
            }
        } else if distance == 100 - self.position {
            self.position = 0;
            self.zero_landed_count += 1;
        } else {
            self.position += distance;
        }
    }
}

#[derive(Debug, PartialEq)]
pub struct RotationSeq {
    rotations: Vec<Rotation>,
}

impl RotationSeq {
    pub fn build(inputs: Vec<&str>) -> Result<RotationSeq, &'static str> {
        let mut rotations = Vec::new();
        let mut rotation;
        for item in inputs {
            rotation = Rotation::build(item)?;
            rotations.push(rotation)
        }
        Ok(RotationSeq { rotations })
    }

    pub fn count_zeros(&self) -> i32 {
        let mut count: i32 = 0;
        let mut position: i32 = 50;
        for rotation in &self.rotations {
            position = match rotation {
                Rotation {
                    direction: Direction::Left,
                    distance,
                } => (position - distance) % 100,
                Rotation {
                    direction: Direction::Right,
                    distance,
                } => (position + distance) % 100,
            };
            if position == 0 {
                count += 1;
            };
        }
        count
    }

    pub fn count_all_zeros(&self) -> i32 {
        let mut count: i32 = 0;
        let mut position: i32 = 50;
        let mut displacement: i32;
        for rotation in &self.rotations {
            match rotation {
                Rotation {
                    direction: Direction::Left,
                    distance,
                } => {
                    displacement = position - distance;
                    if position == 0 {
                        position = displacement % 100;
                        if position < 0 {
                            position += 100;
                        };
                        count += -displacement / 100
                    } else {
                        position = displacement % 100;
                        if position < 0 {
                            count += -displacement / 100 + 1;
                            position += 100;
                        } else if position == 0 {
                            count += 1;
                        };
                    };
                }
                Rotation {
                    direction: Direction::Right,
                    distance,
                } => {
                    displacement = position + distance;
                    position = displacement % 100;
                    count += displacement / 100;
                }
            };

            // println!(
            //     "rotation: {:?}, count: {}, position: {}",
            //     rotation, count, position
            // );
        }
        count
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn direction_ok() {
        assert_eq!(Direction::build("L").unwrap(), Direction::Left);
        assert_eq!(Direction::build("R").unwrap(), Direction::Right);
    }

    #[test]
    #[should_panic]
    fn direction_err() {
        Direction::build("O").unwrap();
    }

    #[test]
    fn rotation_ok() {
        assert_eq!(
            Rotation::build("L1").unwrap(),
            Rotation {
                direction: Direction::Left,
                distance: 1
            }
        );
        assert_eq!(
            Rotation::build("L19").unwrap(),
            Rotation {
                direction: Direction::Left,
                distance: 19
            }
        );
        assert_eq!(
            Rotation::build("L123").unwrap(),
            Rotation {
                direction: Direction::Left,
                distance: 123
            }
        );
        assert_eq!(
            Rotation::build("R2").unwrap(),
            Rotation {
                direction: Direction::Right,
                distance: 2
            }
        );
        assert_eq!(
            Rotation::build("R28").unwrap(),
            Rotation {
                direction: Direction::Right,
                distance: 28
            }
        );
    }

    #[test]
    #[should_panic]
    fn rotation_err() {
        Direction::build("O123").unwrap();
    }

    #[test]
    fn rotation_seq() {
        let input = vec![
            "L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82",
        ];
        assert_eq!(
            RotationSeq::build(input).unwrap(),
            RotationSeq {
                rotations: vec![
                    Rotation {
                        direction: Direction::Left,
                        distance: 68,
                    },
                    Rotation {
                        direction: Direction::Left,
                        distance: 30,
                    },
                    Rotation {
                        direction: Direction::Right,
                        distance: 48,
                    },
                    Rotation {
                        direction: Direction::Left,
                        distance: 5,
                    },
                    Rotation {
                        direction: Direction::Right,
                        distance: 60,
                    },
                    Rotation {
                        direction: Direction::Left,
                        distance: 55,
                    },
                    Rotation {
                        direction: Direction::Left,
                        distance: 1,
                    },
                    Rotation {
                        direction: Direction::Left,
                        distance: 99,
                    },
                    Rotation {
                        direction: Direction::Right,
                        distance: 14,
                    },
                    Rotation {
                        direction: Direction::Left,
                        distance: 82,
                    },
                ]
            }
        )
    }

    #[test]
    fn count_zeros() {
        let input = vec![
            "L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82",
        ];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_zeros(), 3);
    }

    #[test]
    fn count_all_zeros_one() {
        let input = vec!["L68"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 1);
    }

    #[test]
    fn count_all_zeros_two() {
        let input = vec!["L68", "L30"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 1);
    }

    #[test]
    fn count_all_zeros_three() {
        let input = vec!["L68", "L30", "R48"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 2);
    }

    #[test]
    fn count_all_zeros_four() {
        let input = vec!["L68", "L30", "R48", "L5"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 2);
    }

    #[test]
    fn count_all_zeros_five() {
        let input = vec!["L68", "L30", "R48", "L5", "R60"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 3);
    }

    #[test]
    fn count_all_zeros_six() {
        let input = vec!["L68", "L30", "R48", "L5", "R60", "L55"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 4);
    }

    #[test]
    fn count_all_zeros_seven() {
        let input = vec!["L68", "L30", "R48", "L5", "R60", "L55", "L1"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 4);
    }

    #[test]
    fn count_all_zeros_eight() {
        let input = vec!["L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 5);
    }

    #[test]
    fn count_all_zeros_nine() {
        let input = vec!["L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 5);
    }

    #[test]
    fn count_all_zeros_ten() {
        let input = vec![
            "L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82",
        ];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 6);
    }

    #[test]
    fn count_all_zeros_right_big() {
        let input = vec!["R1000"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 10);
    }

    #[test]
    fn count_all_zeros_right_big_start_zero() {
        let input = vec!["L50", "R1000"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 11);
    }

    #[test]
    fn count_all_zeros_left_big() {
        let input = vec!["L1000"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 10);
    }

    #[test]
    fn count_all_zeros_left_big_start_zero() {
        let input = vec!["R50", "L1000"];
        let rotation_seq = RotationSeq::build(input).unwrap();
        assert_eq!(rotation_seq.count_all_zeros(), 11);
    }

    #[test]
    fn rotate_left_from_zero() {
        let mut dial = Dial::new(0);
        dial.rotate_left(10);
        assert_eq!(dial, Dial{position: 90, zero_landed_count: 0, zero_passed_count: 0});

        let mut dial = Dial::new(0);
        dial.rotate_left(100);
        assert_eq!(dial, Dial{position: 0, zero_landed_count: 1, zero_passed_count: 0});

        let mut dial = Dial::new(0);
        dial.rotate_left(110);
        assert_eq!(dial, Dial{position: 90, zero_landed_count: 0, zero_passed_count: 1});

        let mut dial = Dial::new(0);
        dial.rotate_left(200);
        assert_eq!(dial, Dial{position: 0, zero_landed_count: 1, zero_passed_count: 1});

        let mut dial = Dial::new(0);
        dial.rotate_left(210);
        assert_eq!(dial, Dial{position: 90, zero_landed_count: 0, zero_passed_count: 2});
    }

    #[test]
    fn rotate_left_from_nonzero() {
        let mut dial = Dial::new(50);
        dial.rotate_left(10);
        assert_eq!(dial, Dial{position: 40, zero_landed_count: 0, zero_passed_count: 0});

        let mut dial = Dial::new(50);
        dial.rotate_left(50);
        assert_eq!(dial, Dial{position: 0, zero_landed_count: 1, zero_passed_count: 0});

        let mut dial = Dial::new(50);
        dial.rotate_left(60);
        assert_eq!(dial, Dial{position: 90, zero_landed_count: 0, zero_passed_count: 1});

        let mut dial = Dial::new(50);
        dial.rotate_left(150);
        assert_eq!(dial, Dial{position: 0, zero_landed_count: 1, zero_passed_count: 1});

        let mut dial = Dial::new(50);
        dial.rotate_left(160);
        assert_eq!(dial, Dial{position: 90, zero_landed_count: 0, zero_passed_count: 2});
    }

    #[test]
    fn rotate_right_from_zero() {
        let mut dial = Dial::new(0);
        dial.rotate_right(10);
        assert_eq!(dial, Dial{position: 10, zero_landed_count: 0, zero_passed_count: 0});

        let mut dial = Dial::new(0);
        dial.rotate_right(100);
        assert_eq!(dial, Dial{position: 0, zero_landed_count: 1, zero_passed_count: 0});

        let mut dial = Dial::new(0);
        dial.rotate_right(110);
        assert_eq!(dial, Dial{position: 10, zero_landed_count: 0, zero_passed_count: 1});

        let mut dial = Dial::new(0);
        dial.rotate_right(200);
        assert_eq!(dial, Dial{position: 0, zero_landed_count: 1, zero_passed_count: 1});

        let mut dial = Dial::new(0);
        dial.rotate_right(210);
        assert_eq!(dial, Dial{position: 10, zero_landed_count: 0, zero_passed_count: 2});
    }

    #[test]
    fn rotate_right_from_nonzero() {
        let mut dial = Dial::new(50);
        dial.rotate_right(10);
        assert_eq!(dial, Dial{position: 60, zero_landed_count: 0, zero_passed_count: 0});

        let mut dial = Dial::new(50);
        dial.rotate_right(50);
        assert_eq!(dial, Dial{position: 0, zero_landed_count: 1, zero_passed_count: 0});

        let mut dial = Dial::new(50);
        dial.rotate_right(60);
        assert_eq!(dial, Dial{position: 10, zero_landed_count: 0, zero_passed_count: 1});

        let mut dial = Dial::new(50);
        dial.rotate_right(150);
        assert_eq!(dial, Dial{position: 0, zero_landed_count: 1, zero_passed_count: 1});

        let mut dial = Dial::new(50);
        dial.rotate_right(160);
        assert_eq!(dial, Dial{position: 10, zero_landed_count: 0, zero_passed_count: 2});
    }
}
