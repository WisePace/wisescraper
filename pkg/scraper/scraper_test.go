package scraper

import (
	"log"
	"os"
	"testing"
)

func TestReadExistingEmails(t *testing.T) {
	fileName := "test_emails.txt"
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer file.Close()
	defer os.Remove(fileName)

	if _, err := file.WriteString("test@example.com\n"); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	emails, err := ReadExistingEmails(fileName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if _, exists := emails["test@example.com"]; !exists {
		t.Errorf("Expected email 'test@example.com' to be in the map")
	}
}

func TestReadDomains(t *testing.T) {
	fileName := "test_domains.txt"
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer file.Close()
	defer os.Remove(fileName)

	if _, err := file.WriteString("example.com\n"); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	domains, err := ReadDomains(fileName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(domains) != 1 || domains[0] != "example.com" {
		t.Errorf("Expected domain 'example.com', got %v", domains)
	}
}

func TestFindEmails(t *testing.T) {
	domains := []string{"example.com"}
	emailListFile, err := os.Create("test_email_list.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer emailListFile.Close()
	defer os.Remove("test_email_list.txt")

	logger := log.New(os.Stdout, "TEST: ", log.LstdFlags)
	existingEmails := make(map[string]struct{})

	foundEmails, scannedDomains := FindEmails(domains, emailListFile, logger, existingEmails)

	if foundEmails != 0 {
		t.Errorf("Expected 0 found emails, got %d", foundEmails)
	}

	if scannedDomains != 1 {
		t.Errorf("Expected 1 scanned domain, got %d", scannedDomains)
	}
}