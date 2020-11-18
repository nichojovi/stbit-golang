-- CREATE TABLE "movie_log" ---------------------------
CREATE TABLE `movie_log` ( 
	`id` BigInt( 20 ) UNSIGNED AUTO_INCREMENT NOT NULL,
	`user_name` BigInt( 20 ) UNSIGNED NOT NULL,
	`search_word` VarChar( 150 ) NOT NULL,
	`pagination` BigInt( 20 ) UNSIGNED NOT NULL,
	`created_at` DateTime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` DateTime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY ( `id` ))
ENGINE = InnoDB;