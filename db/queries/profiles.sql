-- name: CreateProfile :exec
INSERT INTO profiles (
  user_id, name, age, birth_date, birth_time, birth_place,
  nakshatra, rashi, gender, height, father_name, mother_name,
  total_family_members, qualification, degree, college,
  designation, company_and_city, city, district, state,
  marital_status, add_details, mobile_number, profile_photo_url
) VALUES (
  $1, $2, $3, $4, $5, $6,
  $7, $8, $9, $10, $11, $12,
  $13, $14, $15, $16,
  $17, $18, $19, $20, $21,
  $22, $23, $24, $25
);

-- name: GetProfile :one
SELECT * FROM profiles WHERE user_id = $1;

-- name: UpdateProfile :exec
UPDATE profiles
SET
  name = $2, age = $3, birth_date = $4, birth_time = $5,
  birth_place = $6, nakshatra = $7, rashi = $8, gender = $9,
  height = $10, father_name = $11, mother_name = $12,
  total_family_members = $13, qualification = $14, degree = $15,
  college = $16, designation = $17, company_and_city = $18,
  city = $19, district = $20, state = $21, marital_status = $22,
  add_details = $23, mobile_number = $24, profile_photo_url = $25
WHERE user_id = $1;
