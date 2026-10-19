# Go Refactoring and Design Patterns

Code from the [I'll Be Backend](https://www.youtube.com/@illbebackend) videos: messy Go code gets better, one small, safe refactoring at a time. Every lesson has a video, an [interactive version](https://illbebackend.dev) you can step through, and a folder here with the code after every step.

## Code smells

| Lesson | Steps | Video | Interactive |
|---|---|---|---|
| [Long parameter list: introduce a parameter object](long-parameter-list-introduce-a-parameter-object/) | 12 | [Watch](https://www.youtube.com/watch?v=TZuVEMBOB8k) | [Open](https://illbebackend.dev/lessons/long-parameter-list-introduce-a-parameter-object/go/) |
| [Clean code: replace comments with code](replace-comments-with-code/) | 14 | [Watch](https://www.youtube.com/watch?v=Pq_E5pg411Q) | [Open](https://illbebackend.dev/lessons/replace-comments-with-code/go/) |

## Principles

| Lesson | Steps | Video | Interactive |
|---|---|---|---|
| [DRY: one rule, one place](dry/) | 11 | [Watch](https://www.youtube.com/watch?v=JIynHedsM14) | [Open](https://illbebackend.dev/lessons/dry/go/) |
| [KISS: replace a rule engine with plain code](kiss-unwind-a-clever-abstraction/) | 13 | [Watch](https://www.youtube.com/watch?v=BLqS4lM9p_0) | [Open](https://illbebackend.dev/lessons/kiss-unwind-a-clever-abstraction/go/) |
| [Clean code: rename for intent](rename-for-intent/) | 13 | [Watch](https://www.youtube.com/watch?v=SbTa_iTQl_U) | [Open](https://illbebackend.dev/lessons/rename-for-intent/go/) |
| [YAGNI: delete speculative options](yagni-delete-speculative-options/) | 13 | [Watch](https://www.youtube.com/watch?v=3HBTyq8YpRA) | [Open](https://illbebackend.dev/lessons/yagni-delete-speculative-options/go/) |

## How to use it

- Open a lesson folder and read its README, then the step folders in order.
- `sh check.sh` compiles every step in Docker with `golang:1.23`, the same check every lesson passed before its video went out.
- Found a bug in a lesson? Open an issue.

## License

The code is [MIT No Attribution](LICENSE): use it however you like. The videos, the name and the logo are covered by the [terms](https://illbebackend.dev/terms/).
