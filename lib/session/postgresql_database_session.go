package session
//
//import (
//	"fmt"
//	"time"
//
//	"secretary/alpha/storage"
//	"secretary/alpha/utils"
//)
//
//func PostgreSQLDatabaseSession(resource_id string, access_type string, ttl ...string) (error, map[string]interface{}) {
//    // Pull Resource data from db
//
//    // Build the connection string for the admin user
//    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//        host, port, user, password, dbname)
//
//    // Connect to the database
//    db, err := sql.Open("postgres", psqlInfo)
//    if err != nil {
//        log.Fatalf("Unable to connect to the database: %v", err)
//    }
//    defer db.Close()
//
//    // Verify the connection
//    err = db.Ping()
//    if err != nil {
//        log.Fatalf("Unable to ping the database: %v", err)
//    }
//
//    log.Println("Successfully connected to the database.")
//
//    // User credentials
//    newUsername := "readonly_user"
//    newPassword := "readonly_password"
//    userTTL := "1 day" // Set TTL for the user
//
//    // Create the user with TTL
//    createUserQuery := fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s' VALID UNTIL NOW() + INTERVAL '%s';", newUsername, newPassword, userTTL)
//    _, err = db.Exec(createUserQuery)
//    if err != nil {
//        log.Fatalf("Error creating user: %v", err)
//    }
//    log.Printf("User '%s' created successfully with TTL of %s.\n", newUsername, userTTL)
//
//    // Grant SELECT role on the database
//    grantRoleQuery := fmt.Sprintf("GRANT CONNECT ON DATABASE %s TO %s;", dbname, newUsername)
//    _, err = db.Exec(grantRoleQuery)
//    if err != nil {
//        log.Fatalf("Error granting CONNECT privilege: %v", err)
//    }
//
//    grantSelectQuery := fmt.Sprintf("GRANT SELECT ON ALL TABLES IN SCHEMA public TO %s;", newUsername)
//    _, err = db.Exec(grantSelectQuery)
//    if err != nil {
//        log.Fatalf("Error granting SELECT privilege: %v", err)
//    }
//
//    log.Printf("SELECT privileges granted to user '%s'.\n", newUsername)
//
//    // Ensure future tables grant SELECT permissions automatically
//    autoGrantQuery := fmt.Sprintf(`ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO %s;`, newUsername)
//    _, err = db.Exec(autoGrantQuery)
//    if err != nil {
//        log.Fatalf("Error setting default privileges: %v", err)
//    }
//
//    log.Println("Default SELECT privileges set successfully.")
//}
