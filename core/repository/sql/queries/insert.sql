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
(`id`,`client_id`,`status`,`review_decision`,`reviewed_at`,`decline_reason`,`intake`,`payment`,`scheduling`,`terms_version`,`terms_accepted_at`,`terms_accepted_ip`,`created_at`,`updated_at`)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?);

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

-- name: InsertEngagementAgreement :execresult
INSERT INTO `engagement_agreement`
(`id`,`client_id`,`nda_url`,`status`,`signed_at`,`created_at`,`envelope_id`,`certificate_url`,`contract_url`,`engagement_inquiry_id`,`updated_at`)
VALUES
(?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertEngagementInquiry :execresult
INSERT INTO `engagement_inquiry`
(`id`,`client_id`,`name`,`email`,`phone`,`company`,`project_summary`,`why_more_than_session`,`scope_details`,`budget_range`,`timeline`,`status`,`review_notes`,`created_at`,`updated_at`)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertMagicLink :execresult
INSERT INTO `magic_link`
(`id`,`client_id`,`email`,`token`,`purpose`,`expires_at`,`consumed_at`,`created_at`,`created_ip`)
VALUES
(?,?,?,?,?,?,?,?,?);

