package audit

import (
	"fmt"
	"time"
	"os"
	"io"

	"secretary/alpha/utils"
)

func createAuditFile() (io.Writer, error) {
	directory := "persistence/audit/"
	utils.MakeDir(directory)
	currentDate := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("audit_%s", currentDate)
	file, err := os.OpenFile(directory+filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func Audit(message string) error {
	file, err := createAuditFile()
	if err != nil {
		utils.Logger("fatal", err.Error())
		return err
	}
	defer file.(*os.File).Close()
	
	_, err = file.Write([]byte(utils.CurrentTime()+" - "+message + "\n"))
	if err != nil {
		utils.Logger("fatal", err.Error())
		return err
	}
	return nil
}
