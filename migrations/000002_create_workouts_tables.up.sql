CREATE TABLE IF NOT EXISTS workouts (
    id             UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id        UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name           VARCHAR(255) NOT NULL,
    date           VARCHAR(10) NOT NULL,
    duration       INT NOT NULL DEFAULT 0,
    weight_unit    VARCHAR(10) NOT NULL DEFAULT 'lbs',
    notes          TEXT NOT NULL DEFAULT '',
    completed      BOOLEAN NOT NULL DEFAULT FALSE,
    calories       INT NOT NULL DEFAULT 0,
    avg_heart_rate INT NOT NULL DEFAULT 0,
    exercises      JSONB NOT NULL DEFAULT '[]',
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_workouts_user_id ON workouts (user_id);
