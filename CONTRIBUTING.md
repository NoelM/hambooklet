# Welcome to Amateur Radio Booklet's the contributing guide

The booklet is open to contribution and protected by the CC BY-SA 4.0 license.
You can propose contributions through Pull Requests.

The purpose of the booklet is to be used globally, when you are operating, in the hurry, outside (SOTA, portable, ...), or off-line.

## Requirements

* The editing software [LibreOffice](https://www.libreoffice.org/)
* The [fonts](#fonts)

## Content

If you want to add new content to the booklet, you must validate those rules:

* The content is related to the Amateur Radio regulations or technics
* The content is available globally, all the IARU regions
* The content **MUST NOT** promote any commercial company or commercial project
* The content **MUST** cite its sources
* The content **MUST** be objective

The content **MUST NOT CONTAIN** remarks or judgment about operators depending on their origins, region, or favorite mode.
Some **DO NOT** examples:
> Operators from Somewhere are not able to set properly their emitters
>
> Operators using the mode XYZ are not skilled

## Layouts

### Page

The booklet is intended to be printed, so the content has to fit within the A6 format[^1]. The page margins are the following:

| Inner | Top | Right | Outer |
|-------|-----|-------|---------|
| 1 cm  | 0.5 cm | 0.5 cm | 0.5 cm |

The page layout is mirrored, so:

* Odd Pages, Inner is Left
* Even Pages, Inner is Right

### Fonts

The document uses the following fonts:

* Headings, Tables, Body: [Inter](https://rsms.me/inter/) for readability, accessibility, and internationalization
* Monospace: [Liberation Mono](https://github.com/liberationfonts)

The fonts are available on every operating system without fees.

### Headings

The heading styles are the following:

| Style | Font | Type | Size | Notes |
| --- | --- | --- | --- | --- |
| Title | Inter | Bold | 28 pt | Once on the cover |
| Heading 1 | Inter | Bold | 18 pt | Always preceded by a page break |
| Heading 2 | Inter | Bold | 12 pt | Preceded by a page break if less than 4 text-body lines ahead of the page's bottom |
| Heading 3 | Inter | Bold | 10 pt | Same as previous |

### Body

A _Text Body_ style is defined in the document (Inter Regular, 8 pt). It has to be used everywhere.

### Tables

#### General

* Width, 100% of the page or 9 cm
* Style _Table Heading_ for the head of the table
* Style _Table Vertical_ for vertically oriented content
* Style _Table Content_ for the rest of the table

The table rows must not break across pages and columns.

#### Headers and Sub-Headers

The header uses the _Table Heading_ style and the background _Light Gray 5_.
The header must be repeated among each new page if:

* the table as at least 3 columns,
* or the table is split on 3 pages or more.

Sub-headers could be used, if the table is split across pages and has numerous entries (see Prefix Allocation).
A _Table SubHeading_ style is defined, it must be used with a background _Light Gray 5_.

#### Borders

The table could use border or not, depending its usage. If the table is intended to organize content
without a header (see Morse Code), the table does not require borders.

However, if the table is intended to organize content with a specific header, borders must be defined:

* Every border is active, except is cells are merged
* Border Width: 0.5 pt
* Border Color: Black
* Padding: 0.1 cm everywhere

### Citations

If you reuse content with or without modification which is not yours, you must refer to them author. A style _Quotations_ is defined for long paragraph citations (see Prefix Allocation).

If you add a new cited content, please update the **Reference** chapter.
_Not defined yet: Rules to define reference markers in the document._

### Abbreviations

The booklet is intended to have abbreviations defined, if you create another one, please update the table at the end of the document and sort it alphabetically.

[^1]: [ISO 216 paper formats](https://fr.wikipedia.org/wiki/ISO_216)
