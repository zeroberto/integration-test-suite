package infra

import (
	"database/sql"
	"net"
	"os/exec"
	"strings"
	"time"

	_ "github.com/lib/pq" // Required database driver
)

// DownInfra is responsible for bringing down the infrastructure
func DownInfra(dockerComposeFileName string) {
	cmd := exec.Command("docker-compose", "-f", dockerComposeFileName, "down")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// UpInfra is responsible for raising the infrastructure
func UpInfra(dockerComposeFileName string) {
	cmd := exec.Command("docker-compose", "-f", dockerComposeFileName, "up", "-d")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// GetServiceID is responsible for returning a docker service identifier
func GetServiceID(dockerComposeFileName string, serviceName string) string {
	out, err := exec.Command("docker-compose", "-f", dockerComposeFileName, "ps", "-q", serviceName).Output()
	if err != nil {
		panic(err)
	}

	return strings.TrimSuffix(string(out), "\n")
}

// StopService is responsible for stopping a docker service
func StopService(serviceID string) {
	cmd := exec.Command("docker", "stop", serviceID)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// CheckDBConnection is responsible for checking whether database is accepting connections
func CheckDBConnection(dbType string, dsn string) error {
	db, err := sql.Open(dbType, dsn)
	if err == nil {
		err = db.Ping()
		if err == nil {
			db.Close()
			return nil
		}
	}
	return err
}

// CheckPortIsOpen is responsible for checking that a port is open for connections
func CheckPortIsOpen(host string, port string) bool {
	timeout := time.Second
	if conn, _ := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout); conn != nil {
		defer conn.Close()
		return true
	}
	return false
}
