ALTER TABLE consistency ADD CONSTRAINT consistency_user_week_idx UNIQUE (user_id, week);
ALTER TABLE muscle_volume ADD CONSTRAINT muscle_volume_user_muscle_idx UNIQUE (user_id, muscle_group);
