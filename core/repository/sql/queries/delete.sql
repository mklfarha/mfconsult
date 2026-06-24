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

-- name: DeleteNdaDocument :execresult
DELETE FROM `nda_document`
WHERE
`id` = ?;

