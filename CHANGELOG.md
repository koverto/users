# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][], and this project adheres to
[Semantic Versioning][].

## Unreleased

### Fixed

- Fixed nil pointer panic in `Users.Create` when setting `User.createdAt`

## v0.1.2 - 2020-02-17

### Changed

- Updated to github.com/koverto/uuid@v1.2.1

## v0.1.1 - 2020-02-17

### Changed

- Updated to github.com/koverto/uuid@v1.2.0

## v0.1.0 - 2020-02-17

### Added

- Protobuf API
- `Users.Create`
- `Users.Read`

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html
