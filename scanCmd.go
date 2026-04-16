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

package main

import (
	"github.com/microsoft/sbi/pkg/domain"
	"github.com/microsoft/sbi/pkg/infrastructure/database"
	"github.com/microsoft/sbi/pkg/usecase"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewScanCommand creates the `scan` subcommand.
func NewScanCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Scan container registries and analyze images",
		Long: `Discover image tags from configured registries, pull images,
analyze with Syft and Trivy, and store results in the database.`,
		RunE: runScan,
	}

	cmd.Flags().IntVar(&flgMaxTags, "max-tags", 5, "Maximum tags per repository (0 = all)")
	cmd.Flags().BoolVar(&flgComprehensive, "comprehensive", false, "Enable comprehensive scanning (secrets + misconfigs)")
	cmd.Flags().BoolVar(&flgNoCleanup, "no-cleanup", false, "Keep Docker images after scanning")
	cmd.Flags().BoolVar(&flgUpdateExisting, "update-existing", false, "Rescan existing images")

	return cmd
}

func runScan(cmd *cobra.Command, _ []string) error {
	setLogLevel()

	db, err := database.OpenDB(flgDBPath)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	if err := database.CreateTables(db); err != nil {
		return err
	}

	config := domain.ScanConfig{
		DBPath:            flgDBPath,
		ConfigDir:         flgConfigDir,
		OutputPath:        flgOutputPath,
		TopN:              flgTopN,
		TopNJSON:          flgTopNJSON,
		MaxTagsPerRepo:    flgMaxTags,
		ComprehensiveScan: flgComprehensive,
		CleanupImages:     !flgNoCleanup,
		UpdateExisting:    flgUpdateExisting,
		Verbose:           flgVerbose,
	}

	repo := database.NewRepository(db)
	pipeline, err := usecase.NewPipeline(config, repo)
	if err != nil {
		return err
	}

	log.Info("Starting scan...")

	if err := pipeline.ScanAll(); err != nil {
		return err
	}

	log.Info("Scan complete. Generating report...")

	return pipeline.GenerateReport()
}

func setLogLevel() {
	if flgDebug {
		log.SetLevel(log.DebugLevel)
	} else if flgVerbose {
		log.SetLevel(log.InfoLevel)
	}
}
