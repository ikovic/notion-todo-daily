# notion-todo-daily

Notion API consumer that manages daily tasks

## Installation

Write notes about how to install all the modules after cloning.

## To do
Now that we can fetch the blocks containing `to_do` objects, we should develop an algorithm that does the following:
- for each top level `to_do`:
    - check if it is checked
    - if yes, ignore it and continue
    - if not, copy it and set it as the root object
    - apply the algorithm the same way
Expected output is an object tree that would only consist of unchecked `to_do` objects.
## Learning notes

- What is the advantage of `context.Context` versus just importing packages that we need?
    - `context.Context` internally is using pointers to make sure the values are not copied but instead referenced
    - in some cases it would be easier to just import packages and let them configure themselves in their `init` functions

