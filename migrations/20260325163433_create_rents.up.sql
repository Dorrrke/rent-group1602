CREATE TABLE IF NOT EXISTS rents (
    rid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    car_id UUID NOT NULL,
    user_id UUID NOT NULL,
    hours INT NOT NULL,
    ended BOOLEAN NOT NULL DEFAULT false, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_rents_car
        FOREIGN KEY (car_id) 
        REFERENCES cars(cid)
        ON DELETE CASCADE,

    CONSTRAINT fk_rents_user
        FOREIGN KEY (user_id) 
        REFERENCES users(uid)
        ON DELETE CASCADE
);