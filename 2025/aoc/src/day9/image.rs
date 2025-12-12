struct BlockSpec {
    start: (i64, i64),
    end: (i64, i64),
    value: u8,
}

enum Block {
    Empty,
    Point(BlockSpec),
    VerticalLine(BlockSpec),
    HorizontalLine(BlockSpec),
    Rectangle(BlockSpec),
}

impl Block {
    fn new(start: (i64, i64), end: (i64, i64), value: u8) -> Block {
        if start == end {
            Self::Point(BlockSpec { start, end, value })
        } else if start.0 == end.0 {
            if start.1 < end.1 {
                Self::VerticalLine(BlockSpec { start, end, value })
            } else {
                Self::VerticalLine(BlockSpec {
                    start: end,
                    end: start,
                    value,
                })
            }
        } else if start.1 == end.1 {
            if start.0 < end.0 {
                Self::HorizontalLine(BlockSpec { start, end, value })
            } else {
                Self::HorizontalLine(BlockSpec {
                    start: end,
                    end: start,
                    value,
                })
            }
        } else {
            if start.0 < end.0 && start.1 < end.1 {
                Self::Rectangle(BlockSpec { start, end, value })
            } else if end.0 < start.0 && end.1 < start.1 {
                Self::Rectangle(BlockSpec {
                    start: end,
                    end: start,
                    value,
                })
            } else if start.0 < end.0 && end.1 < start.1 {
                Self::Rectangle(BlockSpec {
                    start: (start.0, end.1),
                    end: (end.0, start.1),
                    value,
                })
            } else {
                Self::Rectangle(BlockSpec {
                    start: (end.0, start.1),
                    end: (start.0, end.1),
                    value,
                })
            }
        }
    }

    fn intersect(&self, other: &Self, value: u8) -> Self {
        match (self, other) {
            (Self::Empty, _) => Self::Empty,
            (_, Self::Empty) => Self::Empty,
            (Self::Point(self_spec), Self::Point(other_spec)) => {
                if self_spec.start == other_spec.start {
                    Self::Point(BlockSpec {
                        start: self_spec.start,
                        end: self_spec.end,
                        value,
                    })
                } else {
                    Self::Empty
                }
            }
            (Self::Point(self_spec), Self::VerticalLine(other_spec)) => {
                if self_spec.start.0 == other_spec.start.0
                    && self_spec.start.1 >= other_spec.start.1
                    && self_spec.start.1 <= other_spec.end.1
                {
                    Self::Point(BlockSpec {
                        start: self_spec.start,
                        end: self_spec.end,
                        value,
                    })
                } else {
                    Self::Empty
                }
            }
            (Self::Point(self_spec), Self::HorizontalLine(other_spec)) => {
                if self_spec.start.1 == other_spec.start.1
                    && self_spec.start.0 >= other_spec.start.0
                    && self_spec.start.0 <= other_spec.end.0
                {
                    Self::Point(BlockSpec {
                        start: self_spec.start,
                        end: self_spec.end,
                        value,
                    })
                } else {
                    Self::Empty
                }
            }
            (Self::Point(self_spec), Self::Rectangle(other_spec)) => {
                if self_spec.start.0 >= other_spec.start.0
                    && self_spec.start.0 <= other_spec.end.0
                    && self_spec.start.1 >= other_spec.start.1
                    && self_spec.start.1 <= other_spec.end.1
                {
                    Self::Point(BlockSpec {
                        start: self_spec.start,
                        end: self_spec.end,
                        value,
                    })
                } else {
                    Self::Empty
                }
            }
            (Self::VerticalLine(self_spec), Self::Point(other_spec)) => {
                if other_spec.start.0 == self_spec.start.0
                    && other_spec.start.1 >= self_spec.start.1
                    && other_spec.start.1 <= self_spec.end.1
                {
                    Self::Point(BlockSpec {
                        start: other_spec.start,
                        end: other_spec.end,
                        value,
                    })
                } else {
                    Self::Empty
                }
            }
            (Self::VerticalLine(self_spec), Self::VerticalLine(other_spec)) => {
                if self_spec.start.0 == other_spec.start.0 {
                    let start = if self_spec.start.1 >= other_spec.start.1 {
                        self_spec.start.1
                    } else {
                        other_spec.start.1
                    };
                    let end = if self_spec.end.1 <= other_spec.end.1 {
                        self_spec.end.1
                    } else {
                        other_spec.end.1
                    };
                    if start < end {
                        Self::VerticalLine(BlockSpec {
                            start: (self_spec.start.0, start),
                            end: (self_spec.start.0, end),
                            value,
                        })
                    } else if start == end {
                        Self::Point(BlockSpec {
                            start: (self_spec.start.0, start),
                            end: (self_spec.start.0, end),
                            value,
                        })
                    } else {
                        Self::Empty
                    }
                } else {
                    Self::Empty
                }
            }
            (Self::VerticalLine(self_spec), Self::HorizontalLine(other_spec)) => {
                if self_spec.start.0 >= other_spec.start.0
                    && self_spec.start.0 <= other_spec.end.0
                    && other_spec.start.1 >= self_spec.start.1
                    && other_spec.start.1 >= self_spec.end.1
                {
                    let start = (self_spec.start.0, other_spec.start.1);
                    Self::Point(BlockSpec {
                        start: start,
                        end: start,
                        value,
                    })
                } else {
                    Self::Empty
                }
            }
            (Self::VerticalLine(self_spec), Self::Rectangle(other_spec)) => {
                if self_spec.start.0 >= other_spec.start.0 && self_spec.start.0 <= other_spec.end.0
                {
                    let start = if self_spec.start.1 >= other_spec.start.1 {
                        self_spec.start.1
                    } else {
                        other_spec.start.1
                    };
                    let end = if self_spec.end.1 <= other_spec.end.1 {
                        self_spec.end.1
                    } else {
                        other_spec.end.1
                    };
                    if start < end {
                        Self::VerticalLine(BlockSpec {
                            start: (self_spec.start.0, start),
                            end: (self_spec.start.0, end),
                            value,
                        })
                    } else if start == end {
                        Self::Point(BlockSpec {
                            start: (self_spec.start.0, start),
                            end: (self_spec.start.0, end),
                            value,
                        })
                    } else {
                        Self::Empty
                    }
                } else {
                    Self::Empty
                }
            }
            (Self::HorizontalLine(self_spec), Self::Point(other_spec)) => {
                if other_spec.start.1 == self_spec.start.1
                    && other_spec.start.0 >= self_spec.start.0
                    && other_spec.start.0 <= self_spec.end.0
                {
                    Self::Point(BlockSpec {
                        start: other_spec.start,
                        end: other_spec.end,
                        value,
                    })
                } else {
                    Self::Empty
                }
            }
            (Self::HorizontalLine(self_spec), Self::VerticalLine(other_spec)) => {
                if other_spec.start.0 >= self_spec.start.0
                    && other_spec.start.0 <= self_spec.end.0
                    && self_spec.start.1 >= other_spec.start.1
                    && self_spec.start.1 >= other_spec.end.1
                {
                    let start = (other_spec.start.0, self_spec.start.1);
                    Self::Point(BlockSpec {
                        start: start,
                        end: start,
                        value,
                    })
                } else {
                    Self::Empty
                }
            }
            (Self::HorizontalLine(self_spec), Self::HorizontalLine(other_spec)) => {
                if self_spec.start.1 == other_spec.start.1 {
                    let start = if self_spec.start.0 >= other_spec.start.0 {
                        self_spec.start.0
                    } else {
                        other_spec.start.0
                    };
                    let end = if self_spec.end.0 <= other_spec.end.0 {
                        self_spec.end.0
                    } else {
                        other_spec.end.0
                    };
                    if start < end {
                        Self::HorizontalLine(BlockSpec {
                            start: (self_spec.start.1, start),
                            end: (self_spec.start.1, end),
                            value,
                        })
                    } else if start == end {
                        Self::Point(BlockSpec {
                            start: (self_spec.start.1, start),
                            end: (self_spec.start.1, end),
                            value,
                        })
                    } else {
                        Self::Empty
                    }
                } else {
                    Self::Empty
                }
            }
            (Self::HorizontalLine(self_spec), Self::Rectangle(other_spec)) => {
                if self_spec.start.1 >= other_spec.start.1 && self_spec.start.1 <= other_spec.end.1
                {
                    let start = if self_spec.start.0 >= other_spec.start.0 {
                        self_spec.start.0
                    } else {
                        other_spec.start.0
                    };
                    let end = if self_spec.end.0 <= other_spec.end.0 {
                        self_spec.end.0
                    } else {
                        other_spec.end.0
                    };
                    if start < end {
                        Self::HorizontalLine(BlockSpec {
                            start: (self_spec.start.0, start),
                            end: (self_spec.start.0, end),
                            value,
                        })
                    } else if start == end {
                        Self::Point(BlockSpec {
                            start: (self_spec.start.0, start),
                            end: (self_spec.start.0, end),
                            value,
                        })
                    } else {
                        Self::Empty
                    }
                } else {
                    Self::Empty
                }
            }
            (Self::Rectangle(self_spec), Self::Point(other_spec)) => {
                if other_spec.start.0 >= self_spec.start.0
                    && other_spec.start.0 <= self_spec.end.0
                    && other_spec.start.1 >= self_spec.start.1
                    && other_spec.start.1 <= self_spec.end.1
                {
                    Self::Point(BlockSpec {
                        start: other_spec.start,
                        end: other_spec.end,
                        value,
                    })
                } else {
                    Self::Empty
                }
            }
            (Self::Rectangle(self_spec), Self::VerticalLine(other_spec)) => {
                if other_spec.start.0 >= self_spec.start.0 && other_spec.start.0 <= self_spec.end.0
                {
                    let start = if other_spec.start.1 >= self_spec.start.1 {
                        other_spec.start.1
                    } else {
                        self_spec.start.1
                    };
                    let end = if other_spec.end.1 <= self_spec.end.1 {
                        other_spec.end.1
                    } else {
                        self_spec.end.1
                    };
                    if start < end {
                        Self::VerticalLine(BlockSpec {
                            start: (other_spec.start.0, start),
                            end: (other_spec.start.0, end),
                            value,
                        })
                    } else if start == end {
                        Self::Point(BlockSpec {
                            start: (other_spec.start.0, start),
                            end: (other_spec.start.0, end),
                            value,
                        })
                    } else {
                        Self::Empty
                    }
                } else {
                    Self::Empty
                }
            }
            (Self::Rectangle(self_spec), Self::HorizontalLine(other_spec)) => {
                if other_spec.start.1 >= self_spec.start.1 && other_spec.start.1 <= self_spec.end.1
                {
                    let start = if other_spec.start.0 >= self_spec.start.0 {
                        other_spec.start.0
                    } else {
                        self_spec.start.0
                    };
                    let end = if other_spec.end.0 <= self_spec.end.0 {
                        other_spec.end.0
                    } else {
                        self_spec.end.0
                    };
                    if start < end {
                        Self::HorizontalLine(BlockSpec {
                            start: (other_spec.start.0, start),
                            end: (other_spec.start.0, end),
                            value,
                        })
                    } else if start == end {
                        Self::Point(BlockSpec {
                            start: (other_spec.start.0, start),
                            end: (other_spec.start.0, end),
                            value,
                        })
                    } else {
                        Self::Empty
                    }
                } else {
                    Self::Empty
                }
            }
            (Self::Rectangle(self_spec), Self::Rectangle(other_spec)) => {
                let start = (
                    if self_spec.start.0 >= other_spec.start.0 {
                        self_spec.start.0
                    } else {
                        other_spec.start.0
                    },
                    if self_spec.start.1 >= other_spec.start.1 {
                        self_spec.start.1
                    } else {
                        other_spec.start.1
                    },
                );
                let end = (
                    if self_spec.end.0 >= other_spec.end.0 {
                        self_spec.end.0
                    } else {
                        other_spec.end.0
                    },
                    if self_spec.end.1 >= other_spec.end.1 {
                        self_spec.end.1
                    } else {
                        other_spec.end.1
                    },
                );
                if start <= end {
                    Block::new(start, end, value)
                } else {
                    Block::Empty
                }
            }
        }
    }
}
