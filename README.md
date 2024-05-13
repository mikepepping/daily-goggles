# Daily Goggles
Simple check list CLI tool written in Go for the learning experience.

The goal is to write this tool with a little external depedencies as possible, trying to stick to just the standard lib.

## Commands

Daily Goggles is installed with the alias `dg` to make commands quicker to type and easier to read.

### Insert
To insert a new task
```
dg insert <<name of the task>>
```

And example might be
```
dg insert Add integrations for popular cloud task software
```

### Print
To view all current tasks
```
dg print
```

This will print out all the current tasks.
Example:
```
+---------------------------------------------------------------------------------+
| ID | NAME                   | STATUS | DONE AT                                  |
+---------------------------------------------------------------------------------+
| 0  | Make tables better     | todo   |                                          |
+---------------------------------------------------------------------------------+
| 1  | Implment Print Command | done   | 2024-03-05 20:39:14.642779942 +1000 AEST |
+---------------------------------------------------------------------------------+
```

### Complete
To mark a task as complete, use the complete command with the ID from the ID column from the `print` command

```
dg complete <<ID>>
```

### Clean
Cleaning moves any tasks with the status of `done` into your history.

```
dg clean
```

### History
To view your history of all completed and cleaned tasks

```
dg history
```

This will print a table similar to the `print` command
