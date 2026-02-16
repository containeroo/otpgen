# otpgen

`otpgen` is a small CLI for generating TOTP (RFC 6238) one-time codes from a shared secret.

## Install

### Homebrew (macOS/Linux)

```bash
brew tap containeroo/tap
brew install --cask otpgen
```

### Build from source

Requirements:
- Go `1.26.0`

```bash
git clone git@github.com:containeroo/otpgen.git
cd otpgen
go build -o otpgen .
```

## Setup and Secret Input

`otpgen` expects one positional argument: a TOTP secret (typically Base32) from your identity provider.

Examples:
- GitHub 2FA secret
- Google Workspace account secret
- Any RFC 6238-compatible provider secret

## Usage

Generate a code:

```bash
otpgen JBSWY3DPEHPK3PXP
```

Generate with a quoted secret:

```bash
otpgen "JBSWY3DPEHPK3PXP"
```

Show version:

```bash
otpgen version
```

Generate completion scripts:

```bash
otpgen completion bash
otpgen completion zsh
otpgen completion powershell
```

## Command Behavior

- Input validation:
  - Exactly one positional argument is required.
  - Whitespace-only secrets are rejected.
- Errors:
  - All command errors are printed to `stderr`.
  - Exit status is non-zero on failure.
- Output:
  - On success, only the generated OTP code is written to `stdout`.

## Security Notes

- Treat TOTP secrets as credentials:
  - Do not paste secrets into shared terminals or shell history where avoidable.
  - Prefer environment-injected or ephemeral input paths in automated environments.
- Output is plain text:
  - Anyone with terminal/log access can read generated codes.
  - Avoid redirecting OTP output into persistent logs.
