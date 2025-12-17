use std::vec;

#[derive(Clone)]
pub struct Block {
    start: (i64, i64),
    end: (i64, i64),
    value: u8,
}

impl Block {
    pub fn new(start: (i64, i64), end: (i64, i64), value: u8) -> Block {
        let (start, end) = if start <= end {
            (start, end)
        } else {
            (end, start)
        };
        Block { start, end, value }
    }

    fn subtract(&self, other: &Block) -> Vec<Block> {
        match self.intersect(other) {
            None => vec![self.clone()],
            Some(intersection) => {
                let value = self.value;
                if self.start == intersection.start && self.end == intersection.end {
                    Vec::new()
                } else if self.start == intersection.start {
                    let end = self.end;
                    if self.end.0 == intersection.end.0 {
                        // XXXX
                        // XXXX
                        // ....
                        // ....
                        let start = (self.start.0, end.1 - intersection.end.1);
                        vec![Block { start, end, value }]
                    } else if self.end.1 == intersection.end.1 {
                        // XX..
                        // XX..
                        // XX..
                        // XX..
                        let start = (end.0 - intersection.end.0, self.start.1);
                        vec![Block { start, end, value }]
                    } else {
                        // XX..
                        // XX..
                        // ....
                        // ....
                        vec![
                            Block {
                                start: (end.0 - intersection.end.0, self.start.1),
                                end: (end.0, end.1 - intersection.end.1),
                                value,
                            },
                            Block {
                                start: (self.start.0, end.1 - intersection.end.1),
                                end,
                                value,
                            },
                        ]
                    }
                } else if self.end == intersection.end {
                    let start = self.start;
                    if self.start.0 == intersection.start.0 {
                        // ....
                        // ....
                        // XXXX
                        // XXXX
                        let end = (self.end.0, self.end.1 - intersection.start.1);
                        vec![Block { start, end, value }]
                    } else if self.start.1 == intersection.start.1 {
                        // ..XX
                        // ..XX
                        // ..XX
                        // ..XX
                        let end = (intersection.end.0 - self.end.0, self.end.1);
                        vec![Block { start, end, value }]
                    } else {
                        // ....
                        // ....
                        // ..XX
                        // ..XX
                        vec![Block {
                            start,
                            end: (self.end.0, intersection.end.1 - self.end.1),
                            value,
                        }]
                    }
                } else {
                    // FIXME: add real else clause
                    Vec::new()
                }
            }
        }
    }

    fn intersect(&self, other: &Block) -> Option<Block> {
        None
    }
}

pub struct Image {
    blocks: Vec<Block>,
}

impl Image {
    pub fn new() -> Image {
        let blocks = Vec::new();
        Image { blocks }
    }

    pub fn add(&mut self, block: Block) {
        self.blocks = self
            .blocks
            .iter()
            .flat_map(|b| b.subtract(&block))
            .collect();
        self.blocks.push(block);
    }
}
