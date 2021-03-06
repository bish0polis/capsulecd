package engine

import (
	"github.com/analogj/capsulecd/pkg/config"
	"github.com/analogj/capsulecd/pkg/pipeline"
	"github.com/analogj/capsulecd/pkg/scm"
)

// Create mock using:
// mockgen -source=pkg/engine/interface.go -destination=pkg/engine/mock/mock_engine.go
type Interface interface {
	Init(pipelineData *pipeline.Data, config config.Interface, sourceScm scm.Interface) error

	GetCurrentMetadata() interface{}
	GetNextMetadata() interface{}

	// Validate that required executables are available for the following build/test/package/etc steps
	ValidateTools() error

	// Assemble the package contents
	// Validate that any required files (like dependency management files) exist
	// Create any recommended optional/missing files we can in the structure. (.gitignore, etc)
	// Read & Bump the version in the metadata file(s)
	// CAN NOT override
	// MUST set CurrentMetadata
	// MUST set NextMetadata
	// REQUIRES pipelineData.GitLocalPath
	AssembleStep() error

	// Compile the source for this package (if required)
	// CAN override
	// USES engine_disable_compile
	// USES engine_cmd_compile
	// REQUIRES pipelineData.GitLocalPath
	CompileStep() error

	// Validate code syntax & execute test runner
	// CAN override
	// Run linter
	// Run unit tests
	// Generate coverage reports
	// USES engine_disable_test
	// USES engine_disable_lint
	// USES engine_disable_security_check
	// USES engine_enable_code_mutation - allows CapsuleCD to modify code using linting tools (only available on some systems)
	// USES engine_cmd_lint
	// USES engine_cmd_test
	// USES engine_cmd_security_check
	TestStep() error

	// Commit any local changes and create a git tag. Nothing should be pushed to remote repository yet.
	// Make sure you remove any unnecessary files from the repo before making the commit
	// CAN NOT override
	// MUST set ReleaseCommit
	// MUST set ReleaseVersion
	// REQUIRES pipelineData.GitLocalPath
	// REQUIRES NextMetadata
	// USES mgr_keep_lock_file
	PackageStep() error

}
