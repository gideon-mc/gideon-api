package gideon

import "github.com/gideon-mc/gideon-api/pkg/db"

func ClaimTables(db *db.DB) {
	// WARNING: It doesn't actually use FOREIGN KEYs due to PlanetScale limitations.

	db.ClaimTable("user", []string{
		"`user_id` varchar(16) PRIMARY KEY",
		"`password_hash` char(64) NOT NULL",
		"`is_admin` boolean NOT NULL DEFAULT false",
		"`is_editor` boolean NOT NULL DEFAULT false",
		"`is_banned` boolean NOT NULL DEFAULT false",
		"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP",
		"`last_seen_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP",
		"`fk_discord_id` char(18) NOT NULL",
	})

	db.ClaimTable("player", []string{
		"`player_id` char(32) PRIMARY KEY",
		"`owner_id` char(32)",
		"`fk_discord_id` char(18) NOT NULL",
		"`fk_attachment_id` bigint",
	})

	db.ClaimTable("gametag", []string{
		"`gametag_id` bigint PRIMARY KEY AUTO_INCREMENT",
		"`name` varchar(16) NOT NULL",
		"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP",
		"`fk_player_id` char(32) NOT NULL",
	})

	db.ClaimTable("discord", []string{
		"`discord_id` char(18) PRIMARY KEY",
		"`slug` char(32) NOT NULL",
		"`is_hidden` boolean NOT NULL DEFAULT true",
		"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP",
		"`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
	})

	db.ClaimTable("attachment", []string{
		"`attachment_id` bigint PRIMARY KEY AUTO_INCREMENT",
		"`urls` text NOT NULL",
		"`is_hidden` boolean NOT NULL DEFAULT true",
		"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP",
		"`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
	})

	db.ClaimTable("player_clan", []string{
		"`player_clan_id` bigint PRIMARY KEY AUTO_INCREMENT",
		"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP",
		"`fk_role_id` bigint NOT NULL",
		"`fk_player_id` bigint NOT NULL",
		"`fk_clan_id` bigint NOT NULL",
	})

	db.ClaimTable("role", []string{
		"`role_id` bigint PRIMARY KEY AUTO_INCREMENT",
		"`name` varchar(16) NOT NULL",
		"`color_hex` char(6) NOT NULL",
	})

	db.ClaimTable("alliance", []string{
		"`alliance_id` bigint PRIMARY KEY AUTO_INCREMENT",
		"`name` varchar(16) NOT NULL",
	})

	db.ClaimTable("clan", []string{
		"`clan_id` bigint PRIMARY KEY AUTO_INCREMENT",
		"`name` varchar(16) NOT NULL",
		"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP",
		"`fk_alliance_id` bigint",
	})
}
