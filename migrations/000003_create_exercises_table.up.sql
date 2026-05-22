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

INSERT INTO exercises (name, description, muscle_group, equipment, difficulty, instructions) VALUES
('Bench Press', 'Barbell bench press targeting chest, shoulders, and triceps', 'chest', 'barbell', 'intermediate', ARRAY['Lie flat on bench with feet on floor', 'Grip bar slightly wider than shoulder-width', 'Lower bar to mid-chest', 'Press bar up to full arm extension']),
('Incline Dumbbell Press', 'Dumbbell press on incline bench for upper chest', 'chest', 'dumbbell', 'intermediate', ARRAY['Set bench to 30-45 degree incline', 'Hold dumbbells at shoulder height', 'Press up until arms are extended', 'Lower slowly back to start']),
('Push-Up', 'Bodyweight chest exercise', 'chest', 'bodyweight', 'beginner', ARRAY['Start in plank position with hands shoulder-width', 'Lower body until chest nearly touches floor', 'Push back up to start position', 'Keep core engaged throughout']),
('Cable Fly', 'Cable crossover for chest isolation', 'chest', 'cable', 'intermediate', ARRAY['Set pulleys to high position', 'Stand in center, grab handles', 'Bring hands together in front of chest', 'Slowly return to start']),
('Pull-Up', 'Compound back exercise targeting lats and biceps', 'back', 'bodyweight', 'intermediate', ARRAY['Grip bar with palms facing away', 'Hang with arms fully extended', 'Pull up until chin clears bar', 'Lower with control to full hang']),
('Barbell Row', 'Bent-over barbell row for mid-back thickness', 'back', 'barbell', 'intermediate', ARRAY['Hinge at hips with slight knee bend', 'Grip bar shoulder-width', 'Pull bar to lower ribcage', 'Lower bar with control']),
('Lat Pulldown', 'Cable pulldown targeting lats', 'back', 'cable', 'beginner', ARRAY['Sit at pulldown station with thighs secured', 'Grip bar wider than shoulder-width', 'Pull bar down to upper chest', 'Slowly return to full extension']),
('Deadlift', 'Full-body compound lift primarily targeting back and legs', 'back', 'barbell', 'advanced', ARRAY['Stand with feet hip-width, bar over midfoot', 'Hinge at hips and grip bar', 'Drive through heels to stand', 'Lower bar with controlled hinge']),
('Squat', 'Barbell back squat for quadriceps and glutes', 'legs', 'barbell', 'intermediate', ARRAY['Position bar on upper back', 'Stand feet shoulder-width', 'Squat down to parallel or below', 'Drive up through heels to stand']),
('Leg Press', 'Machine-based compound leg exercise', 'legs', 'machine', 'beginner', ARRAY['Sit in leg press machine', 'Place feet shoulder-width on platform', 'Lower platform until knees at 90 degrees', 'Press back up without locking knees']),
('Romanian Deadlift', 'Hamstring and glute focused hip hinge', 'legs', 'barbell', 'intermediate', ARRAY['Hold bar at hip height', 'Hinge at hips pushing butt back', 'Lower bar along legs until hamstring stretch', 'Drive hips forward to return']),
('Walking Lunge', 'Dynamic lunge for quad and glute development', 'legs', 'dumbbell', 'beginner', ARRAY['Hold dumbbells at sides', 'Step forward into lunge', 'Lower back knee toward ground', 'Drive through front heel to stand and step forward with other leg']),
('Overhead Press', 'Standing barbell press for shoulders', 'shoulders', 'barbell', 'intermediate', ARRAY['Stand with feet shoulder-width', 'Hold bar at shoulder height', 'Press bar overhead to full extension', 'Lower bar back to shoulders']),
('Lateral Raise', 'Dumbbell lateral raise for side delts', 'shoulders', 'dumbbell', 'beginner', ARRAY['Stand with dumbbells at sides', 'Raise arms out to sides to shoulder height', 'Keep slight bend in elbows', 'Lower with control']),
('Face Pull', 'Cable face pull for rear delts and rotator cuff', 'shoulders', 'cable', 'beginner', ARRAY['Set pulley at upper-chest height', 'Grip rope attachment with both hands', 'Pull toward face with elbows high', 'Squeeze shoulder blades at peak']),
('Barbell Curl', 'Standing barbell curl for biceps', 'arms', 'barbell', 'beginner', ARRAY['Stand with bar at hip height', 'Curl bar toward shoulders', 'Keep elbows pinned to sides', 'Lower with control']),
('Tricep Pushdown', 'Cable pushdown for triceps', 'arms', 'cable', 'beginner', ARRAY['Attach straight bar to high pulley', 'Grip bar with elbows at sides', 'Press bar down until arms straight', 'Slowly return to start']),
('Hammer Curl', 'Dumbbell hammer curl for brachialis and biceps', 'arms', 'dumbbell', 'beginner', ARRAY['Hold dumbbells with neutral grip', 'Curl toward shoulders', 'Keep palms facing each other', 'Lower with control']),
('Skull Crusher', 'Lying tricep extension with EZ bar', 'arms', 'barbell', 'intermediate', ARRAY['Lie on bench holding EZ bar above chest', 'Lower bar toward forehead', 'Extend arms back to start', 'Keep upper arms stationary']),
('Plank', 'Core stability exercise', 'core', 'bodyweight', 'beginner', ARRAY['Start in forearm plank position', 'Keep body in straight line from head to heels', 'Engage core and glutes', 'Hold position for desired duration']),
('Hanging Leg Raise', 'Advanced core exercise targeting lower abs', 'core', 'bodyweight', 'advanced', ARRAY['Hang from pull-up bar', 'Raise legs until parallel to ground', 'Lower with control', 'Avoid swinging']),
('Cable Crunch', 'Cable crunch for rectus abdominis', 'core', 'cable', 'beginner', ARRAY['Attach rope to high pulley', 'Kneel facing away from pulley', 'Crunch forward bringing elbows to knees', 'Slowly return to start']),
('Hip Thrust', 'Barbell hip thrust for glute activation and growth', 'glutes', 'barbell', 'intermediate', ARRAY['Sit on floor with back against bench', 'Place bar across hips', 'Drive hips up squeezing glutes at top', 'Lower with control']),
('Standing Calf Raise', 'Standing calf raise for gastrocnemius', 'calves', 'machine', 'beginner', ARRAY['Stand on calf raise machine', 'Place shoulders under pads', 'Rise up on toes as high as possible', 'Lower heels below platform stretch']),
('Seated Calf Raise', 'Seated calf raise for soleus', 'calves', 'machine', 'beginner', ARRAY['Sit in seated calf raise machine', 'Place toes on platform', 'Raise heels as high as possible', 'Lower with control']),
('Wrist Curl', 'Wrist curl for forearm flexors', 'forearms', 'dumbbell', 'beginner', ARRAY['Sit on bench holding dumbbell', 'Rest forearm on thigh', 'Curl wrist up', 'Lower with control']),
('Barbell Shrug', 'Barbell shrug for trapezius', 'traps', 'barbell', 'beginner', ARRAY['Stand holding bar at hip height', 'Shrug shoulders straight up', 'Hold peak contraction briefly', 'Lower with control']),
('Burpee', 'Full-body conditioning exercise', 'full-body', 'bodyweight', 'intermediate', ARRAY['Stand with feet shoulder-width', 'Drop into squat and place hands on floor', 'Kick feet back into plank', 'Jump feet back to squat and explode up']);
