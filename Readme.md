### Battle Simulator

Code that simulates a battle between 2 combatants that can have various abilities.

##### Usage

> go run main.go

#### Tests

For running tests, run:

> go test ./...

**Note about Code Coverage**

The code coverage for the project is around `~67%`. This is due to the fact that I didn't test at all the component (LogCommentator / log_commentator.go) that reports what's happening during a duel. That component is responsible with "presentation logic" of the duel and contains lots of calls to `log.Print`.
