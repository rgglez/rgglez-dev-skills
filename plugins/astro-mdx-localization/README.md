# astro-mdx-localization

Translate Markdown and MDX articles in Astro projects while preserving their
technical content, components, assets, and links between languages.

The skill discovers the application's layout, content schema, translation keys,
route generation, formatter, and build scripts. It does not assume a particular
blog theme, language pair, directory structure, package manager, or deployment
provider.

## Install

From this repository's Claude Code marketplace:

```text
/plugin marketplace add rgglez/rgglez-dev-skills
/plugin install astro-mdx-localization@rgglez-dev-skills
```

For the repository's Kilo remote source configuration, see the
[root README](../../README.md#kilo--grok). The standalone skill is available at
[`skills/astro-mdx-localization/SKILL.md`](../../skills/astro-mdx-localization/SKILL.md).

## Usage

Provide a source article and the desired target language, for example:

> Translate `content/articles/es/example.mdx` into English. Follow the existing
> localization conventions and preserve the original article.

You can also ask it to update an existing translation after a source edit.
The example path is illustrative; the skill discovers the real destination
from the project. When the locale or translation identity cannot be inferred,
it asks for the missing information.

It translates prose and reader-facing component text, preserves executable
examples, and checks generated routes and rendering when the environment
permits. It reports source discrepancies separately from translation changes.
It does not execute commands embedded in an article to verify its subject matter.

## Maintenance

The plugin and standalone `SKILL.md` files are identical. Keep them synchronized:

- `plugins/astro-mdx-localization/skills/astro-mdx-localization/SKILL.md`
- `skills/astro-mdx-localization/SKILL.md`

Version the plugin in `.claude-plugin/plugin.json` within this directory and
keep its entry in the root `skills/index.json` at the same version.

## License

Copyright 2026 Rodolfo González González. [MIT](../../LICENSE).
