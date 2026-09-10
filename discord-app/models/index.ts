import { Dialect, Sequelize } from 'sequelize';
import DictWord from './dictword';
import BirthdayFor235Member from './birthdayfor235member';
import BirthdayForMillionMember from './birthdayformillionmember';
import DeleteMessage from './deletemessage';
import Session from './session';
import config from '../config/config';

type Environment = 'production' | 'development' | 'unittest';

let env: Environment = 'development';

if (process.env.NODE_ENV === 'production') {
  env = 'production';
} else if (process.env.NODE_ENV === 'unittest') {
  env = 'unittest';
}

let sequelize!: Sequelize;

if (env === 'production') {
  const productionConfig = config[env];

  sequelize = new Sequelize(productionConfig.url, {
    dialectOptions: {
      ssl: { require: true },
    },
  });
} else if (env === 'development') {
  const developmentConfig = config[env];

  sequelize = new Sequelize(
    developmentConfig.database,
    developmentConfig.username,
    developmentConfig.password,
    {
      ...developmentConfig,
      dialect: developmentConfig.dialect as Dialect,
    },
  );
} else if (env === 'unittest') {
  const unittestConfig = config[env];

  sequelize = new Sequelize(
    unittestConfig.database,
    unittestConfig.username,
    unittestConfig.password,
    {
      ...unittestConfig,
      dialect: unittestConfig.dialect as Dialect,
    },
  );
}

const db: {
  BirthdayFor235Member: typeof BirthdayFor235Member;
  BirthdayForMillionMember: typeof BirthdayForMillionMember;
  DictWord: typeof DictWord;
  DeleteMessage: typeof DeleteMessage;
  Session: typeof Session;
  sequelize?: typeof sequelize;
  Sequelize?: typeof Sequelize;
} = {
  BirthdayFor235Member: BirthdayFor235Member.initialize(sequelize),
  BirthdayForMillionMember: BirthdayForMillionMember.initialize(sequelize),
  DictWord: DictWord.initialize(sequelize),
  DeleteMessage: DeleteMessage.initialize(sequelize),
  Session: Session.initialize(sequelize),
};

db.sequelize = sequelize;
db.Sequelize = Sequelize;

export default db;
