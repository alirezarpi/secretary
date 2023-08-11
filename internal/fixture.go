package storage

import "secretary/alpha/utils"

func (admin *User) adminFixture() error {
	err := admin.CreateUser("admin", utils.GenerateRandomPassword(32), true)
	if err != nil {
		utils.Logger("warn", err.Error())
	}
	utils.Logger("info", "admin user successfully initiated")
	return nil
}

func RunFixtures() error {
	adminFixture()
	utils.Logger("info", "Fixtures successfully initiated")
}
