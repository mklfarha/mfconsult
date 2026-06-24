

-- client selects:
-- name: FetchClientById :many
SELECT `id`,`name`,`email`,`timezone`,`notes`,`created_at`,`updated_at`
FROM `client`
WHERE 
    `id` = ? ;

        
-- name: FetchClientByIdForUpdate :many
SELECT `id`,`name`,`email`,`timezone`,`notes`,`created_at`,`updated_at`
FROM `client`
WHERE 
    `id` = ? 
FOR UPDATE;
        




-- webhook_event selects:
-- name: FetchWebhookEventById :many
SELECT `id`,`source`,`event_type`,`payload`,`processed_at`,`created_at`
FROM `webhook_event`
WHERE 
    `id` = ? ;

        
-- name: FetchWebhookEventByIdForUpdate :many
SELECT `id`,`source`,`event_type`,`payload`,`processed_at`,`created_at`
FROM `webhook_event`
WHERE 
    `id` = ? 
FOR UPDATE;
        




-- booking selects:
-- name: FetchBookingById :many
SELECT `id`,`client_id`,`status`,`review_decision`,`reviewed_at`,`decline_reason`,`pay_link_token`,`pay_link_expires_at`,`portal_token`,`intake`,`payment`,`scheduling`,`terms_version`,`terms_accepted_at`,`terms_accepted_ip`,`created_at`,`updated_at`
FROM `booking`
WHERE 
    `id` = ? ;

        
-- name: FetchBookingByIdForUpdate :many
SELECT `id`,`client_id`,`status`,`review_decision`,`reviewed_at`,`decline_reason`,`pay_link_token`,`pay_link_expires_at`,`portal_token`,`intake`,`payment`,`scheduling`,`terms_version`,`terms_accepted_at`,`terms_accepted_ip`,`created_at`,`updated_at`
FROM `booking`
WHERE 
    `id` = ? 
FOR UPDATE;
        




-- booking_document selects:
-- name: FetchBookingDocumentById :many
SELECT `id`,`booking_id`,`kind`,`url`,`label`,`purge_after`,`created_at`
FROM `booking_document`
WHERE 
    `id` = ? ;

        
-- name: FetchBookingDocumentByIdForUpdate :many
SELECT `id`,`booking_id`,`kind`,`url`,`label`,`purge_after`,`created_at`
FROM `booking_document`
WHERE 
    `id` = ? 
FOR UPDATE;
        




-- booking_recap selects:
-- name: FetchBookingRecapById :many
SELECT `id`,`booking_id`,`body`,`published_at`,`created_at`
FROM `booking_recap`
WHERE 
    `id` = ? ;

        
-- name: FetchBookingRecapByIdForUpdate :many
SELECT `id`,`booking_id`,`body`,`published_at`,`created_at`
FROM `booking_recap`
WHERE 
    `id` = ? 
FOR UPDATE;
        




-- nda_document selects:
-- name: FetchNdaDocumentById :many
SELECT `id`,`client_id`,`url`,`status`,`signed_at`,`created_at`,`envelope_id`,`certificate_url`
FROM `nda_document`
WHERE 
    `id` = ? ;

        
-- name: FetchNdaDocumentByIdForUpdate :many
SELECT `id`,`client_id`,`url`,`status`,`signed_at`,`created_at`,`envelope_id`,`certificate_url`
FROM `nda_document`
WHERE 
    `id` = ? 
FOR UPDATE;
        


