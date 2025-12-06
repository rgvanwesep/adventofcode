use aoc::day1::RotationSeq;
use aoc::{day2, day3, day4, day5};
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
        (2, 1) => {
            println!(
                "Result: {}",
                day2::sum_invalid_ids(inputs[0].as_str(), &day2::is_invalid)
            );
            Ok(())
        }
        (2, 2) => {
            println!(
                "Result: {}",
                day2::sum_invalid_ids(inputs[0].as_str(), &day2::is_invalid_all_repeats)
            );
            Ok(())
        }
        (3, 1) => {
            println!(
                "Result: {}",
                day3::sum_joltage(inputs.iter().map(String::as_str).collect(), 2)
            );
            Ok(())
        }
        (3, 2) => {
            println!(
                "Result: {}",
                day3::sum_joltage(inputs.iter().map(String::as_str).collect(), 12)
            );
            Ok(())
        }
        (4, 1) => {
            println!(
                "Result: {}",
                day4::count_rolls(inputs.iter().map(String::as_str).collect())
            );
            Ok(())
        }
        (4, 2) => {
            println!(
                "Result: {}",
                day4::count_removable_rolls(inputs.iter().map(String::as_str).collect())
            );
            Ok(())
        }
        (5, 1) => {
            println!(
                "Result: {}",
                day5::count_fresh_ids(inputs.iter().map(String::as_str).collect())
            );
            Ok(())
        }
        (5, 2) => {
            println!(
                "Result: {}",
                day5::count_total_fresh_ids(inputs.iter().map(String::as_str).collect())
            );
            Ok(())
        }
        _ => Err(format!("No match for Day {}, Part {}", args.day, args.part).into()),
    }
}
