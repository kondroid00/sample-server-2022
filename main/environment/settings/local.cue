db: {
	dbconfigs: [{
		name:     "main"
		dbms:     "mysql"
		user:     "sample"
		password: "bne585W3SjZBwKzB8jZyiqM5JN2q2MYQ"
		protocol: "tcp"
		host:     "db"
		port:     3306
		schema:   "sample_local"
		option:   "charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true&sql_safe_updates=ON&sql_mode=%27TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY%27"
		max_open: 50
		max_idle: 10
		conn_lifetime_sec: 10
	}, {
		name:     "read"
		dbms:     "mysql"
		user:     "sample"
		password: "bne585W3SjZBwKzB8jZyiqM5JN2q2MYQ"
		protocol: "tcp"
		host:     "db"
		port:     3306
		schema:   "sample_local"
		option:   "charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true&sql_safe_updates=ON&sql_mode=%27TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY%27"
		max_open: 50
		max_idle: 10
		conn_lifetime_sec: 10
	}]
	debug_mode: true
}
sentry: {
	config: {
		dsn: "https://example.com"
	}
	debug_mode: false
}