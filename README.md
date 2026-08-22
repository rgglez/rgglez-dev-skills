# rgglez-dev-skills

A Claude Code plugin marketplace: development skills, one plugin per topic.

## Install

```
/plugin marketplace add rgglez/rgglez-dev-skills
```

Then install whichever plugins you want:

```
/plugin install go-dependency-direction@rgglez-dev-skills
```

`/plugin marketplace update rgglez-dev-skills` pulls later changes.

## Plugins

| Plugin | What it does |
| --- | --- |
| [`go-dependency-direction`](plugins/go-dependency-direction/) | Makes Claude follow one architectural rule when it writes Go: the import graph is the architecture. Caller-declared minimal interfaces, one-way package arrows. Ships a runnable example and an `AGENTS.md` variant for other agents. |

## Layout

```
.claude-plugin/marketplace.json     the catalog
plugins/<name>/
  .claude-plugin/plugin.json        that plugin's manifest
  skills/<name>/SKILL.md            the skill Claude loads
  README.md                         that plugin's docs
LICENSE
```

One plugin per skill, so people install only what they want. Adding another means a new directory under `plugins/` and one more entry in `marketplace.json`.

## Working on this repo

Validate before pushing — `--strict` turns warnings into errors and catches misspelled manifest fields:

```sh
claude plugin validate . --strict
claude plugin validate ./plugins/go-dependency-direction --strict
```

Test the marketplace locally before it goes public:

```
/plugin marketplace add /path/to/this/repo
/plugin install go-dependency-direction@rgglez-dev-skills
```

Version each plugin in its own `plugin.json` only. Setting `version` in `marketplace.json` as well is ignored without warning, and a stale manifest value silently wins.

## License

Copyright 2026 Rodolfo González González. [MIT](LICENSE).

Individual plugins credit their own sources; see each plugin's README.
