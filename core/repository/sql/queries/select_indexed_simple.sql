

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
SELECT `id`,`client_id`,`status`,`review_decision`,`reviewed_at`,`decline_reason`,`intake`,`payment`,`scheduling`,`terms_version`,`terms_accepted_at`,`terms_accepted_ip`,`created_at`,`updated_at`
FROM `booking`
WHERE 
    `id` = ? ;

        
-- name: FetchBookingByIdForUpdate :many
SELECT `id`,`client_id`,`status`,`review_decision`,`reviewed_at`,`decline_reason`,`intake`,`payment`,`scheduling`,`terms_version`,`terms_accepted_at`,`terms_accepted_ip`,`created_at`,`updated_at`
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
        




-- engagement_agreement selects:
-- name: FetchEngagementAgreementById :many
SELECT `id`,`client_id`,`nda_url`,`status`,`signed_at`,`created_at`,`envelope_id`,`certificate_url`,`contract_url`,`engagement_inquiry_id`,`updated_at`
FROM `engagement_agreement`
WHERE 
    `id` = ? ;

        
-- name: FetchEngagementAgreementByIdForUpdate :many
SELECT `id`,`client_id`,`nda_url`,`status`,`signed_at`,`created_at`,`envelope_id`,`certificate_url`,`contract_url`,`engagement_inquiry_id`,`updated_at`
FROM `engagement_agreement`
WHERE 
    `id` = ? 
FOR UPDATE;
        




-- engagement_inquiry selects:
-- name: FetchEngagementInquiryById :many
SELECT `id`,`client_id`,`name`,`email`,`phone`,`company`,`project_summary`,`why_more_than_session`,`scope_details`,`budget_range`,`timeline`,`status`,`review_notes`,`created_at`,`updated_at`
FROM `engagement_inquiry`
WHERE 
    `id` = ? ;

        
-- name: FetchEngagementInquiryByIdForUpdate :many
SELECT `id`,`client_id`,`name`,`email`,`phone`,`company`,`project_summary`,`why_more_than_session`,`scope_details`,`budget_range`,`timeline`,`status`,`review_notes`,`created_at`,`updated_at`
FROM `engagement_inquiry`
WHERE 
    `id` = ? 
FOR UPDATE;
        




-- magic_link selects:
-- name: FetchMagicLinkById :many
SELECT `id`,`client_id`,`email`,`token`,`purpose`,`expires_at`,`consumed_at`,`created_at`,`created_ip`
FROM `magic_link`
WHERE 
    `id` = ? ;

        
-- name: FetchMagicLinkByIdForUpdate :many
SELECT `id`,`client_id`,`email`,`token`,`purpose`,`expires_at`,`consumed_at`,`created_at`,`created_ip`
FROM `magic_link`
WHERE 
    `id` = ? 
FOR UPDATE;
        


