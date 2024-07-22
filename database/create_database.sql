SELECT 'CREATE DATABASE db_time_tracker' 
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'db_time_tracker')\gexec