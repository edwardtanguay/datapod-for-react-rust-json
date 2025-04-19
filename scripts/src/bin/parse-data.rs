use std::fs::File;
use std::io::Write;
use serde::Serialize;

mod qtools;

#[derive(Serialize)]
struct Flashcard {
    #[allow(dead_code)]
    suuid: String,
    #[allow(dead_code)]
    category: String,
    #[allow(dead_code)]
    front: String,
    #[allow(dead_code)]
    back: String,
}

fn save_flashcards_to_file(flashcards: &[Flashcard], output_file: &str) -> Result<(), std::io::Error> {
    let json = serde_json::to_string_pretty(flashcards).expect("Failed to serialize flashcards");
    let mut file = File::create(output_file)?;
    file.write_all(json.as_bytes())?;
    Ok(())
}

fn main() {
    let input_file = "../data/flashcards.txt";

    let lines = match qtools::qfil::get_lines_from_file(input_file) {
        Ok(lines) => lines,
        Err(err) => {
            eprintln!("Error reading file {}: {}", input_file, err);
            return;
        }
    };

    println!("Read {} lines from {}", lines.len(), input_file);

    for (index, line) in lines.iter().enumerate() {
        println!("Line {}: {}", index + 1, line);
    }

    let mut flashcards = Vec::new();

    for i in (0..lines.len()).step_by(4) {
        let category = lines.get(i).cloned().unwrap_or_default();
        let front = lines.get(i + 1).cloned().unwrap_or_default();
        let back = lines.get(i + 2).cloned().unwrap_or_default();

        if !category.is_empty() && !front.is_empty() && !back.is_empty() {
            let suuid = qtools::qstr::generate_suuid();
            flashcards.push(Flashcard {
                suuid,
                category,
                front,
                back,
            });
        } else {
            eprintln!("WARNING: Skipping incomplete flashcard at lines {}-{}", i + 1, i + 3);
        }
    }

    println!("Generated {} flashcards", flashcards.len());

    for (index, flashcard) in flashcards.iter().enumerate() {
        println!(
            "Flashcard {}:\n  UUID: {}\n  Category: {}\n  Front: {}\n  Back: {}\n",
            index + 1,
            flashcard.suuid,
            flashcard.category,
            flashcard.front,
            flashcard.back
        );
    }

    // Save flashcards to JSON file
    let output_file = "../parseddata/flashcards.json";
    match save_flashcards_to_file(&flashcards, output_file) {
        Ok(_) => println!("Wrote {} flashcards to {}", flashcards.len(), output_file),
        Err(err) => eprintln!("Error writing flashcards to JSON file: {}", err),
    }
}
