# SOLID Principles Summary

## S — Single Responsibility Principle (SRP)

Break responsibilities into smaller structs/classes.

> One struct = one reason to change.

---

## O — Open/Closed Principle (OCP)

Define interfaces so new behavior is added by implementing them,
not by modifying existing code.

> Open for extension, closed for modification.

---

## L — Liskov Substitution Principle (LSP)

If a subtype cannot fully satisfy an interface contract,
the problem is usually that the interface is too broad.

> The subtype is the victim forced to lie.  
> Split the interface into smaller behavioral contracts.

---

## I — Interface Segregation Principle (ISP)

Consumers (structs) should not depend on methods they do not use (from interfaces).

> The consumer is the victim forced to depend on unnecessary behavior.  
> Split large interfaces into focused ones.

---

## D — Dependency Inversion Principle (DIP)

Depend on abstractions, not concrete implementations.

> Inject interfaces (into high-level modules) from outside instead of instantiating concrete types (low-level modules) internally.
