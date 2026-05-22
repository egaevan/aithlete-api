CREATE TABLE IF NOT EXISTS exercises (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    muscle_group VARCHAR(50) NOT NULL,
    equipment VARCHAR(100) NOT NULL DEFAULT '',
    difficulty VARCHAR(20) NOT NULL DEFAULT 'beginner',
    instructions TEXT[] NOT NULL DEFAULT '{}',
    image_url VARCHAR(500) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_exercises_muscle_group ON exercises(muscle_group);
CREATE INDEX idx_exercises_difficulty ON exercises(difficulty);
