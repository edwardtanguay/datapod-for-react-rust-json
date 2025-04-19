mod qtools;

fn main() {
    qtools::qdev::debug("we are inside the create-page script");
    qtools::qcli::message("this script will create a page in the React app", "info");
}