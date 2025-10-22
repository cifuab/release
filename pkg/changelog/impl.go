/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package changelog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/blang/semver/v4"
	"github.com/sirupsen/logrus"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"

	"sigs.k8s.io/mdtoc/pkg/mdtoc"
	"sigs.k8s.io/release-sdk/git"
	"sigs.k8s.io/release-sdk/github"
	"sigs.k8s.io/release-sdk/object"
	"sigs.k8s.io/release-utils/helpers"
	"sigs.k8s.io/release-utils/http"

	"k8s.io/release/pkg/cve"
	"k8s.io/release/pkg/notes"
	"k8s.io/release/pkg/notes/document"
	"k8s.io/release/pkg/notes/options"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . impl
//go:generate /usr/bin/env bash -c "cat ../../hack/boilerplate/boilerplate.generatego.txt changelogfakes/fake_impl.go > changelogfakes/_fake_impl.go && mv changelogfakes/_fake_impl.go changelogfakes/fake_impl.go"

type impl interface {
	// Used in `Run()`
	TagStringToSemver(tag string) (semver.Version, error)
	OpenRepo(repoPath string) (*git.Repo, error)
	CurrentBranch(repo *git.Repo) (branch string, err error)
	RevParse(repo *git.Repo, rev string) (string, error)
	RevParseTag(repo *git.Repo, rev string) (string, error)
	CreateDownloadsTable(
		w io.Writer, bucket, tars, images, prevTag, newTag string,
	) error
	LatestGitHubTagsPerBranch() (github.TagsPerBranch, error)
	GenerateTOC(markdown string) (string, error)
	DependencyChanges(from, to string) (string, error)
	Checkout(repo *git.Repo, rev string, args ...string) error

	// Used in `generateReleaseNotes()`
	ValidateAndFinish(opts *options.Options) error
	GatherReleaseNotes(opts *options.Options) (*notes.ReleaseNotes, error)
	NewDocument(
		releaseNotes *notes.ReleaseNotes, previousRev, currentRev string,
	) (*document.Document, error)
	RenderMarkdownTemplate(
		document *document.Document, bucket, tars, images, templateSpec string,
	) (string, error)

	// Used in `writeMarkdown()`
	RepoDir(repo *git.Repo) string
	WriteFile(filename string, data []byte, perm os.FileMode) error
	Stat(name string) (os.FileInfo, error)
	ReadFile(filename string) ([]byte, error)

	// Used in `writeHTML()`
	MarkdownToHTML(
		markdown string, writer io.Writer, opts ...parser.ParseOption,
	) error
	ParseHTMLTemplate(text string) (*template.Template, error)
	TemplateExecute(
		tpl *template.Template, wr io.Writer, data any,
	) error
	Abs(path string) (string, error)

	// Used in `lookupRemoteReleaseNotes()`
	GetURLResponse(url string) (string, error)

	// Used in `commitChanges()`
	Add(repo *git.Repo, filename string) error
	Commit(repo *git.Repo, msg string) error
	Rm(repo *git.Repo, force bool, files ...string) error
	CloneCVEData() (cveDir string, err error)
}

type defaultImpl struct{}

func (*defaultImpl) TagStringToSemver(tag string) (semver.Version, error) {
	return helpers.TagStringToSemver(tag)
}

func (*defaultImpl) OpenRepo(repoPath string) (*git.Repo, error) {
	return git.OpenRepo(repoPath)
}

func (*defaultImpl) CurrentBranch(repo *git.Repo) (string, error) {
	return repo.CurrentBranch()
}

func (*defaultImpl) RevParse(repo *git.Repo, rev string) (string, error) {
	return repo.RevParse(rev)
}

func (*defaultImpl) RevParseTag(repo *git.Repo, rev string) (string, error) {
	return repo.RevParseTag(rev)
}

func (*defaultImpl) CreateDownloadsTable(
	w io.Writer, bucket, tars, images, prevTag, newTag string,
) error {
	return document.CreateDownloadsTable(w, bucket, tars, images, prevTag, newTag)
}

func (*defaultImpl) LatestGitHubTagsPerBranch() (github.TagsPerBranch, error) {
	return github.New().LatestGitHubTagsPerBranch()
}

func (*defaultImpl) GenerateTOC(markdown string) (string, error) {
	return mdtoc.GenerateTOC([]byte(markdown), mdtoc.Options{
		Dryrun:     false,
		SkipPrefix: false,
		MaxDepth:   mdtoc.MaxHeaderDepth,
	})
}

func (*defaultImpl) DependencyChanges(from, to string) (string, error) {
	return notes.NewDependencies().Changes(from, to)
}

func (*defaultImpl) Checkout(repo *git.Repo, rev string, args ...string) error {
	return repo.Checkout(rev, args...)
}

func (*defaultImpl) ValidateAndFinish(opts *options.Options) error {
	return opts.ValidateAndFinish()
}

func (*defaultImpl) GatherReleaseNotes(
	opts *options.Options,
) (*notes.ReleaseNotes, error) {
	return notes.GatherReleaseNotes(opts)
}

func (*defaultImpl) NewDocument(
	releaseNotes *notes.ReleaseNotes, previousRev, currentRev string,
) (*document.Document, error) {
	return document.New(releaseNotes, previousRev, currentRev)
}

func (*defaultImpl) RenderMarkdownTemplate(
	doc *document.Document, bucket, tars, images, templateSpec string,
) (string, error) {
	return doc.RenderMarkdownTemplate(bucket, tars, images, templateSpec)
}

func (*defaultImpl) RepoDir(repo *git.Repo) string {
	return repo.Dir()
}

func (*defaultImpl) WriteFile(
	filename string, data []byte, perm os.FileMode,
) error {
	return os.WriteFile(filename, data, perm)
}

func (*defaultImpl) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (*defaultImpl) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (*defaultImpl) MarkdownToHTML(
	markdown string, writer io.Writer, opts ...parser.ParseOption,
) error {
	return goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	).Convert([]byte(markdown), writer, opts...)
}

func (*defaultImpl) ParseHTMLTemplate(text string) (*template.Template, error) {
	return template.New("html").Parse(text)
}

func (*defaultImpl) TemplateExecute(
	tpl *template.Template, wr io.Writer, data any,
) error {
	return tpl.Execute(wr, data)
}

func (*defaultImpl) Abs(path string) (string, error) {
	return filepath.Abs(path)
}

func (*defaultImpl) GetURLResponse(url string) (string, error) {
	content, err := http.NewAgent().Get(url)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func (*defaultImpl) Add(repo *git.Repo, filename string) error {
	return repo.Add(filename)
}

func (*defaultImpl) Commit(repo *git.Repo, msg string) error {
	return repo.Commit(msg)
}

func (*defaultImpl) Rm(repo *git.Repo, force bool, files ...string) error {
	return repo.Rm(force, files...)
}

// CloneCVEData copies the CVE data maps from the release bucket.
func (*defaultImpl) CloneCVEData() (cveDir string, err error) {
	tmpdir, err := os.MkdirTemp(os.TempDir(), "cve-maps-")
	if err != nil {
		return "", fmt.Errorf("creating temporary dir for CVE data: %w", err)
	}

	// Create a new GCS client
	gcs := object.NewGCS()

	remoteSrc, err := gcs.NormalizePath(object.GcsPrefix + filepath.Join(cve.Bucket, cve.Directory))
	if err != nil {
		return "", fmt.Errorf("normalizing cve bucket path: %w", err)
	}

	bucketExists, err := gcs.PathExists(remoteSrc)
	if err != nil {
		return "", fmt.Errorf("checking if CVE directory exists: %w", err)
	}

	if !bucketExists {
		logrus.Warnf("CVE data maps not found in bucket location: %s", remoteSrc)

		return "", nil
	}

	// Rsync the CVE files to the temp dir
	if err := gcs.RsyncRecursive(remoteSrc, tmpdir); err != nil {
		return "", fmt.Errorf("copying release directory to bucket: %w", err)
	}

	logrus.Infof("Successfully synchronized CVE data maps from %s", remoteSrc)

	return tmpdir, nil
}
