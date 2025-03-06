package services

import "testing"

func TestGetCityByZipSuccess(t *testing.T) {
	city, err := GetCityByZip("01001000")
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	if city != "São Paulo" {
		t.Errorf("Esperava São Paulo, mas recebeu %s", city)
	}
}

func TestGetCityByZipNotFound(t *testing.T) {
	city, err := GetCityByZip("00000000")
	if err == nil {
		t.Errorf("Esperava erro, mas recebeu %s", city)
	}
}

func TestGetCityByZipInvalidCEP(t *testing.T) {
	city, err := GetCityByZip("123")
	if err == nil {
		t.Errorf("Esperava erro, mas recebeu %s", city)
	}
}

func TestGetCityByZipInvalidCEPFormat(t *testing.T) {
	city, err := GetCityByZip("1234567")
	if err == nil {
		t.Errorf("Esperava erro, mas recebeu %s", city)
	}
}
