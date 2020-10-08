package infra

import "testing"

func TestUpInfra(t *testing.T) {
	fileName := "tests/docker-compose.yml"

	UpInfra(fileName)
	defer DownInfra(fileName)

	got := GetServiceID(fileName, "test")

	if got == "" {
		t.Errorf("UpInfra() failed, at least one service was expected")
	}
}

func TestUpInfraWithEnvs(t *testing.T) {
	expected := "testing"

	fileName := "tests/docker-compose.yml"

	UpInfraWithEnvs(fileName, map[string]string{
		"TEST": "testing",
	})
	defer DownInfra(fileName)

	serviceID := GetServiceID(fileName, "test")

	got := GetContainerEnvValue(serviceID, "TEST")

	if got == "" {
		t.Errorf("UpInfraWithEnvs() failed, expected %v, got %v", expected, nil)
	}
}
