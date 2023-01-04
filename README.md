# jcb

A TUI personal budgeting program that is fast, familiar (heavily inspired by Vim and Mutt) and powerful.

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
- Local data. Private. Nothing is sent to the cloud.


## The UI

The user interface has been inspired quite a bit by the Mutt email editor and the Vim text editor.


### Transaction Attributes

Every transaction has a set of three attributes which are displayed in the first column of the transactions table.

Attributes are:

- `C`: Transaction is committed.
- `n`: Transaction has a note.
- `+`: Transaction is modified, but not saved.


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

