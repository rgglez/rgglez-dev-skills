---
name: astro-mdx-localization
description: Translate or update localized Markdown and MDX content in Astro projects while preserving content identity, executable examples, components, assets, and locale navigation. Use when a source article and target locale are provided or inferable from the request.
license: MIT
---

# Astro MDX localization

Translate the requested content using the project's own localization model.
Support any source and target languages; do not assume a particular theme,
directory layout, package manager, or hosting provider.

## Discover the project conventions

Before writing the translation, inspect:

- The source file, any existing target translation, and a nearby localized
  article. An existing target may contain edits worth preserving.
- Applicable repository instructions and the content collection schema or
  loader. Locate these within the actual Astro application, which may be in
  a monorepo subdirectory.
- The routing and translation lookup used by this content. Determine whether
  translations share a filename, an explicit identifier, a slug, or another key.
  Do not assume filenames and slugs serve the same purpose.
- The application package scripts, package manager, formatter configuration,
  and Markdown/MDX plugins, including generated table-of-contents behavior.

Use this evidence to determine the destination, localized URL, metadata fields,
and validation commands. Ask only if the target language or a consequential
convention remains ambiguous. Create only the requested locales.

## Preserve identity and meaning

- Preserve the article's scope, structure, technical claims, qualifications,
  numbers, and versions. Translate naturally without adding new material.
  If the source appears incorrect, report the issue separately rather than
  silently revising it as part of translation.
- Translate reader-facing titles, descriptions, headings, prose, table cells,
  image alt text, captions, and component labels.
- Set the locale and localize the slug only as required by the project's
  schema and routing. Keep the established translation key intact. A shared
  filename is mandatory only when the project uses it to link translations.
- Preserve authorship, dates, publication flags, and machine-readable metadata
  unless the request or project conventions require changes. Translate
  descriptive tags only when the tag taxonomy supports it. Do not invent fields
  or mechanically translate identifiers, license names, or enum values.
- Reuse existing assets. Preserve the resolved asset targets; adjust relative
  paths if the destination depth changes. Keep working aliases and imports.
  Generate or replace images only when requested.

## Keep examples and MDX functional

- Preserve executable commands, options, filenames, identifiers, environment
  variables, API names, hashes, and technical output. Translate explanatory
  code comments when appropriate.
- Translate user-facing strings in a provided script only when they are not
  protocol values or otherwise consumed by code. Keep matching sample output
  and troubleshooting text consistent. Preserve verbatim diagnostic messages
  emitted by external tools. Do not run article commands merely to translate
  them: they may install software, modify data, or contact services.
- Preserve JSX/component names, imports, expressions, and functional props.
  Translate textual props according to their meaning, not all strings blindly.
  Maintain whitespace needed for Markdown inside components to render.
- Keep source and external links unless an authoritative localized equivalent
  is known. For internal links, verify that the target locale page actually
  exists before changing its URL. Follow the site's fallback behavior when it
  does not; do not invent routes.
- Follow the project's HTML-versus-Markdown link convention. Avoid angle-bracket
  URL autolinks that the MDX parser treats as JSX. Check that formatting has not
  split an inline link into a separate paragraph.
- If the table of contents is generated, use the heading recognized for the
  target locale; do not also write a manual list. Check localized anchor links
  against generated heading IDs. If a new locale needs routing or plugin
  configuration, identify that requirement and keep changes within the task's
  authorized scope.

## Validate and report

1. Review source and target side by side for omitted paragraphs, changed
   meaning, broken components, missing links, and unintended code changes.
   Compare code blocks with allowances for deliberately translated comments
   and messages; byte-for-byte equality is not sufficient on its own.
2. Format only the changed files using the project's configured formatter from
   the application directory. Review its changes, especially JSX and inline
   links. In a Git checkout, check whitespace and the final diff, including
   newly created files, without changing unrelated work.
3. Run the existing build or relevant content validation using the declared
   package manager. Do not assume a fixed executable path or install a new
   toolchain merely because a command is missing. Distinguish content errors
   from environment restrictions; request permitted access when needed instead
   of changing valid source to accommodate the sandbox.
4. Inspect the generated page or available preview for translated metadata,
   heading/TOC rendering, asset resolution, and links between locales where
   the site provides them. Do not assume a fixed build output directory or that
   all routes produce static HTML. Run a search-index check only if the project
   includes one in its workflow.

Report the destination, localized URL when verified, completed checks, and any
remaining validation limits. Do not describe retained verification dates in the
article as tests performed during the translation task.
