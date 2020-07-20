# Secret

An implementation of Shamir's secret sharing.

## What?

Secret sharing splits a secret (typically a password, key or similar short data) into a number of shares to be given to trusted parties, with a threshold of required shares to reassemble lower than the total number.

In other words, a secret can be split into 10 shares, and with a threshold of 2 you can recreate the secret with any pair of shares.

## Recommendations

This implementation technically supports a very large number of shares, with pretty huge secrets (several megabytes). You shouldn't, but you could make use of that. The resulting shares will be larger than the input secret, so it can get clunky.

Keeping the number of shares in the dozens at most is probably smart, and the size of the input secret shouldn't normally be required to be larger than a typical login certificate.
