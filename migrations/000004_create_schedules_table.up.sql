CREATE TABLE IF NOT EXISTS schedules (
    id        UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id   UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    date      VARCHAR(10) NOT NULL,
    time      VARCHAR(10) NOT NULL DEFAULT '',
    title     VARCHAR(255) NOT NULL,
    duration  VARCHAR(50) NOT NULL DEFAULT '',
    type      VARCHAR(50) NOT NULL DEFAULT '',
    notes     TEXT NOT NULL DEFAULT '',
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_schedules_user_id ON schedules (user_id);
CREATE INDEX IF NOT EXISTS idx_schedules_user_date ON schedules (user_id, date);
