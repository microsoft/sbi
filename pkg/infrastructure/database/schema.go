//     MIT License
//
//     Copyright (c) Microsoft Corporation.
//
//     Permission is hereby granted, free of charge, to any person obtaining a copy
//     of this software and associated documentation files (the "Software"), to deal
//     in the Software without restriction, including without limitation the rights
//     to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//     copies of the Software, and to permit persons to whom the Software is
//     furnished to do so, subject to the following conditions:
//
//     The above copyright notice and this permission notice shall be included in all
//     copies or substantial portions of the Software.
//
//     THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//     IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//     FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//     AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//     LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//     OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//     SOFTWARE

package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite" // Pure Go SQLite driver
)

// OpenDB opens or creates the SQLite database at the given path.
func OpenDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	// SQLite only supports one writer at a time. Limiting to a single
	// connection ensures PRAGMAs (like foreign_keys) stay in effect for
	// every operation, since database/sql may otherwise hand out new
	// connections that lack the per-connection PRAGMA settings.
	db.SetMaxOpenConns(1)

	// Enable WAL mode and foreign keys
	pragmas := []string{
		"PRAGMA journal_mode=WAL",
		"PRAGMA foreign_keys=ON",
	}

	for _, p := range pragmas {
		if _, err := db.Exec(p); err != nil {
			_ = db.Close()
			return nil, fmt.Errorf("setting pragma %q: %w", p, err)
		}
	}

	return db, nil
}

// CreateTables creates all required tables if they don't exist.
func CreateTables(db *sql.DB) error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS images (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE NOT NULL,
			registry TEXT,
			repository TEXT,
			tag TEXT,
			digest TEXT,
			size_bytes INTEGER,
			layers INTEGER,
			created_date TEXT,
			scan_timestamp TEXT,
			base_os_name TEXT,
			base_os_version TEXT,
			total_vulnerabilities INTEGER DEFAULT 0,
			critical_vulnerabilities INTEGER DEFAULT 0,
			high_vulnerabilities INTEGER DEFAULT 0,
			medium_vulnerabilities INTEGER DEFAULT 0,
			low_vulnerabilities INTEGER DEFAULT 0,
			negligible_vulnerabilities INTEGER DEFAULT 0,
			unknown_vulnerabilities INTEGER DEFAULT 0,
			vulnerability_scan_timestamp TEXT,
			vulnerability_scanner TEXT,
			secrets_found INTEGER DEFAULT 0,
			config_issues INTEGER DEFAULT 0,
			license_issues INTEGER DEFAULT 0,
			comprehensive_scan_timestamp TEXT,
			comprehensive_scanner TEXT,
			UNIQUE(registry, repository, tag)
		)`,
		`CREATE TABLE IF NOT EXISTS languages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			image_id INTEGER NOT NULL,
			language TEXT NOT NULL,
			version TEXT,
			major_minor TEXT,
			package_name TEXT,
			package_type TEXT,
			architecture TEXT,
			vendor TEXT,
			verified INTEGER DEFAULT 0,
			FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE CASCADE
		)`,
		`CREATE INDEX IF NOT EXISTS idx_languages_image_id ON languages(image_id)`,
		`CREATE INDEX IF NOT EXISTS idx_languages_lang_ver ON languages(language, version)`,
		`CREATE TABLE IF NOT EXISTS vulnerabilities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			image_id INTEGER NOT NULL,
			vulnerability_id TEXT,
			severity TEXT,
			package_name TEXT,
			package_version TEXT,
			fixed_version TEXT,
			description TEXT,
			cvss_score REAL,
			FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE CASCADE
		)`,
		`CREATE INDEX IF NOT EXISTS idx_vulnerabilities_severity ON vulnerabilities(severity)`,
		`CREATE INDEX IF NOT EXISTS idx_vulnerabilities_image_id ON vulnerabilities(image_id)`,
		`CREATE TABLE IF NOT EXISTS package_managers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			image_id INTEGER NOT NULL,
			name TEXT,
			version TEXT,
			language TEXT,
			FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS capabilities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			image_id INTEGER NOT NULL,
			capability TEXT,
			FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE CASCADE
		)`,
		`CREATE INDEX IF NOT EXISTS idx_capabilities_capability ON capabilities(capability)`,
		`CREATE INDEX IF NOT EXISTS idx_capabilities_image_id ON capabilities(image_id)`,
		`CREATE TABLE IF NOT EXISTS system_packages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			image_id INTEGER NOT NULL,
			name TEXT,
			version TEXT,
			package_type TEXT,
			FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS security_findings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			image_id INTEGER NOT NULL,
			finding_type TEXT,
			severity TEXT,
			rule_id TEXT,
			title TEXT,
			description TEXT,
			file_path TEXT,
			category TEXT,
			message TEXT,
			FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE CASCADE
		)`,
		`CREATE INDEX IF NOT EXISTS idx_security_findings_type ON security_findings(finding_type)`,
		`CREATE INDEX IF NOT EXISTS idx_security_findings_severity ON security_findings(severity)`,
		`CREATE INDEX IF NOT EXISTS idx_security_findings_image_id ON security_findings(image_id)`,
	}

	for _, stmt := range statements {
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("executing schema statement: %w", err)
		}
	}

	return nil
}
