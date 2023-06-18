# jcb

A terminal-based personal budgeting program that is fast, familiar (heavily inspired by Vim and Mutt) and powerful.

(Still under heavy development but usable)


## Getting Started

Watch this video to get started with a new budget.

[![Getting Started](https://user-images.githubusercontent.com/131466/210484368-1f06f2b2-20b9-49f9-8283-87846da8fbed.png)](https://u.pcloud.link/publink/show?code=XZAWeeVZItk1CMJmI1fBHOlwonuAJmkWr22k)

And this video shows you how to reconcile your budget with your bank statement.

[![Reconciliation](https://user-images.githubusercontent.com/131466/210484542-844bb59d-e49e-4103-bc50-04963660a06e.png)](https://u.pcloud.link/publink/show?code=XZc9eeVZE6jAm2PQMvbxiOQDeGcwmfnRyTzy)


## Features

- Tag and simultaneously operate on many transactions.
- Create repeating transactions.
- Fast navigation with vimlike key bindings.
- Import from TSV.
- Reports and forecasts
- Local data. Private. Nothing is sent to the cloud.

## Building

jcb is written in Go. To build, you need to first install [Go](https://go.dev/doc/install) v1.18+.

Running `go build -o jcb cmd/main.go` will produce a binary for your system at `./jcb`. To install it, run `sudo mv ./jcb /usr/local/bin`.


## Data Formats

### Import/Export

Transactions can be imported from or exported to [tab separated values](https://en.wikipedia.org/wiki/Tab-separated_values) (TSV). The format is:

```
Date	Category	Description	Amount	Notes
```

- The format of `Date` is `YYYY-MM-DD`.
- Category is a single word that is less than 11 characters.
- Description is a string that is less than 33 characters.
- Amount is number that takes the form `<dollars>.<cents>`.
- Notes is a string that is less than 201 characters.

### Savefile

The savefile is a regular [SQLite](https://en.wikipedia.org/wiki/SQLite) database file. You can query or modify it with the `sqlite3` command or anything else that understands SQLite databases.

When the application starts, it copies the database to `.<savefile>.tmp`. Saving will write the data back to `<savefile>`.

The default location of the savefile is `${HOME}/.local/share/jcb/data.db`.


## The UI

The user interface has been inspired quite a bit by the Mutt email editor and the Vim text editor.


### Attributes

Every transaction and budget has a set of attributes which are displayed in the first column of the table.

#### Transaction

Attributes are:

- `C`: Transaction is committed.
- `n`: Transaction has a note.
- `s`: Don't deduct from the budget.
- `+`: Transaction is modified, but not saved.

#### Budget

- `C`: All in this time frame have been committed.
- `n`: The budget has a note.
- `c`: The budget is cumulative. It will retain the remaining credit (or debt) from the prior months budget.
- `+`: Budget is modified, but not saved.


### Info Panel

At the bottom-right of your screen you will find an info panel. It provides an overview of the transactions table. It might look something like this:

```
[13:27] [1] [0]
```

That tells you that:
- The thirteenth transaction is selected.
- There are a total of twenty-seven transactions.
- That one transaction has been modified but not saved.
- That zero transactions are tagged.

