use crate::qtools;

pub fn show_title(ex_code: &str, title: &str) {
    let full_title = format!("EX{}: {}", ex_code, title.to_uppercase()); 
    qtools::qcli::message(&full_title, "success");
}

pub fn ex001() {
    show_title("001", "numbers");
    let percent1: f32 = 0.5;
    let percent2: f64 = 0.5;
    let number_of_items: i8 = 5;

    let sum1 = percent1 + percent2 as f32;
    let sum2: f64 = percent1 as f64 + number_of_items as f64;
    let sum3: f64 = percent2 + number_of_items as f64;
    qtools::qcli::message(&format!("percent1: {}", percent1), "info");
    qtools::qcli::message(&format!("percent2: {}", percent2), "info");
    qtools::qcli::message(&format!("number of items: {}", number_of_items), "info");
    qtools::qcli::message(&format!("sum1: {}", sum1), "info");
    qtools::qcli::message(&format!("sum2: {}", sum2), "info");
    qtools::qcli::message(&format!("sum3: {}", sum3), "info");
}

pub fn ex002() {
    println!("this is 002");
}