import { Sequelize, Model, DataTypes } from 'sequelize';

export default class Session extends Model {
  public session_id!: string;

  public expires_at!: Date;

  static initialize(sequelize: Sequelize) {
    this.init(
      {
        session_id: {
          type: DataTypes.STRING,
          allowNull: false,
          unique: true,
        },
        expires_at: {
          type: DataTypes.DATE,
          allowNull: false,
        },
      },
      {
        sequelize,
        modelName: 'Session',
        tableName: 'sessions',
        timestamps: false,
      },
    );

    this.removeAttribute('id');

    return this;
  }
}
