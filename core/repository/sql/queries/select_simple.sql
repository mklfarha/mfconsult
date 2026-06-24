-- name: FetchClient :many
SELECT `id`,`name`,`email`,`timezone`,`notes`,`created_at`,`updated_at`
FROM `client`;

-- name: FetchWebhookEvent :many
SELECT `id`,`source`,`event_type`,`payload`,`processed_at`,`created_at`
FROM `webhook_event`;

-- name: FetchBooking :many
SELECT `id`,`client_id`,`status`,`review_decision`,`reviewed_at`,`decline_reason`,`pay_link_token`,`pay_link_expires_at`,`portal_token`,`intake`,`payment`,`scheduling`,`terms_version`,`terms_accepted_at`,`terms_accepted_ip`,`created_at`,`updated_at`
FROM `booking`;

-- name: FetchBookingDocument :many
SELECT `id`,`booking_id`,`kind`,`url`,`label`,`purge_after`,`created_at`
FROM `booking_document`;

-- name: FetchBookingRecap :many
SELECT `id`,`booking_id`,`body`,`published_at`,`created_at`
FROM `booking_recap`;

-- name: FetchNdaDocument :many
SELECT `id`,`client_id`,`url`,`status`,`signed_at`,`created_at`,`envelope_id`,`certificate_url`
FROM `nda_document`;

