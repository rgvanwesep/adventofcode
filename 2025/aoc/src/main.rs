use aoc::day1::RotationSeq;
use clap::Parser;
use std::error::Error;
use std::io;

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    #[arg(short, long)]
    day: u8,

    #[arg(short, long)]
    part: u8,
}

fn main() -> Result<(), Box<dyn Error>> {
    let args = Args::parse();

    let mut inputs = Vec::new();
    for line in io::stdin().lines() {
        inputs.push(line.unwrap())
    }

    match (args.day, args.part) {
        (1, 1) => {
            println!(
                "Result: {}",
                RotationSeq::build(inputs.iter().map(String::as_str).collect())?.count_zeros()
            );
            Ok(())
        }
        (1, 2) => {
            println!(
                "Result: {}",
                RotationSeq::build(inputs.iter().map(String::as_str).collect())?.count_all_zeros()
            );
            Ok(())
        }
        _ => Err(format!("No match for Day {}, Part {}", args.day, args.part).into()),
    }
}
