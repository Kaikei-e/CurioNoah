pub fn summarize(target_text: String, summary_length: usize) -> String {
    String::from("summarize")
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn success() {
        assert_eq!(summarize("Quick brown fox jumps over the lazy dog".to_string(), 5), "summarize");
    }
}