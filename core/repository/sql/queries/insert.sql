-- name: InsertClient :execresult
INSERT INTO `client`
(`id`,`name`,`email`,`timezone`,`notes`,`created_at`,`updated_at`)
VALUES
(?,?,?,?,?,?,?);

-- name: InsertWebhookEvent :execresult
INSERT INTO `webhook_event`
(`id`,`source`,`event_type`,`payload`,`processed_at`,`created_at`)
VALUES
(?,?,?,?,?,?);

-- name: InsertBooking :execresult
INSERT INTO `booking`
(`id`,`client_id`,`status`,`review_decision`,`reviewed_at`,`decline_reason`,`pay_link_token`,`pay_link_expires_at`,`portal_token`,`intake`,`payment`,`scheduling`,`terms_version`,`terms_accepted_at`,`terms_accepted_ip`,`created_at`,`updated_at`)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertBookingDocument :execresult
INSERT INTO `booking_document`
(`id`,`booking_id`,`kind`,`url`,`label`,`purge_after`,`created_at`)
VALUES
(?,?,?,?,?,?,?);

-- name: InsertBookingRecap :execresult
INSERT INTO `booking_recap`
(`id`,`booking_id`,`body`,`published_at`,`created_at`)
VALUES
(?,?,?,?,?);

-- name: InsertNdaDocument :execresult
INSERT INTO `nda_document`
(`id`,`client_id`,`url`,`status`,`signed_at`,`created_at`,`envelope_id`,`certificate_url`)
VALUES
(?,?,?,?,?,?,?,?);

