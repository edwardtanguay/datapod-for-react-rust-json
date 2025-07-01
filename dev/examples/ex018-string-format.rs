fn main() {
    let mut title = String::from("The Report 002");
    title.push_str(" (original)");
    let message: String = format!("Message: {}", title);
    println!("{}", message);
}
