# Managing Upstream Subtree

This document describes how to update and reconcile the backpack git subtree.

## Detecting Drift

To see if there are changes between the current application's `.backpack`
directory and the upstream backpack repo, we can run the following commands:

```sh
git remote add -f backpack git@github.com:yanglinz/backpack.git
git fetch backpack
git diff backpack/master master:.backpack/
```

## Updating Backpack

To update backpack subtree to the latest version, we can run the following
commands:

```sh
git remote add -f backpack git@github.com:yanglinz/backpack.git
git fetch backpack
git subtree pull --prefix=.backpack backpack master --squash
```

> Note that because git subtree primarily rely on commit message metadata for
> diff-ing when pulling and pushing, it's important to avoid anything that may
> modify commit messages, like github's "squash" feature. When in doubt, keep
> subtree modifications in its own atomic commit and push directly to the main
> line `master` branch.

## Pushing to Upstream Backpack

There are cases where we may have made changes inside the `.backpack` directory.
We should reconcile these changes with the upstream repo.

```sh
git remote add -f backpack git@github.com:yanglinz/backpack.git
git fetch backpack
git subtree push --prefix=.backpack backpack <some-branch-name>
```

This will create `<some-brannch-name>` in the backpack repo.
