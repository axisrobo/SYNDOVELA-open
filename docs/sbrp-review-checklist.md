# SBRP v1 Review Checklist

Status: gate checklist. SBRP v1 may not be frozen until at least one
runtime built independently of AxisRobo has completed this review.

## Why the gate exists

A protocol only its author can implement is not open. The review is not
a courtesy: it is the test of whether an outsider can build a conformant
runtime from the specification alone, using only what is published in
this repository.

## Reviewer instructions

Read `docs/sbrp.md`, the schema at the SBRP v1 contract, and the
conformance checklist, then attempt to implement or simulate each item
below. Do not read any AxisRobo source.

## Runtime role

| # | Item | Reviewer answer |
| --- | --- | --- |
| 1 | Can you construct a `RuntimeDescriptor` from the spec alone? | |
| 2 | Can you serve `describe` without ambiguity about field semantics? | |
| 3 | Can you implement `fetch`/`validate`/`load`/`activate` ordering, or is the sequencing implied but unspecified? | |
| 4 | Is the fail-closed rule ("only ACTIVE accepts invocations") stated precisely enough to test? | |
| 5 | Can you decide what `report` must contain after each operation? | |
| 6 | Is the double-verification requirement implementable (re-verify digest locally)? | |
| 7 | Is there any behaviour you could only learn by reading AxisRobo source? If so, which? | |
| 8 | Are the six mandatory rules each independently testable? | |

## Control plane role

| # | Item | Reviewer answer |
| --- | --- | --- |
| 9 | Can you build a resolver that decides eligibility from a descriptor alone? | |
| 10 | Is the vendor-neutrality rule (never branch on `implementation`) enforceable in practice? | |
| 11 | Can you produce a lock that a runtime would accept, without reading source? | |

## General

| # | Item | Reviewer answer |
| --- | --- | --- |
| 12 | Was the schema sufficient, or did you need examples for disambiguation? | |
| 13 | Is anything overspecified in a way that forecloses a legitimate runtime design? | |
| 14 | Is anything underspecified that forces guessing? | |

## Outcome

| Result | Action |
| --- | --- |
| All items pass with no source reading | Proceed to freeze |
| Source reading required | Fix the spec; the gap is a spec defect, not a reviewer failure |
| Ambiguity without source | Write a clarifying test and add it to the conformance suite |
