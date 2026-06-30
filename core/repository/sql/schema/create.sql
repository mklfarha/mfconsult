CREATE TABLE IF NOT EXISTS `client` (
    `id` CHAR(36) NOT NULL,
    `name` VARCHAR(255),
    `email` VARCHAR(512) NOT NULL,
    `timezone` VARCHAR(64),
    `notes` TEXT,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `webhook_event` (
    `id` VARCHAR(128) NOT NULL,
    `source` INT,
    `event_type` VARCHAR(128),
    `payload` TEXT,
    `processed_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `booking` (
    `id` CHAR(36) NOT NULL,
    `client_id` CHAR(36) NOT NULL,
    `status` INT NOT NULL,
    `review_decision` INT NOT NULL,
    `reviewed_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `decline_reason` TEXT,
    `intake` JSON,
    `payment` JSON,
    `scheduling` JSON,
    `terms_version` VARCHAR(32),
    `terms_accepted_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `terms_accepted_ip` VARCHAR(64),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
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
    PRIMARY KEY (`id`),
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
    PRIMARY KEY (`id`),
    CONSTRAINT `booking_has_recap`
        FOREIGN KEY (`booking_id`)
        REFERENCES `booking` (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `engagement_agreement` (
    `id` CHAR(36) NOT NULL,
    `client_id` CHAR(36) NOT NULL,
    `nda_url` VARCHAR(512),
    `status` INT,
    `signed_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `envelope_id` VARCHAR(128),
    `certificate_url` VARCHAR(512),
    `contract_url` VARCHAR(512),
    `engagement_inquiry_id` CHAR(36) NOT NULL,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    CONSTRAINT `client_has_agreement`
        FOREIGN KEY (`client_id`)
        REFERENCES `client` (`id`),
    CONSTRAINT `engagement_has_agreement`
        FOREIGN KEY (`engagement_inquiry_id`)
        REFERENCES `engagement_inquiry` (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `engagement_inquiry` (
    `id` CHAR(36) NOT NULL,
    `client_id` CHAR(36),
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(512) NOT NULL,
    `phone` VARCHAR(32),
    `company` VARCHAR(255),
    `project_summary` VARCHAR(255) NOT NULL,
    `why_more_than_session` TEXT,
    `scope_details` VARCHAR(255),
    `budget_range` VARCHAR(128),
    `timeline` VARCHAR(128),
    `status` INT,
    `review_notes` TEXT,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    CONSTRAINT `client_has_inquiries`
        FOREIGN KEY (`client_id`)
        REFERENCES `client` (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `magic_link` (
    `id` CHAR(36) NOT NULL,
    `client_id` CHAR(36) NOT NULL,
    `email` VARCHAR(512),
    `token` VARCHAR(128) NOT NULL,
    `purpose` INT,
    `expires_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `consumed_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `created_ip` VARCHAR(64),
    PRIMARY KEY (`id`),
    CONSTRAINT `client_has_magic_links`
        FOREIGN KEY (`client_id`)
        REFERENCES `client` (`id`)
) ENGINE = InnoDB;

