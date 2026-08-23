# rgglez-dev-skills

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
![GitHub all releases](https://img.shields.io/github/downloads/rgglez/rgglez-dev-skills/total)
![GitHub issues](https://img.shields.io/github/issues/rgglez/rgglez-dev-skills)
![GitHub commit activity](https://img.shields.io/github/commit-activity/y/rgglez/rgglez-dev-skills)
[![GitHub release](https://img.shields.io/github/release/rgglez/rgglez-dev-skills.svg)](https://github.com/rgglez/rgglez-dev-skills/releases/)
![GitHub stars](https://img.shields.io/github/stars/rgglez/rgglez-dev-skills?style=social)
![GitHub forks](https://img.shields.io/github/forks/rgglez/rgglez-dev-skills?style=social)

A marketplace for development skills.

- Claude Code: plugin marketplace (`.claude-plugin/`).
- Kilo (Grok): remote skills source via `skills.urls` + `index.json` (Agent Skills standard).

## Install

```
/plugin marketplace add rgglez/rgglez-dev-skills
```

Then install whichever plugins you want:

```
/plugin install go-dependency-direction@rgglez-dev-skills
```

`/plugin marketplace update rgglez-dev-skills` pulls later changes.

## Kilo / Grok

Add as a remote skills source (no UI marketplace registration needed):

```jsonc
// kilo.jsonc (project or global)
{
  "skills": {
    "urls": ["https://raw.githubusercontent.com/rgglez/rgglez-dev-skills/main/skills"]
  }
}
```

Then `/reload` or start a new session. The skill loads on-demand when the task matches its description.

For always-on project rules instead, copy `portable/AGENTS.md` to your repo root as `AGENTS.md`.

## Plugins

| Plugin | What it does |
| --- | --- |
| [`go-dependency-direction`](plugins/go-dependency-direction/) | Makes the agent follow one architectural rule when it writes Go: the import graph is the architecture. Caller-declared minimal interfaces, one-way package arrows. Ships a runnable example and an `AGENTS.md` variant for other agents. Available for Claude (plugin) and Kilo/Grok (remote skill). |

## Layout

```
.claude-plugin/marketplace.json     the catalog (Claude)
plugins/<name>/
  .claude-plugin/plugin.json        that plugin's manifest
  skills/<name>/SKILL.md            the skill Claude loads
  README.md                         that plugin's docs
skills/                             Kilo remote skills source
  <name>/
    SKILL.md
  index.json
LICENSE
```

One plugin per skill (Claude). Kilo skills are flat under `skills/`. Adding another Claude plugin means a new dir under `plugins/` + entry in `marketplace.json`. Kilo additions go under `skills/`.

## Working on this repo

Validate before pushing — `--strict` turns warnings into errors and catches misspelled manifest fields:

```sh
claude plugin validate . --strict
claude plugin validate ./plugins/go-dependency-direction --strict
```

For the Kilo side, ensure `skills/<name>/SKILL.md` has valid frontmatter (`name` + `description`, `name` matches dir) and `skills/index.json` version matches the corresponding `plugin.json`.

Test the marketplace locally before it goes public:

```
 /plugin marketplace add /path/to/this/repo
 /plugin install go-dependency-direction@rgglez-dev-skills
```

For Kilo, point `skills.urls` at the `skills/` subdir and start a session (or `/reload`).

Version each plugin in its own `plugin.json` only. Setting `version` in `marketplace.json` as well is ignored without warning, and a stale manifest value silently wins. Keep the `skills/index.json` version in sync with the plugin manifest.

## License

Copyright 2026 Rodolfo González González. [MIT](LICENSE).

Individual plugins credit their own sources; see each plugin's README.
