# go-dependency-direction

A Claude Code skill that makes Claude follow one architectural rule when it writes Go: **the import graph is the architecture**. Folder names are not.

Distilled from [*Forget clean architecture — Master dependency direction in Go*](https://blog.stackademic.com/forget-clean-architecture-master-dependency-direction-in-go-2c875193e940) (Yaninyz witty, Stackademic, Jul 2026).

## Contents

```
skills/go-dependency-direction/SKILL.md   the skill itself (Claude Code)
portable/AGENTS.md                        trimmed variant for other agents — see below
example/                                  a runnable Go program that obeys it
```

## What the skill enforces

- **The caller declares the interface.** The package that *invokes* a dependency defines it, listing only the methods it invokes. Not beside the implementation.
- **Small interfaces, usually unexported.** One method beats five. Accept interfaces, return structs.
- **Two checks before any import.** Which way does the new arrow run? Is one already running the opposite way? Plus, per package: strip the project to this package alone — does it still build?
- **Layered designs are only as good as their arrows.** `handler → service → repository`, no return path. It names the two shortcuts that erode it: the service reaching for the handler's error mapping, the repository borrowing the service's DTO.
- **Check it, don't trust it.** `go list -f '{{.ImportPath}} -> {{.Imports}}' ./...` prints the real graph, one line per package.
- **Cases it does not cover.** `main`/wiring imports freely — that is its purpose. No interface for one implementation with no test fake. Reuse stdlib interfaces (`io.Writer`, `fmt.Stringer`) when they already match the shape you need — they are caller-shaped already, and they compose with the rest of the ecosystem.

## Install

As a plugin, from the marketplace this repo publishes:

```
/plugin marketplace add rgglez/rgglez-dev-skills
/plugin install go-dependency-direction@rgglez-dev-skills
```

Or drop the skill file in by hand — the directory name becomes the skill name:

```sh
mkdir -p ~/.claude/skills/go-dependency-direction
cp skills/go-dependency-direction/SKILL.md ~/.claude/skills/go-dependency-direction/
```

Project-scoped instead of global: use `.claude/skills/` inside the repo you want it in.

## How it triggers

No slash command. The `description` in the frontmatter is the trigger — Claude loads the skill on its own when a task involves writing, refactoring, or reviewing Go, designing package layout, defining interfaces, or wiring repository/service/handler layers.

If it fails to load on a task where it should, widen the `description` with more trigger terms. Leave the body alone: the body is what Claude follows *after* loading, the description is what gets it loaded.

## Checking it works

Ask Claude for anything where one package needs a collaborator — storage, an HTTP client for a third-party API, a notifier, a cache, a clock, a metrics sink, a queue publisher. Under the skill it declares a minimal interface in the *consuming* package. Without it, it tends to export a fat interface from the implementation package and make every caller import it.

Storage is just the most common instance; the tell is the same everywhere: which side owns the interface, and how many methods it has.

[`example/`](example/) is the reference outcome: `invoicing` declares a one-method `sender`, `memmail` satisfies it without knowing it exists, and only `main` imports both. Neither package appears in the other's import list.

## Other agents (Codex, Grok, …)

The rules are plain markdown and port anywhere. What does **not** port is the loading machinery: the YAML frontmatter, the `~/.claude/skills/` location, and the conditional auto-load that keeps the body out of context until a Go task shows up.

[`portable/AGENTS.md`](portable/AGENTS.md) is the trimmed variant — same rules, one code pair instead of the antipattern/pattern pair, no frontmatter. Copy it into the root of a Go repo:

```sh
cp portable/AGENTS.md /path/to/your-go-repo/AGENTS.md
```

Codex reads `AGENTS.md` from the repo root. For Grok it depends on the client — if it honors `AGENTS.md` or a rules file, same move; otherwise paste the contents as a system instruction. Check the current docs for either, the conventions move.

Two caveats:

- **It is always loaded**, not conditional. It costs context on every turn, `git status` included. That is why the portable version is trimmed, and why it belongs only in repos that are actually Go.
- **It merges with whatever else the repo's `AGENTS.md` says.** If one already exists, append these rules as a section instead of overwriting it.

Keeping the two in sync is manual. `SKILL.md` is the source; edit it first, then port the change.

## License

Copyright 2026 Rodolfo González González.

[MIT](../../LICENSE) — covers the skill text, the portable variant, and the example program, all written for this repo.

The underlying ideas come from the Stackademic article linked above; ideas are not owned, but its author's prose is. Every rule here is worded independently and the code examples are original, so nothing in this repo reproduces the article. The PDF used as source material is `.gitignore`d and must not be committed or redistributed.
