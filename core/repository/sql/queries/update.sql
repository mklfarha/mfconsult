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
`client_id` = ?, `status` = ?, `review_decision` = ?, `reviewed_at` = ?, `decline_reason` = ?, `pay_link_token` = ?, `pay_link_expires_at` = ?, `portal_token` = ?, `intake` = ?, `payment` = ?, `scheduling` = ?, `terms_version` = ?, `terms_accepted_at` = ?, `terms_accepted_ip` = ?, `created_at` = ?, `updated_at` = ?
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

-- name: UpdateNdaDocument :exec
UPDATE `nda_document`
SET
`client_id` = ?, `url` = ?, `status` = ?, `signed_at` = ?, `created_at` = ?, `envelope_id` = ?, `certificate_url` = ?
WHERE
`id` = ?;

