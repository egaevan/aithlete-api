CREATE TABLE IF NOT EXISTS body_weight (
    id                  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id             UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    date                VARCHAR(10) NOT NULL,
    weight              DOUBLE PRECISION NOT NULL DEFAULT 0,
    body_fat_percentage DOUBLE PRECISION NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_body_weight_user_id ON body_weight (user_id);

CREATE TABLE IF NOT EXISTS strength_progression (
    id          UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    exercise    VARCHAR(255) NOT NULL,
    date        VARCHAR(10) NOT NULL,
    one_rep_max DOUBLE PRECISION NOT NULL DEFAULT 0,
    volume      DOUBLE PRECISION NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_strength_user_id ON strength_progression (user_id);

CREATE TABLE IF NOT EXISTS consistency (
    id                 UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id            UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    week               VARCHAR(10) NOT NULL,
    workouts_completed INT NOT NULL DEFAULT 0,
    workouts_planned   INT NOT NULL DEFAULT 0,
    streak             INT NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_consistency_user_id ON consistency (user_id);

CREATE TABLE IF NOT EXISTS muscle_volume (
    id           UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id      UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    muscle_group VARCHAR(100) NOT NULL,
    volume       DOUBLE PRECISION NOT NULL DEFAULT 0,
    sessions     INT NOT NULL DEFAULT 0,
    trend        VARCHAR(20) NOT NULL DEFAULT 'stable'
);

CREATE INDEX IF NOT EXISTS idx_muscle_volume_user_id ON muscle_volume (user_id);
