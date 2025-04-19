use rand::Rng;

pub fn generate_suuid() -> String {
    let mut rng = rand::rng();
    let random_string: String = (0..6)
        .map(|_| rng.random_range(0..62)) // Generate a random index
        .map(|idx| {
            if idx < 10 {
                // 0-9
                (b'0' + idx as u8) as char
            } else if idx < 36 {
                // A-Z
                (b'A' + (idx - 10) as u8) as char
            } else {
                // a-z
                (b'a' + (idx - 36) as u8) as char
            }
        })
        .collect();
    random_string
}