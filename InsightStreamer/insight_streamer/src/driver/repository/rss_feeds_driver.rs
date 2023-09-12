use sqlx::{MySql, Pool};


#[derive(Clone)]
pub struct DatabasePool {
    pub(crate) pool: Pool<MySql>,
}

#[derive(Debug, Clone)]
pub struct FeedRepository {
    pub pool: Pool<MySql>,
}

impl FeedRepository {
    pub fn new(pool: DatabasePool) -> Self {
        FeedRepository { pool: pool.pool }
    }
}