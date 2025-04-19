pub fn message(line: &str, kind: &str) {
    match kind {
        "info" => println!("ℹ️  {}", line),
        "error" => println!("❌ {}", line),
        "warning" => println!("️🟨 {}", line),
        "doing" => println!("⏳ {}", line),
        _ => println!("✅ {}", line),
    }
}