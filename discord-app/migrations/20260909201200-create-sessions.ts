/** @type {import('sequelize-cli').Migration} */
module.exports = {
  up: async (queryInterface: any, Sequelize: any) => {
    await queryInterface.createTable('sessions', {
      session_id: {
        allowNull: false,
        unique: true,
        type: Sequelize.STRING,
      },
      expires_at: {
        allowNull: false,
        type: Sequelize.DATE,
      },
    });
  },

  down: async (queryInterface: any, Sequelize: any) => {
    await queryInterface.dropTable('sessions');
  },
};
