-- +migrate Up
ALTER TABLE `users`
ADD COLUMN `notification_token` VARCHAR(255) DEFAULT NULL;

-- +migrate Down
ALTER TABLE `users`
DROP COLUMN `notification_token`;