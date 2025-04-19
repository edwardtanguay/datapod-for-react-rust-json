use chrono::Utc;

pub fn debug(line: &str) -> &str {
    let timestamp = Utc::now().format("%Y-%m-%d %H:%M:%S").to_string(); // Format the timestamp
    println!("🛠️  {} - {}", timestamp, line);
    line
}