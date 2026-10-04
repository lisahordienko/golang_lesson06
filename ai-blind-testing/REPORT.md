# Task 3 — AI blind test generation

> Homework — Task 3 (Lesson 6: File I/O, JSON and Testing).
> This task is **not auto-graded** by CI — it's reviewed by your mentor.
> The workflow only checks that you actually filled this file in
> (see the "Task 3 — AI report present" step in the Actions log).

## Instructions

1. Pick one function you wrote for Task 1 or Task 2 (e.g. `LoadTodos` or
   `ValidateEmail`).
2. Give an AI assistant **only the function signature** — no explanation
   of the logic, no existing test file. Ask it to generate a full
   table-driven test suite for that signature.
3. Run the AI-generated tests against your implementation.
4. Fill in the sections below.

---

## Function under test

`func ValidatePhone(s string) bool` from `validate/validate.go`.

## Prompt you gave the AI

Prompt 1, sent to a separate test-generation agent:

```text
You are doing a blind test-generation exercise. You receive only this function signature and no implementation or test files: `func ValidatePhone(s string) bool`.

Generate a complete table-driven Go test suite candidate. Since the signature does not define a phone-number policy, state a reasonable policy assumption, then provide cases as Go table entries (`name`, `input`, `want`) covering valid inputs, empty input, malformed characters, and length boundaries. Do not inspect or search any repository files. Return the policy, proposed table entries, and any ambiguity that cannot be resolved from the signature alone. This exact request, including only the signature as function-specific input, should be recorded in a homework report.
```

No follow-up prompt was used. The generator was separately instructed not to inspect repository files; the main coding session had already seen the starter tests before this isolated generation run.

## Edge cases the AI found that you had missed

The starter phone table covered only one valid plus-prefixed number, letters,
and the empty string. The generated suite added digit-only numbers, minimum and
maximum lengths, a misplaced or repeated `+`, whitespace, separators, and
non-ASCII digits. I ran the generated cases unchanged against `ValidatePhone`:
13 passed and 3 failed. The failures were `1234567`, `2025550123`, and
`123456789012345`, which the generator expected to accept without a leading
`+`. Its policy assumption was an optional `+` and 7 to 15 digits.

## Edge cases you had that the AI missed

My project-specific cases included a leading zero after `+` (`+012345678`),
the validator's chosen 8-digit minimum (`+12345678`), and an embedded plus
(`+1+2345678`). The generated suite did not cover these because its assumed
format allowed unprefixed digit strings and used a different minimum length.

## Cases where the AI's expected output was wrong

Relative to this implementation's documented policy, the three unprefixed
strings above have incorrect expected values: this validator requires a
leading `+` and applies a locally chosen minimum of 8 digits. That minimum is
not inferred from the signature or claimed as a universal E.164 requirement.
Those strings are not universally invalid phone numbers; the signature alone
cannot establish the intended country, formatting rules, or minimum length.
The mismatch exposed an underspecified contract, not a validator defect under
the chosen policy.

## What you'd change about your own test-writing process after this

I would write down the accepted format and its length limits before adding
examples, then test both sides of every boundary. I would also keep a separate
checklist for punctuation and Unicode, and review AI-generated expectations
against the documented contract instead of treating them as authoritative.
The blind run showed that a function signature alone is not enough to infer a
phone-number policy, so the assumed contract must be reported alongside the
generated tests.
