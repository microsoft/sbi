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
	"time"

	"github.com/microsoft/sbi/pkg/domain"
	log "github.com/sirupsen/logrus"
)

// Repository provides database operations for image data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository with the given database connection.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ImageExists checks whether an image with the given name already exists.
func (r *Repository) ImageExists(name string) (bool, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM images WHERE name = ?", name).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("checking image existence: %w", err)
	}

	return count > 0, nil
}

// InsertImage inserts an image and all its related data in a transaction.
// err is a named result so the deferred rollback always sees failures from
// every return path (including short-declaration shadows in nested scopes).
func (r *Repository) InsertImage(img *domain.ImageRecord) (err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}
	// Defer is registered only after Begin succeeds, so tx is always non-nil here.
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Errorf("rollback failed: %v", rbErr)
			}
		}
	}()

	scanTS := time.Now().UTC().Format(time.RFC3339)

	res, err := tx.Exec(`
		INSERT INTO images (
			name, registry, repository, tag, digest, size_bytes, layers,
			created_date, scan_timestamp, base_os_name, base_os_version,
			total_vulnerabilities, critical_vulnerabilities, high_vulnerabilities,
			medium_vulnerabilities, low_vulnerabilities, negligible_vulnerabilities,
			unknown_vulnerabilities, vulnerability_scan_timestamp, vulnerability_scanner,
			secrets_found, config_issues, license_issues,
			comprehensive_scan_timestamp, comprehensive_scanner
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(name) DO UPDATE SET
			digest=excluded.digest, size_bytes=excluded.size_bytes, layers=excluded.layers,
			created_date=excluded.created_date, scan_timestamp=excluded.scan_timestamp,
			base_os_name=excluded.base_os_name, base_os_version=excluded.base_os_version,
			total_vulnerabilities=excluded.total_vulnerabilities,
			critical_vulnerabilities=excluded.critical_vulnerabilities,
			high_vulnerabilities=excluded.high_vulnerabilities,
			medium_vulnerabilities=excluded.medium_vulnerabilities,
			low_vulnerabilities=excluded.low_vulnerabilities,
			negligible_vulnerabilities=excluded.negligible_vulnerabilities,
			unknown_vulnerabilities=excluded.unknown_vulnerabilities,
			vulnerability_scan_timestamp=excluded.vulnerability_scan_timestamp,
			vulnerability_scanner=excluded.vulnerability_scanner,
			secrets_found=excluded.secrets_found,
			config_issues=excluded.config_issues,
			license_issues=excluded.license_issues,
			comprehensive_scan_timestamp=excluded.comprehensive_scan_timestamp,
			comprehensive_scanner=excluded.comprehensive_scanner`,
		img.Name, img.Registry, img.Repository, img.Tag, img.Digest,
		img.SizeBytes, img.Layers, img.CreatedDate, scanTS,
		img.BaseOSName, img.BaseOSVersion,
		img.TotalVulnerabilities, img.CriticalVulnerabilities, img.HighVulnerabilities,
		img.MediumVulnerabilities, img.LowVulnerabilities, img.NegligibleVulnerabilities,
		img.UnknownVulnerabilities, img.VulnerabilityScanTimestamp, img.VulnerabilityScanner,
		img.SecretsFound, img.ConfigIssues, img.LicenseIssues,
		img.ComprehensiveScanTimestamp, img.ComprehensiveScanner,
	)
	if err != nil {
		return fmt.Errorf("inserting image: %w", err)
	}

	// On upsert (ON CONFLICT DO UPDATE), LastInsertId() is unreliable — it may
	// return 0 or a stale value from a previous insert. Always query by name.
	_ = res
	var imageID int64
	err = tx.QueryRow("SELECT id FROM images WHERE name = ?", img.Name).Scan(&imageID)
	if err != nil {
		return fmt.Errorf("getting image id: %w", err)
	}

	// Clear old related data on upsert.
	// Assign to the named err (not :=) so deferred rollback observes failures.
	relTables := []string{"languages", "vulnerabilities", "package_managers", "capabilities", "system_packages", "security_findings"}
	for _, t := range relTables {
		if _, err = tx.Exec(fmt.Sprintf("DELETE FROM %s WHERE image_id = ?", t), imageID); err != nil {
			return fmt.Errorf("clearing %s: %w", t, err)
		}
	}

	// Insert languages
	for _, l := range img.Languages {
		_, err = tx.Exec(`INSERT INTO languages (image_id, language, version, major_minor, package_name, package_type, architecture, vendor, verified)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			imageID, l.Language, l.Version, l.MajorMinor, l.PackageName, l.PackageType, l.Architecture, l.Vendor, l.Verified)
		if err != nil {
			return fmt.Errorf("inserting language: %w", err)
		}
	}

	// Insert vulnerabilities
	for _, v := range img.Vulnerabilities {
		_, err = tx.Exec(`INSERT INTO vulnerabilities (image_id, vulnerability_id, severity, package_name, package_version, fixed_version, description, cvss_score)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			imageID, v.VulnerabilityID, v.Severity, v.PackageName, v.PackageVersion, v.FixedVersion, v.Description, v.CVSSScore)
		if err != nil {
			return fmt.Errorf("inserting vulnerability: %w", err)
		}
	}

	// Insert package managers
	for _, pm := range img.PackageManagers {
		_, err = tx.Exec(`INSERT INTO package_managers (image_id, name, version, language) VALUES (?, ?, ?, ?)`,
			imageID, pm.Name, pm.Version, pm.Language)
		if err != nil {
			return fmt.Errorf("inserting package manager: %w", err)
		}
	}

	// Insert capabilities
	for _, c := range img.Capabilities {
		_, err = tx.Exec(`INSERT INTO capabilities (image_id, capability) VALUES (?, ?)`,
			imageID, c.Capability)
		if err != nil {
			return fmt.Errorf("inserting capability: %w", err)
		}
	}

	// Insert system packages
	for _, sp := range img.SystemPackages {
		_, err = tx.Exec(`INSERT INTO system_packages (image_id, name, version, package_type) VALUES (?, ?, ?, ?)`,
			imageID, sp.Name, sp.Version, sp.PackageType)
		if err != nil {
			return fmt.Errorf("inserting system package: %w", err)
		}
	}

	// Insert security findings
	for _, sf := range img.SecurityFindings {
		_, err = tx.Exec(`INSERT INTO security_findings (image_id, finding_type, severity, rule_id, title, description, file_path, category, message)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			imageID, sf.FindingType, sf.Severity, sf.RuleID, sf.Title, sf.Description, sf.FilePath, sf.Category, sf.Message)
		if err != nil {
			return fmt.Errorf("inserting security finding: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

// QueryLanguages returns distinct languages ordered by best security posture.
// Languages whose top image has the fewest critical → high → total vulns appear first.
// The special "base" language (no-runtime images) always sorts last.
func (r *Repository) QueryLanguages() ([]string, error) {
	rows, err := r.db.Query(`
		SELECT LOWER(l.language) AS lang
		FROM languages l
		JOIN images i ON l.image_id = i.id
		WHERE l.language != ''
		GROUP BY lang
		ORDER BY CASE WHEN LOWER(l.language) = 'base' THEN 1 ELSE 0 END ASC,
		         MIN(i.critical_vulnerabilities) ASC,
		         MIN(i.high_vulnerabilities) ASC,
		         MIN(i.total_vulnerabilities) ASC,
		         lang ASC`)
	if err != nil {
		return nil, fmt.Errorf("querying languages: %w", err)
	}
	defer func() { _ = rows.Close() }()

	var languages []string
	for rows.Next() {
		var lang string
		if err := rows.Scan(&lang); err != nil {
			return nil, fmt.Errorf("scanning language: %w", err)
		}

		languages = append(languages, lang)
	}

	return languages, rows.Err()
}

// QueryTopImages returns the top N images for a given language, ranked by
// critical → high → total vulnerabilities → size ascending.
func (r *Repository) QueryTopImages(language string, topN int) ([]domain.RecommendedImage, error) {
	rows, err := r.db.Query(`
		SELECT i.name, l.version, i.critical_vulnerabilities, i.high_vulnerabilities,
		       i.total_vulnerabilities, COALESCE(i.size_bytes, 0),
		       COALESCE(i.created_date, ''), i.digest
		FROM images i
		JOIN languages l ON i.id = l.image_id
		WHERE LOWER(l.language) = ?
		GROUP BY i.id, l.version
		ORDER BY i.critical_vulnerabilities ASC,
		         i.high_vulnerabilities ASC,
		         i.total_vulnerabilities ASC,
		         COALESCE(i.size_bytes, 0) ASC
		LIMIT ?`, language, topN)
	if err != nil {
		return nil, fmt.Errorf("querying top images: %w", err)
	}
	defer func() { _ = rows.Close() }()

	var results []domain.RecommendedImage
	for rows.Next() {
		var img domain.RecommendedImage
		var digest sql.NullString

		if err := rows.Scan(&img.Name, &img.Version, &img.CriticalVulnerabilities,
			&img.HighVulnerabilities, &img.TotalVulnerabilities, &img.SizeBytes,
			&img.CreatedDate, &digest); err != nil {
			return nil, fmt.Errorf("scanning image row: %w", err)
		}

		if digest.Valid {
			img.Digest = digest.String
		}

		results = append(results, img)
	}

	return results, rows.Err()
}

// QueryBaseOSes returns distinct base OS names for a given language, sorted
// alphabetically with "Other" (empty/unknown) last.
func (r *Repository) QueryBaseOSes(language string) ([]string, error) {
	rows, err := r.db.Query(`
		SELECT DISTINCT COALESCE(NULLIF(i.base_os_name, ''), 'Other') AS os
		FROM images i
		JOIN languages l ON i.id = l.image_id
		WHERE LOWER(l.language) = ?
		ORDER BY
			CASE WHEN COALESCE(NULLIF(i.base_os_name, ''), 'Other') = 'Other' THEN 1 ELSE 0 END,
			os ASC`, language)
	if err != nil {
		return nil, fmt.Errorf("querying base OSes: %w", err)
	}
	defer func() { _ = rows.Close() }()

	var oses []string
	for rows.Next() {
		var os string
		if err := rows.Scan(&os); err != nil {
			return nil, fmt.Errorf("scanning base OS: %w", err)
		}
		oses = append(oses, os)
	}

	return oses, rows.Err()
}

// QueryTopImagesByOS returns the top N images for a given language and base OS,
// ranked by critical → high → total vulnerabilities → size ascending.
// Pass "Other" as baseOS to match images with empty/unknown OS.
// Pass topN <= 0 for unlimited results.
func (r *Repository) QueryTopImagesByOS(language, baseOS string, topN int) ([]domain.RecommendedImage, error) {
	limit := topN
	if limit <= 0 {
		limit = -1 // SQLite LIMIT -1 means no limit
	}

	rows, err := r.db.Query(`
		SELECT i.name, l.version, i.critical_vulnerabilities, i.high_vulnerabilities,
		       i.total_vulnerabilities, COALESCE(i.size_bytes, 0),
		       COALESCE(i.created_date, ''), i.digest,
		       COALESCE(NULLIF(i.base_os_name, ''), 'Other')
		FROM images i
		JOIN languages l ON i.id = l.image_id
		WHERE LOWER(l.language) = ?
		  AND COALESCE(NULLIF(i.base_os_name, ''), 'Other') = ?
		GROUP BY i.id, l.version
		ORDER BY i.critical_vulnerabilities ASC,
		         i.high_vulnerabilities ASC,
		         i.total_vulnerabilities ASC,
		         COALESCE(i.size_bytes, 0) ASC
		LIMIT ?`, language, baseOS, limit)
	if err != nil {
		return nil, fmt.Errorf("querying top images by OS: %w", err)
	}
	defer func() { _ = rows.Close() }()

	var results []domain.RecommendedImage
	for rows.Next() {
		var img domain.RecommendedImage
		var digest sql.NullString

		if err := rows.Scan(&img.Name, &img.Version, &img.CriticalVulnerabilities,
			&img.HighVulnerabilities, &img.TotalVulnerabilities, &img.SizeBytes,
			&img.CreatedDate, &digest, &img.BaseOSName); err != nil {
			return nil, fmt.Errorf("scanning image row: %w", err)
		}

		if digest.Valid {
			img.Digest = digest.String
		}

		results = append(results, img)
	}

	return results, rows.Err()
}

// ClearDatabase deletes all data from all tables.
func (r *Repository) ClearDatabase() error {
	tables := []string{"security_findings", "system_packages", "capabilities", "package_managers", "vulnerabilities", "languages", "images"}
	for _, t := range tables {
		if _, err := r.db.Exec(fmt.Sprintf("DELETE FROM %s", t)); err != nil {
			return fmt.Errorf("clearing table %s: %w", t, err)
		}
	}

	return nil
}

// QueryAllImageDetails returns every image with its child data populated
// (languages, vulnerabilities, system packages, package managers).
// Uses batched eager loading within a read transaction for snapshot consistency.
func (r *Repository) QueryAllImageDetails() ([]domain.ImageRecord, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("beginning read transaction: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	// Step 1: Load all images
	imgRows, err := tx.Query(`
		SELECT id, name, registry, repository, tag, digest,
		       COALESCE(size_bytes, 0), layers, COALESCE(created_date, ''),
		       COALESCE(scan_timestamp, ''), COALESCE(base_os_name, ''),
		       COALESCE(base_os_version, ''),
		       total_vulnerabilities, critical_vulnerabilities, high_vulnerabilities,
		       medium_vulnerabilities, low_vulnerabilities, negligible_vulnerabilities,
		       unknown_vulnerabilities
		FROM images
		ORDER BY name ASC`)
	if err != nil {
		return nil, fmt.Errorf("querying images: %w", err)
	}

	imageMap := make(map[int64]*domain.ImageRecord)
	var imageOrder []int64

	for imgRows.Next() {
		var img domain.ImageRecord
		if err := imgRows.Scan(
			&img.ID, &img.Name, &img.Registry, &img.Repository, &img.Tag, &img.Digest,
			&img.SizeBytes, &img.Layers, &img.CreatedDate,
			&img.ScanTimestamp, &img.BaseOSName, &img.BaseOSVersion,
			&img.TotalVulnerabilities, &img.CriticalVulnerabilities, &img.HighVulnerabilities,
			&img.MediumVulnerabilities, &img.LowVulnerabilities, &img.NegligibleVulnerabilities,
			&img.UnknownVulnerabilities,
		); err != nil {
			_ = imgRows.Close()
			return nil, fmt.Errorf("scanning image: %w", err)
		}

		imageMap[img.ID] = &img
		imageOrder = append(imageOrder, img.ID)
	}

	if err := imgRows.Close(); err != nil {
		return nil, fmt.Errorf("closing image rows: %w", err)
	}

	if err := imgRows.Err(); err != nil {
		return nil, fmt.Errorf("iterating images: %w", err)
	}

	if len(imageMap) == 0 {
		return nil, nil
	}

	// Step 2: Load languages
	langRows, err := tx.Query(`
		SELECT image_id, language, COALESCE(version, ''), COALESCE(major_minor, ''),
		       COALESCE(package_name, ''), COALESCE(package_type, ''), verified
		FROM languages ORDER BY image_id, id`)
	if err != nil {
		return nil, fmt.Errorf("querying languages: %w", err)
	}

	for langRows.Next() {
		var imageID int64
		var l domain.Language
		var verified int

		if err := langRows.Scan(&imageID, &l.Language, &l.Version, &l.MajorMinor,
			&l.PackageName, &l.PackageType, &verified); err != nil {
			_ = langRows.Close()
			return nil, fmt.Errorf("scanning language: %w", err)
		}

		l.Verified = verified != 0

		if img, ok := imageMap[imageID]; ok {
			img.Languages = append(img.Languages, l)
		}
	}

	if err := langRows.Close(); err != nil {
		return nil, fmt.Errorf("closing language rows: %w", err)
	}

	if err := langRows.Err(); err != nil {
		return nil, fmt.Errorf("iterating languages: %w", err)
	}

	// Step 3: Load vulnerabilities
	vulnRows, err := tx.Query(`
		SELECT image_id, COALESCE(vulnerability_id, ''), COALESCE(severity, ''),
		       COALESCE(package_name, ''), COALESCE(package_version, ''),
		       COALESCE(fixed_version, ''), COALESCE(description, ''),
		       COALESCE(cvss_score, 0)
		FROM vulnerabilities ORDER BY image_id, id`)
	if err != nil {
		return nil, fmt.Errorf("querying vulnerabilities: %w", err)
	}

	for vulnRows.Next() {
		var imageID int64
		var v domain.Vulnerability

		if err := vulnRows.Scan(&imageID, &v.VulnerabilityID, &v.Severity,
			&v.PackageName, &v.PackageVersion, &v.FixedVersion,
			&v.Description, &v.CVSSScore); err != nil {
			_ = vulnRows.Close()
			return nil, fmt.Errorf("scanning vulnerability: %w", err)
		}

		if img, ok := imageMap[imageID]; ok {
			img.Vulnerabilities = append(img.Vulnerabilities, v)
		}
	}

	if err := vulnRows.Close(); err != nil {
		return nil, fmt.Errorf("closing vulnerability rows: %w", err)
	}

	if err := vulnRows.Err(); err != nil {
		return nil, fmt.Errorf("iterating vulnerabilities: %w", err)
	}

	// Step 4: Load system packages
	spRows, err := tx.Query(`
		SELECT image_id, COALESCE(name, ''), COALESCE(version, ''),
		       COALESCE(package_type, '')
		FROM system_packages ORDER BY image_id, id`)
	if err != nil {
		return nil, fmt.Errorf("querying system packages: %w", err)
	}

	for spRows.Next() {
		var imageID int64
		var sp domain.SystemPackage

		if err := spRows.Scan(&imageID, &sp.Name, &sp.Version, &sp.PackageType); err != nil {
			_ = spRows.Close()
			return nil, fmt.Errorf("scanning system package: %w", err)
		}

		if img, ok := imageMap[imageID]; ok {
			img.SystemPackages = append(img.SystemPackages, sp)
		}
	}

	if err := spRows.Close(); err != nil {
		return nil, fmt.Errorf("closing system package rows: %w", err)
	}

	if err := spRows.Err(); err != nil {
		return nil, fmt.Errorf("iterating system packages: %w", err)
	}

	// Step 5: Load package managers
	pmRows, err := tx.Query(`
		SELECT image_id, COALESCE(name, ''), COALESCE(version, ''),
		       COALESCE(language, '')
		FROM package_managers ORDER BY image_id, id`)
	if err != nil {
		return nil, fmt.Errorf("querying package managers: %w", err)
	}

	for pmRows.Next() {
		var imageID int64
		var pm domain.PackageManager

		if err := pmRows.Scan(&imageID, &pm.Name, &pm.Version, &pm.Language); err != nil {
			_ = pmRows.Close()
			return nil, fmt.Errorf("scanning package manager: %w", err)
		}

		if img, ok := imageMap[imageID]; ok {
			img.PackageManagers = append(img.PackageManagers, pm)
		}
	}

	if err := pmRows.Close(); err != nil {
		return nil, fmt.Errorf("closing package manager rows: %w", err)
	}

	if err := pmRows.Err(); err != nil {
		return nil, fmt.Errorf("iterating package managers: %w", err)
	}

	// Assemble in original order
	results := make([]domain.ImageRecord, 0, len(imageOrder))
	for _, id := range imageOrder {
		results = append(results, *imageMap[id])
	}

	return results, nil
}

// GetDatabaseStats returns basic statistics about the database.
func (r *Repository) GetDatabaseStats() (map[string]int, error) {
	stats := make(map[string]int)
	tables := []string{"images", "languages", "vulnerabilities", "package_managers", "capabilities", "system_packages", "security_findings"}

	for _, t := range tables {
		var count int
		if err := r.db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", t)).Scan(&count); err != nil {
			return nil, fmt.Errorf("counting %s: %w", t, err)
		}

		stats[t] = count
	}

	return stats, nil
}
