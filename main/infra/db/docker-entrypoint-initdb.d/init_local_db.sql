CREATE USER sample IDENTIFIED BY 'bne585W3SjZBwKzB8jZyiqM5JN2q2MYQ';

DROP DATABASE IF EXISTS sample_local;
CREATE DATABASE IF NOT EXISTS sample_local CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
GRANT ALL ON `sample_local`.* TO 'sample'@'%';

DROP DATABASE IF EXISTS sample_test;
CREATE DATABASE IF NOT EXISTS sample_test CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
GRANT ALL ON `sample_test`.* TO 'sample'@'%';

