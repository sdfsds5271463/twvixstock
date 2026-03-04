SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
DROP TABLE IF EXISTS `stock`;
CREATE TABLE `stock`  (
`id` int UNSIGNED NOT NULL AUTO_INCREMENT,
`type` enum('TAIEX','VIXTWN') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
`date` date NULL DEFAULT NULL,
`close` decimal(10, 2) NULL DEFAULT NULL,
PRIMARY KEY (`id`) USING BTREE,
INDEX `date`(`date`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;
-- ----------------------------
SET FOREIGN_KEY_CHECKS = 1;
INSERT INTO `stock` (type, date, close) VALUES ('TAIEX', '2026-02-11', 17800.50);
INSERT INTO `stock` (type, date, close) VALUES ('VIXTWN', '2026-02-12', 15.20);