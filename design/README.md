## Interface and Composition Design

Composition goes beyond the mechanics of type embedding and is more than just a paradigm. It is the key for maintaining stability in your software by having the ability to adapt to the data and transformation changes that are coming.

## Notes

* This is much more than the mechanics of type embedding.
* Declare types and implement workflows with composition in mind.
* Understand the problem you are trying to solve first. This means understanding the data.
* The goal is to reduce and minimize cascading changes across your software.
* Interfaces provide the highest form of composition.
* Don't group types by a common DNA but by a common behavior.
* Everyone can work together when we focus when we do and not what we are.

## Quotes

_"Methods are valid when it is practical or reasonable for a piece of data to expose a capability." - William Kennedy_

## Design Guidelines

**Design Philosophy:**

* Interfaces give programs structure.
* Interfaces encourage design by composition.
* Interfaces enable and enforce clean divisions between components.
  * The standardization of interfaces can set clear and consistent expectations.
* Decoupling means reducing the dependencies between components and the types they use.
  * This leads to correctness, quality and performance.
* Interfaces allow you to group concrete types by what they do.
  * Don't group types by a common DNA but by a common behavior.
  * Everyone can work together when we focus on what we do and not who we are.
* Interfaces help your code decouple itself from change.
  * You must do your best to understand what could change and use interfaces to decouple.
  * Interfaces with more than one method have more than one reason to change.
  * Uncertainty about change is not a license to guess but a directive to STOP and learn more.
* You must distinguish between code that:
  * defends against fraud vs protects against accidents

**Validation:**

Declare an interface when:
* users of the API need to provide an implementation detail.
* API’s have multiple implementations they need to maintain internally.
* parts of the API that can change have been identified and require decoupling.

Don't declare an interface:
* for the sake of using an interface.
* to generalize an algorithm.
* when users can declare their own interfaces.
* if it's not clear how the ineterface makes the code better.

### Exercise 1

Using the template, declare a set of concrete types that implement the set of predefined interface types. Then create values of these types and use them to complete a set of predefined tasks.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/uY6KMprfMR)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/nbd3gnLlih))

Solution:
```go
package main

// Add import(s).
import "fmt"

// administrator represents a person or other entity capable of administering
// hardware and software infrastructure.
type administrator interface {
	administrate(system string)
}

// developer represents a person or other entity capable of writing software.
type developer interface {
	develop(system string)
}

// =============================================================================

// adminlist represents a group of administrators.
type adminlist struct {
	list []administrator
}

// Enqueue adds an administrator to the adminlist.
func (l *adminlist) Enqueue(a administrator) {
	l.list = append(l.list, a)
}

// Dequeue removes an administrator from the adminlist.
func (l *adminlist) Dequeue() administrator {
	a := l.list[0]
	l.list = l.list[1:]
	return a
}

// =============================================================================

// devlist represents a group of developers.
type devlist struct {
	list []developer
}

// Enqueue adds a developer to the devlist.
func (l *devlist) Enqueue(d developer) {
	l.list = append(l.list, d)
}

// Dequeue removes a developer from the devlist.
func (l *devlist) Dequeue() developer {
	d := l.list[0]
	l.list = l.list[1:]
	return d
}

// =============================================================================

// Declare a concrete type named sysadmin with a name field of type string.
type sysadmin struct {
	name string
}

// Declare a method named administrate for the sysadmin type, implementing the
// administrator interface. administrate should print out the name of the
// sysadmin, as well as the system they are administering.
func (s *sysadmin) administrate(system string) {
	fmt.Println(s.name, "is administering", system)
}

// Declare a concrete type named programmer with a name field of type string.
type programmer struct {
	name string
}

// Declare a method named develop for the programmer type, implementing the
// developer interface. develop should print out the name of the
// programmer, as well as the system they are coding.
func (p *programmer) develop(system string) {
	fmt.Println(p.name, "is developing", system)
}

// Declare a concrete type named company. Declare it as the composition of
// the administrator and developer interface types.
type company struct {
	administrator
	developer
}

// =============================================================================

func main() {

	// Create a variable named admins of type adminlist.
	var admins adminlist

	// Create a variable named devs of type devlist.
	var devs devlist

	// Enqueue a new sysadmin onto admins.
	admins.Enqueue(&sysadmin{"John"})

	// Enqueue two new programmers onto devs.
	devs.Enqueue(&programmer{"Mary"})
	devs.Enqueue(&programmer{"Steve"})

	// Create a variable named cmp of type company, and initialize it by
	// hiring (dequeuing) an administrator from admins and a developer from devs.
	cmp := company{
		administrator: admins.Dequeue(),
		developer:     devs.Dequeue(),
	}

	// Enqueue the company value on both lists since the company implements
	// each interface.
	admins.Enqueue(&cmp)
	devs.Enqueue(&cmp)

	// A set of tasks for administrators and developers to perform.
	tasks := []struct {
		needsAdmin bool
		system     string
	}{
		{needsAdmin: false, system: "xenia"},
		{needsAdmin: true, system: "pillar"},
		{needsAdmin: false, system: "omega"},
	}

	// Iterate over tasks.
	for _, task := range tasks {

		// Check if the task needs an administrator else use a developer.
		if task.needsAdmin {

			// Dequeue an administrator value from the admins list and
			// call the administrate method.
			adm := admins.Dequeue()
			adm.administrate(task.system)

			continue
		}

		// Dequeue a developer value from the devs list and
		// call the develop method.
		dev := devs.Dequeue()
		dev.develop(task.system)
	}
}
```
