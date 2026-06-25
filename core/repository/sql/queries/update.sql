-- name: UpdateClient :exec
UPDATE `client`
SET
`name` = ?, `email` = ?, `timezone` = ?, `notes` = ?, `created_at` = ?, `updated_at` = ?
WHERE
`id` = ?;

-- name: UpdateWebhookEvent :exec
UPDATE `webhook_event`
SET
`source` = ?, `event_type` = ?, `payload` = ?, `processed_at` = ?, `created_at` = ?
WHERE
`id` = ?;

-- name: UpdateBooking :exec
UPDATE `booking`
SET
`client_id` = ?, `status` = ?, `review_decision` = ?, `reviewed_at` = ?, `decline_reason` = ?, `intake` = ?, `payment` = ?, `scheduling` = ?, `terms_version` = ?, `terms_accepted_at` = ?, `terms_accepted_ip` = ?, `created_at` = ?, `updated_at` = ?
WHERE
`id` = ?;

-- name: UpdateBookingDocument :exec
UPDATE `booking_document`
SET
`booking_id` = ?, `kind` = ?, `url` = ?, `label` = ?, `purge_after` = ?, `created_at` = ?
WHERE
`id` = ?;

-- name: UpdateBookingRecap :exec
UPDATE `booking_recap`
SET
`booking_id` = ?, `body` = ?, `published_at` = ?, `created_at` = ?
WHERE
`id` = ?;

-- name: UpdateEngagementAgreement :exec
UPDATE `engagement_agreement`
SET
`client_id` = ?, `nda_url` = ?, `status` = ?, `signed_at` = ?, `created_at` = ?, `envelope_id` = ?, `certificate_url` = ?, `contract_url` = ?, `engagement_inquiry_id` = ?, `updated_at` = ?
WHERE
`id` = ?;

-- name: UpdateEngagementInquiry :exec
UPDATE `engagement_inquiry`
SET
`client_id` = ?, `name` = ?, `email` = ?, `phone` = ?, `company` = ?, `project_summary` = ?, `why_more_than_session` = ?, `scope_details` = ?, `budget_range` = ?, `timeline` = ?, `status` = ?, `review_notes` = ?, `created_at` = ?, `updated_at` = ?
WHERE
`id` = ?;

-- name: UpdateMagicLink :exec
UPDATE `magic_link`
SET
`client_id` = ?, `email` = ?, `token` = ?, `purpose` = ?, `expires_at` = ?, `consumed_at` = ?, `created_at` = ?, `created_ip` = ?
WHERE
`id` = ?;

