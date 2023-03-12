pub fn summarize_usecase(target_text: String, summary_length: usize) -> String {
    String::from("summarize")
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn success() {
        assert_eq!(summarize_usecase("Quick brown fox jumps over the lazy dog".to_string(), 5), "summarize");
    }
}


