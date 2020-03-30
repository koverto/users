# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][], and this project adheres to
[Semantic Versioning][].

## Unreleased

## v1.0.1 - 2020-03-30

### Added

- `Client` helper struct

### Changed

- Updated to github.com/koverto/micro@v2.0.1

## v1.0.0 - 2020-03-05

No changes; bumping version to reflect feature-complete state.

## v0.2.0 - 2020-03-05

### Changed

- Updated to github.com/koverto/micro@v1.2.0
- Updated to github.com/koverto/uuid@v1.3.0

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
