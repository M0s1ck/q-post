ALTER TABLE community.users
DROP COLUMN followers_count,
ADD COLUMN followees_count,
ADD COLUMN friends_count;