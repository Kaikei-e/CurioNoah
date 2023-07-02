use chrono::{DateTime, Utc};

pub enum AuditLogAction {
    Upsert,
    Delete,
    Fail,
}

impl AuditLogAction {
    pub fn convert_to_string(&self) -> String {
        match self {
            AuditLogAction::Upsert => "update".to_string(),
            AuditLogAction::Delete => "delete".to_string(),
            AuditLogAction::Fail => "fail".to_string(),
        }
    }
}


pub struct AuditLog {
    pub action: AuditLogAction,
    pub updated_at: DateTime<Utc>,
}

// struct Updated{
//     action_name: String,
//     updated_at: DateTime<Utc>,
// }
//
// struct Deleted{
//     action_name: String,
//     updated_at: DateTime<Utc>,
// }
//
// struct Failed{
//     action_name: String,
//     updated_at: DateTime<Utc>,
// }