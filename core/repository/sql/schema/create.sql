CREATE TABLE IF NOT EXISTS `client` (
    `id` CHAR(36) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(512) NOT NULL,
    `timezone` VARCHAR(64),
    `notes` TEXT,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `webhook_event` (
    `id` VARCHAR(128) NOT NULL,
    `source` INT,
    `event_type` VARCHAR(128),
    `payload` TEXT,
    `processed_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `booking` (
    `id` CHAR(36) NOT NULL,
    `client_id` CHAR(36) NOT NULL,
    `status` INT NOT NULL,
    `review_decision` INT NOT NULL,
    `reviewed_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `decline_reason` TEXT,
    `pay_link_token` VARCHAR(64),
    `pay_link_expires_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `portal_token` VARCHAR(64),
    `intake` JSON,
    `payment` JSON,
    `scheduling` JSON,
    `terms_version` VARCHAR(32),
    `terms_accepted_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `terms_accepted_ip` VARCHAR(64),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `client_has_bookings`
        FOREIGN KEY (`client_id`)
        REFERENCES `client` (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `booking_document` (
    `id` CHAR(36) NOT NULL,
    `booking_id` CHAR(36) NOT NULL,
    `kind` INT,
    `url` VARCHAR(512),
    `label` VARCHAR(255),
    `purge_after` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `booking_has_documents`
        FOREIGN KEY (`booking_id`)
        REFERENCES `booking` (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `booking_recap` (
    `id` CHAR(36) NOT NULL,
    `booking_id` CHAR(36) NOT NULL,
    `body` TEXT,
    `published_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `booking_has_recap`
        FOREIGN KEY (`booking_id`)
        REFERENCES `booking` (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `nda_document` (
    `id` CHAR(36) NOT NULL,
    `client_id` CHAR(36) NOT NULL,
    `url` VARCHAR(512),
    `status` INT,
    `signed_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `envelope_id` VARCHAR(128),
    `certificate_url` VARCHAR(512),
    CONSTRAINT `client_has_nda`
        FOREIGN KEY (`client_id`)
        REFERENCES `client` (`id`)
) ENGINE = InnoDB;

