use wasm_bindgen::prelude::*;

#[wasm_bindgen]
extern {
    pub fn alert(s: &str);
}

#[wasm_bindgen]
pub fn greet(name: &str) {
    alert(&format!("Hello, {}!", name));
}

#[wasm_bindgen]
pub fn n(r: u8, g: u8, b: u8) -> String {
    return String::from("")
}

#[wasm_bindgen]
pub fn t(r: u8, g: u8, b: u8) -> bool {
    return false
}
