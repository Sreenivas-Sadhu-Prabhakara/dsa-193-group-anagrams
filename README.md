# Group Anagrams

> NeetCode 150 · Arrays & Hashing · Blind 75 — the letter-count array IS the map key: Go arrays are comparable, so no sorting per word. Go, stdlib only, tested.

## Problem

Bucket a list of words so each bucket holds exactly the words that are anagrams of one another.

## Approach

Anagrams share a canonical form; the usual one is the sorted word (O(k log k) per word). The sharper Go-specific move: the canonical form can be the 26-letter *count array itself*, because fixed-size arrays — unlike slices — are comparable and therefore legal map keys. Counting is O(k) per word with no allocation beyond the key. A first-seen order list makes the output deterministic (maps iterate randomly in Go — a real trap when tests compare directly). The `aab`/`abb` test pins that multiplicity, not just letter membership, distinguishes groups.

**Complexity:** O(total characters) time, O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
