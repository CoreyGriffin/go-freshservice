# Contributing

First of all, thanks for contributing!

This document provides some basic guidelines for contributing to this repository. To propose improvements, feel free to submit a pull request.

### Submitting Issues

Github issues are welcome, feel free to submit error reports and feature requests! Make sure to add enough details to explain your use case.

### Features & Fixes

- Fork the repository (https://github.com/CoreyGriffin/go-freshservice)
- Create your feature branch (git checkout -b my-new-feature)
- Commit your changes `git commit -am 'Add some feature'`
- Before pushing review the [PR guidelines below](#pull-request-guidelines)
- Push to the branch `git push origin my-new-feature`
- Create a new Pull Request
### Pull Request Guidelines

1. Ensure PR template is filled out and all relevant action items in the checklist are covered. PR titles should include one of the following: 
- `[FEATURE]`   - To be used with new features such as adding a new endpoint
- `[BUG FIX]`   - To be used with any bug fixes
- `[REFACTOR]`  - To be used when code change is specific to refactoring and no new features are added
- `[MISC]`      - To be used with documentation, dependency updates, etc.
  
2. Unless what you are doing is absolutely trivial, add unit tests. Good unit tests come in bundles, and usually test for both the expected and how one handles the unexpected case. To ensure your tests pass please run `go test -v ./...` or `make pr-prep` from the root of this repo. 
   
Smaller scoped PR's make the review process easier on everyone. Therefore, if you would prefer to submit your tests in a follow up PR please notate that in the description of the PR so the reviewer knows and be a good citizen.
