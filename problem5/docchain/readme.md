# Consensus-Breaking Change

## What is a Consensus-Breaking Change?

In a blockchain network, consensus is the agreement among nodes about the state of the blockchain. All nodes in the network follow a set of rules to validate transactions, create blocks, and update the state of the blockchain. When these rules are changed in a way that is not backwards-compatible, it's called a consensus-breaking change.

A consensus-breaking change can cause nodes in the network to disagree about the state of the blockchain. For example, some nodes might consider a transaction valid according to the new rules, while others might consider the same transaction invalid according to the old rules. This disagreement can lead to a fork in the blockchain, with different nodes maintaining different versions of the blockchain's history.

## Why Does Our Change Break Consensus?

In our case, we have introduced a new field, `category`, to the `document` type in our blockchain application. Previously, a `document` only required a `title` and a `body`. Now, each `document` must also include a `category`.

This change breaks consensus because it alters the rules for validating transactions. Nodes running the old version of the software will reject transactions that include the `category` field because these transactions do not conform to the old rules. On the other hand, nodes running the new version of the software will reject transactions that do not include the `category` field because these transactions do not conform to the new rules.

As a result, the network can split into two: one group of nodes that follows the old rules and another group that follows the new rules. This is why adding the `category` field to the `document` type is a consensus-breaking change.
