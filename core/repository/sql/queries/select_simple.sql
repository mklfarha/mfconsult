-- name: FetchClient :many
SELECT `id`,`name`,`email`,`timezone`,`notes`,`created_at`,`updated_at`
FROM `client`;

-- name: FetchWebhookEvent :many
SELECT `id`,`source`,`event_type`,`payload`,`processed_at`,`created_at`
FROM `webhook_event`;

-- name: FetchBooking :many
SELECT `id`,`client_id`,`status`,`review_decision`,`reviewed_at`,`decline_reason`,`intake`,`payment`,`scheduling`,`terms_version`,`terms_accepted_at`,`terms_accepted_ip`,`created_at`,`updated_at`
FROM `booking`;

-- name: FetchBookingDocument :many
SELECT `id`,`booking_id`,`kind`,`url`,`label`,`purge_after`,`created_at`
FROM `booking_document`;

-- name: FetchBookingRecap :many
SELECT `id`,`booking_id`,`body`,`published_at`,`created_at`
FROM `booking_recap`;

-- name: FetchEngagementAgreement :many
SELECT `id`,`client_id`,`nda_url`,`status`,`signed_at`,`created_at`,`envelope_id`,`certificate_url`,`contract_url`,`engagement_inquiry_id`,`updated_at`
FROM `engagement_agreement`;

-- name: FetchEngagementInquiry :many
SELECT `id`,`client_id`,`name`,`email`,`phone`,`company`,`project_summary`,`why_more_than_session`,`scope_details`,`budget_range`,`timeline`,`status`,`review_notes`,`created_at`,`updated_at`
FROM `engagement_inquiry`;

-- name: FetchMagicLink :many
SELECT `id`,`client_id`,`email`,`token`,`purpose`,`expires_at`,`consumed_at`,`created_at`,`created_ip`
FROM `magic_link`;

