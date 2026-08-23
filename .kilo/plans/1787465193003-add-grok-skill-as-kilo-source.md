# Plan: Make rgglez-dev-skills a Kilo (Grok) skills source via remote URLs

## Context
- Repo is a Claude Code plugin marketplace (`/.claude-plugin/marketplace.json`, `plugins/<name>/.claude-plugin/plugin.json`, `plugins/<name>/skills/<name>/SKILL.md`).
- Only skill today: `go-dependency-direction` (architectural import-direction rules for Go).
- SKILL.md already follows the shared Agent Skills spec (YAML frontmatter + Markdown) used by Kilo.
- Kilo loads skills on-demand via `description`; supports remote sources with `skills.urls` + `index.json` manifest.
- Repo already ships `portable/AGENTS.md` (trimmed rules) for project-level use in Grok/Codex/etc.
- User goal: make *this repo itself* a usable Kilo marketplace source (not just docs or central-repo PR).

## Key Decisions (resolved)
- **Distribution**: enable via Kilo remote skills (`kilo.jsonc` → `skills.urls`). (Central Kilo-Org/kilo-marketplace submission is out of scope.)
- **Layout**: add top-level `skills/` (parallel to `plugins/`) for clean Kilo URLs.
- **Content**: create *adapted* Kilo variant (strip Claude-only persistence/off-switch/session text; keep core rules + examples identical). Port manually like the existing `portable/AGENTS.md`.
- **Versioning**: reuse the version string from `plugins/go-dependency-direction/.claude-plugin/plugin.json` (currently 0.2.0) in `index.json`.
- **Claude side untouched**: `plugins/` layout, manifests, and `claude plugin validate` remain exactly as-is.
- Out of scope: Kilo "Agent" entries, MCP, auto-port scripts, changes to central kilo-marketplace repo, bundling example/ into the skill.

## Implementation Task List (ordered, minimal)

1. Create the Kilo skills tree at repo root:
   ```
   mkdir -p skills/go-dependency-direction
   ```

2. Create `skills/go-dependency-direction/SKILL.md` (adapted):
   - Frontmatter (exact):
     ```yaml
     ---
     name: go-dependency-direction
     description: Applies dependency-direction rules when writing or reviewing Go code — consumer-defined minimal interfaces, no domain→infrastructure imports, one-way package arrows. Load whenever writing, refactoring, or reviewing Go packages, designing package layout, defining interfaces, or wiring repository/service/handler layers.
     license: MIT
     ---
     ```
   - Body: copy the architecture rules, examples, checks, "Cases this does not cover", and credit from the current `plugins/go-dependency-direction/skills/go-dependency-direction/SKILL.md` (or the trimmed `portable/AGENTS.md`).
   - Adapt *only* the activation/persistence section:
     - Replace Claude "Persistence / session / stop go-dd" wording with Kilo-friendly equivalent (on-demand via description; force with name mention; "stop go-dependency-direction" / "normal mode" / "ignora las reglas..."; mention `/reload` after edits).
     - Remove or generalize any "Claude Code skill", `~/.claude/skills/`, or `/plugin` language.
     - Keep Spanish synonyms if present.
   - Ensure the file is < ~500 lines; no shell `!`` commands for now.
   - `name` must exactly match directory name.

3. Create `skills/index.json` (at `skills/index.json`):
   ```json
   {
     "skills": [
       {
         "name": "go-dependency-direction",
         "version": "0.2.0",
         "files": ["SKILL.md"]
       }
     ]
   }
   ```
   - `version` must match the string in `plugins/go-dependency-direction/.claude-plugin/plugin.json`.

4. Update root `README.md`:
   - Add Kilo/Grok usage example under "Install" or a new "Kilo / Grok" section.
   - Show:
     ```jsonc
     // in kilo.jsonc (project or global)
     {
       "skills": {
         "urls": ["https://raw.githubusercontent.com/rgglez/rgglez-dev-skills/main/skills"]
       }
     }
     ```
   - After change: `/reload` or new session.
   - Note: this gives the *adapted* on-demand skill. For project-wide rules still copy `portable/AGENTS.md` to repo root as `AGENTS.md`.
   - Update the Plugins table or "Layout" if needed to mention the new `skills/` tree.
   - Add a short "Kilo skills source" sentence in the intro.

5. Update `plugins/go-dependency-direction/README.md`:
   - In the existing "## Other agents (Codex, Grok, …)" section, add Kilo-specific paragraph:
     - Remote config example (pointing at `/skills`).
     - Distinction: this `SKILL.md` (Kilo-adapted, on-demand) vs. `portable/AGENTS.md` (always-on project rules).
     - "Keeping the two [Claude SKILL + Kilo SKILL] in sync is manual. Edit the Claude `SKILL.md` (source of the rules) first, then port the change to `skills/go-dependency-direction/SKILL.md` (only adapt the activation text)."
   - Keep all existing Claude install + portable instructions.

6. (Minimal) Update the "Working on this repo" validation section in root README:
   - Keep the two `claude plugin validate` commands.
   - Add: "For the Kilo side: ensure `skills/go-dependency-direction/SKILL.md` frontmatter has `name` + `description`, `name` matches its directory, and `skills/index.json` version matches the plugin manifest."

7. No other files (no changes to `.claude-plugin/`, no `kilo.jsonc` in this repo, no symlinks, no example/ duplication).

## Risks & Mitigations
- Content/rules drift between Claude SKILL.md and Kilo variant → explicit "edit Claude first, then port" note in both READMEs (already the pattern for portable/).
- Version not bumped in lockstep → document "when you bump the plugin version, also update index.json".
- Remote cache in Kilo → version field forces refresh on next discovery.
- Duplicate skill if user also installs the Claude plugin + has `.claude/skills/` compat enabled in Kilo → Kilo priority rules apply; document the native remote as preferred for pure-Kilo users.
- Windows symlink issues avoided (we copy the adapted content).

## Validation Steps (for the impl agent + user)
- `claude plugin validate . --strict` and `claude plugin validate ./plugins/go-dependency-direction --strict` still pass (new `skills/` dir is outside any plugin).
- `skills/go-dependency-direction/SKILL.md`:
  - Valid YAML frontmatter with required `name` + `description`.
  - `name: go-dependency-direction` exactly matches parent dir.
  - No Claude-only load mechanics left in the body.
- `skills/index.json` is valid JSON, version exactly matches the one in the plugin manifest, `files` lists only what exists.
- Grep the two READMEs for "kilo", "grok", "skills.urls", "go-dependency-direction" (Kilo path).
- (Manual) Conceptually verify a raw URL would serve:
  - `.../skills/index.json`
  - `.../skills/go-dependency-direction/SKILL.md`
- (If Kilo available) In a test workspace set the `skills.urls`, start session, give a Go layering task, confirm the skill description appears and rules are followed; "stop go-dependency-direction" deactivates for the turn.

## Open / Out of Scope (explicitly)
- Submitting `skills/go-dependency-direction/` to https://github.com/Kilo-Org/kilo-marketplace (would be a separate PR; this repo becomes an independent remote source).
- Adding Kilo "Agent" items (e.g. from `portable/AGENTS.md`) or MCP.
- Bundling extra files (`example/`, `references/`) in the index.
- CI / script to keep the adapted file in sync.
- Changing how the Claude plugin is laid out or validated.
- Support for Kilo's built-in Marketplace UI (requires the central repo).

Plan is implementation-ready. The tasks above + the adaptation rules for the body are sufficient for another agent to execute without further design decisions.
