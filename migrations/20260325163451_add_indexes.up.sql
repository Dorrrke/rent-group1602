CREATE INDEX IF NOT EXISTS idx_rents_user_id ON rents(user_id);
CREATE INDEX IF NOT EXISTS idx_rents_car_id ON rents(car_id);
CREATE INDEX IF NOT EXISTS idx_cars_number ON cars(number);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);