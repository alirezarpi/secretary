package internal

import "secretary/alpha/utils"

func adminFixture() error {
	admin := &User{}
	username := "admin"
	password := utils.GenerateRandomPassword(32)
	err := admin.CreateUser(username, password, true)
	if err != nil {
		utils.Logger("warn", err.Error())
		return nil
	}
	utils.Logger("info", "admin user successfully initiated")
	utils.Logger("info", "------------------------------------------")
	utils.Logger("info", "username: "+username)
	utils.Logger("info", "password: "+password)
	utils.Logger("info", "------------------------------------------")
	return nil
}

//func permissionFixture() error {
//	admin := &User{}
//	username := "admin"
//	password := utils.GenerateRandomPassword(32)
//	err := admin.CreateUser(username, password, true)
//	if err != nil {
//		utils.Logger("warn", err.Error())
//		return nil
//	}
//	utils.Logger("info", "admin user successfully initiated")
//	utils.Logger("info", "------------------------------------------")
//	utils.Logger("info", "username: " + username)
//	utils.Logger("info", "password: " + password)
//	utils.Logger("info", "------------------------------------------")
//	return nil
//}
//
//func roleFixture() error {
//	admin := &User{}
//	username := "admin"
//	password := utils.GenerateRandomPassword(32)
//	err := admin.CreateUser(username, password, true)
//	if err != nil {
//		utils.Logger("warn", err.Error())
//		return nil
//	}
//	utils.Logger("info", "admin user successfully initiated")
//	utils.Logger("info", "------------------------------------------")
//	utils.Logger("info", "username: " + username)
//	utils.Logger("info", "password: " + password)
//	utils.Logger("info", "------------------------------------------")
//	return nil
//}

func RunFixtures() error {
	adminFixture()
	utils.Logger("info", "Fixtures successfully initiated")
	return nil
}
