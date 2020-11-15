package main

import (
	"log"
	"log/syslog"
	"os"
)

const logTransport = "tcp"
const logServer = "crm.asart.local:514"

/* *********************************************************************** */
/*                          Methodes of the package                        */
/* *********************************************************************** */

// RemoteSyslog sends message to remote syslog server
// In case of the remote syslog unavailable, the method writes
// to local syslog
// @message: message string
// @severity: syslog severity level
func RemoteSyslog(message string, severity syslog.Priority) error {
	logger, err := syslog.Dial(logTransport, logServer, severity, os.Args[0])

	if err != nil {
		localLogger, err := syslog.New(severity, os.Args[0])
		if err != nil {
			return err
		}
		log.SetOutput(localLogger)
		log.Print(message)
		return err
	}
	log.SetOutput(logger)
	log.Print(message)
	return nil
}
