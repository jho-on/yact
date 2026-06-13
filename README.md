# Yet Another CLI Todo (YACT)

yact is a simple, hopefully fast and lightweight CLI application for managing todo lists and tasks.

yact is fully local, and all data is stored in a JSON file.

## Commands Structure

All commands follow a similar structure:

    yact <command> <params> <flags>
    Ex: yact add liro -d laro -e 3d -p 2 -list "Cool List 😎"

### Components

- command

  A short word that represents an operation you can perform. In most cases, it is followed by parameters. In the example, we pass one parameter, the task title.

- params
  Required information that the command needs to operate. In most cases, it will be just one parameter. They must not be empty.

- flags
  Flags appear as key-value pairs. They can be written in short form (`-d`, `-l`) or long form (`--description`, `--list`). All flags must be followed by a valid value.
