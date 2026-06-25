-- name: DeleteClient :execresult
DELETE FROM `client`
WHERE
`id` = ?;

-- name: DeleteWebhookEvent :execresult
DELETE FROM `webhook_event`
WHERE
`id` = ?;

-- name: DeleteBooking :execresult
DELETE FROM `booking`
WHERE
`id` = ?;

-- name: DeleteBookingDocument :execresult
DELETE FROM `booking_document`
WHERE
`id` = ?;

-- name: DeleteBookingRecap :execresult
DELETE FROM `booking_recap`
WHERE
`id` = ?;

-- name: DeleteEngagementAgreement :execresult
DELETE FROM `engagement_agreement`
WHERE
`id` = ?;

-- name: DeleteEngagementInquiry :execresult
DELETE FROM `engagement_inquiry`
WHERE
`id` = ?;

-- name: DeleteMagicLink :execresult
DELETE FROM `magic_link`
WHERE
`id` = ?;

