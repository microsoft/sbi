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
	"github.com/spf13/cobra"
)

var (
	flgDBPath         string
	flgConfigDir      string
	flgOutputPath     string
	flgTopN           int
	flgTopNJSON       int
	flgMaxTags        int
	flgComprehensive  bool
	flgNoCleanup      bool
	flgUpdateExisting bool
	flgVerbose        bool
	flgDebug          bool
	flgDetailed       bool
)

// NewRootCommand creates the root cobra command.
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "sbi",
		Short: "Scan container images and generate daily security recommendations by language",
		Long: `Scans container registries (MCR, Docker Hub, etc.) for base images,
analyzes them with Syft (SBOM) and Trivy (vulnerabilities), stores results in SQLite,
and generates a markdown report ranking images by security posture per language.`,
	}

	rootCmd.PersistentFlags().StringVar(&flgDBPath, "database", "azure_linux_images.db", "Path to SQLite database")
	rootCmd.PersistentFlags().StringVar(&flgConfigDir, "config-dir", "config", "Path to configuration directory")
	rootCmd.PersistentFlags().StringVar(&flgOutputPath, "output", "docs/daily_recommendations.md", "Path to output report file")
	rootCmd.PersistentFlags().IntVar(&flgTopN, "top-n", 10, "Number of top images per language per base OS (0 = all)")
	rootCmd.PersistentFlags().IntVar(&flgTopNJSON, "json-top-n", 20, "Number of top images per language per base OS in JSON report (0 = all)")
	rootCmd.PersistentFlags().BoolVarP(&flgVerbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVarP(&flgDebug, "debug", "d", false, "Enable debug output")
	rootCmd.PersistentFlags().BoolVar(&flgDetailed, "detailed", false, "Generate detailed per-image JSON report with packages and vulnerabilities")

	rootCmd.AddCommand(NewScanCommand())
	rootCmd.AddCommand(NewReportCommand())
	rootCmd.AddCommand(NewResetDBCommand())

	return rootCmd
}
