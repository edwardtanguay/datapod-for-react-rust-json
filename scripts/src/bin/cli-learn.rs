mod qtools;
mod developer;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() > 1 {
        let param = &args[1]; 
        // let message = format!("EX{}", param); 
        // qtools::qcli::message(&message, "success");

        match param.as_str() {
            "001" => developer::learn::ex001(),
            "002" => developer::learn::ex002(),
            _ => qtools::qcli::message("Invalid parameter provided", "error"),
        }
    } else {
        qtools::qcli::message("no parameter provided", "error");
        qtools::qcli::message("Usage: npm run learn <exercise-number>", "info");
        qtools::qcli::message("Example: npm run 001", "info");
    }
}