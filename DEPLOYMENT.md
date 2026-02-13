# HttpYum Deployment Guide

This document provides step-by-step instructions for deploying new versions of HttpYum. The project uses an automated CI/CD pipeline with conventional commits for seamless releases.

## Overview

HttpYum uses a sophisticated automated deployment system with:
- **Conventional commit-based** automatic versioning
- **Cross-platform releases** (Linux, macOS, Windows - amd64 & arm64)
- **GitHub Actions** for CI/CD automation
- **GitHub Pages** for documentation deployment
- **Professional installation scripts** for end users

## Deployment Methods

### Method 1: Automatic Deployment (Recommended)

The project automatically creates releases based on conventional commit messages. No manual intervention required!

#### Step 1: Write Code Changes
Make your code changes as usual using standard development practices.

#### Step 2: Use Conventional Commits
When committing your changes, use conventional commit format:

```bash
# For new features (triggers minor version bump)
git commit -m "feat: add support for GraphQL requests"
git commit -m "feat(parser): implement multipart form data parsing"

# For bug fixes (triggers patch version bump)
git commit -m "fix: resolve HTTP timeout handling issue"
git commit -m "fix(ui): correct viewport rendering on small terminals"

# For performance improvements (triggers patch version bump)
git commit -m "perf: optimize request parsing for large files"

# For breaking changes (triggers major version bump)
git commit -m "feat!: redesign configuration file format"
git commit -m "feat: change default timeout to 30s

BREAKING CHANGE: Default request timeout changed from 10s to 30s"

# For non-release commits (no version bump)
git commit -m "docs: update README with new examples"
git commit -m "chore: update dependencies"
git commit -m "refactor: simplify request validation logic"
```

#### Step 3: Push to Main Branch
```bash
git push origin main
```

#### Step 4: Automatic Release Process
Once pushed, the automation will:

1. **Analyze commits** since last release using conventional commit format
2. **Calculate new version** (major.minor.patch) based on commit types
3. **Create and push new tag** (if version bump needed)
4. **Trigger release workflow** automatically
5. **Build cross-platform binaries** (6 platforms total)
6. **Create GitHub release** with auto-generated release notes
7. **Deploy documentation** to GitHub Pages

#### Step 5: Verify Release
- Check [GitHub Releases](https://github.com/aritra1999/httpyum/releases) for your new release
- Verify all platform binaries are present
- Test installation using the provided scripts

### Method 2: Manual Release (For Special Cases)

If you need to create a release manually or override the automation:

#### Step 1: Ensure Code Quality
```bash
# Run all quality checks
make all

# Or run individual checks
make fmt && make lint && make test
```

#### Step 2: Manual Version Bump (Optional)
```bash
# Manually bump version if needed
make bump-version VERSION=1.2.3
```

#### Step 3: Create Release via GitHub CLI
```bash
# Create release manually
make release VERSION=v1.2.3
```

## Understanding the Deployment Pipeline

### Workflow Files
- **`.github/workflows/test.yml`** - Runs tests on every PR/push
- **`.github/workflows/auto-release.yml`** - Analyzes commits and creates tags
- **`.github/workflows/release.yml`** - Builds binaries and creates GitHub releases
- **`.github/workflows/deploy.yml`** - Deploys documentation site

### Version Calculation Rules
- **Major (x.0.0)**: `feat!:`, `fix!:`, or `BREAKING CHANGE:` in commit body
- **Minor (x.y.0)**: `feat:` commits (new features)
- **Patch (x.y.z)**: `fix:`, `perf:`, `refactor:` commits
- **No Release**: `docs:`, `chore:`, `test:`, `ci:` commits

### Build Targets
The release builds for these platforms:
- **Linux**: amd64, arm64
- **macOS**: amd64 (Intel), arm64 (Apple Silicon)
- **Windows**: amd64, arm64

## Installation Methods for End Users

Your releases will be available through multiple installation methods:

### Quick Install (Recommended for Users)
```bash
curl -fsSL https://raw.githubusercontent.com/aritra1999/httpyum/main/scripts/install.sh | bash
```

### Manual Download
Users can download binaries directly from the [GitHub Releases](https://github.com/aritra1999/httpyum/releases) page.

### Update Existing Installation
```bash
curl -fsSL https://raw.githubusercontent.com/aritra1999/httpyum/main/scripts/update.sh | bash
```

### From Source
```bash
go install github.com/aritra1999/httpyum/cmd/httpyum@latest
```

## Troubleshooting Deployments

### Release Not Created
- Check if commits use proper conventional commit format
- Verify commits include changes that warrant a release (not just docs/chore)
- Check GitHub Actions logs for any failures

### Build Failures
- Ensure all tests pass locally: `make test`
- Check if code compiles on all platforms: `make release-local`
- Review GitHub Actions workflow logs

### Documentation Not Deployed
- Check if changes exist in `/site` directory
- Verify SvelteKit build succeeds locally
- Review deploy.yml workflow logs

## Best Practices

### Commit Messages
- Use clear, descriptive commit messages
- Follow conventional commit format consistently
- Include scope when relevant: `feat(parser): add new feature`
- Explain the "why" in commit body for complex changes

### Before Deployment
```bash
# Always run before pushing
make fmt          # Format code
make lint         # Run linters
make test         # Run all tests
make test-race    # Test with race detector
```

### Release Verification
After each release:
1. Test installation script on different platforms
2. Verify binary functionality
3. Check release notes accuracy
4. Test documentation site updates

## Emergency Procedures

### Rollback a Release
```bash
# Delete problematic tag and release
git tag -d v1.2.3
git push origin :refs/tags/v1.2.3

# Delete GitHub release manually from web interface
# Create hotfix with proper version bump
```

### Hotfix Release
```bash
# Create hotfix branch from problematic release
git checkout v1.2.3
git checkout -b hotfix/critical-fix

# Make minimal fix
# Commit with conventional format
git commit -m "fix: resolve critical security issue"

# Push to main to trigger automated release
git checkout main
git cherry-pick hotfix/critical-fix
git push origin main
```

## Monitoring Deployments

- **GitHub Actions**: Monitor workflow runs in the Actions tab
- **Codecov**: Track test coverage changes
- **GitHub Releases**: Verify release artifacts and download counts
- **Documentation Site**: Ensure docs updates deploy correctly

## Support

For deployment issues:
1. Check GitHub Actions workflow logs
2. Review this deployment guide
3. Check project's GitHub Issues for similar problems
4. Create new issue with deployment logs if needed

---

**Remember**: The automated deployment system handles most scenarios. Manual intervention should only be needed for special cases or troubleshooting.