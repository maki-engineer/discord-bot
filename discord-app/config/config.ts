require('dotenv').config();

export = {
  production: {
    url: process.env.POSTGRES_URL!,
    dialect: 'postgres',
    dialectOptions: { ssl: { require: true } },
    timezone: '+09:00',
  },
  development: {
    username: process.env.POSTGRES_USER_DEVELOPMENT!,
    password: process.env.POSTGRES_PASSWORD_DEVELOPMENT!,
    database: process.env.POSTGRES_DB_DEVELOPMENT!,
    host: 'development-db',
    port: 5432,
    dialect: 'postgres',
    timezone: '+09:00',
  },
  unittest: {
    username: process.env.POSTGRES_USER_UNITTEST!,
    password: process.env.POSTGRES_PASSWORD_UNITTEST!,
    database: process.env.POSTGRES_DB_UNITTEST!,
    host: '127.0.0.1',
    port: 5433,
    dialect: 'postgres',
    timezone: '+09:00',
  },
};
