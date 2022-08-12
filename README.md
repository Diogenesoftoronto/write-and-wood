# Write-and-wood

a creative writing game that challenges writers to answer prompts.


## About

When I was in highschool, I had a creative writing class where the teacher had us answer a prompt given with a given theme within the first fifteen minutes of the class. We then shared what some of us had written.

This application takes such an excersise and makes it into a game. Taking that to the logical extreme, it creates a series of free form challenges with a wide variety of themes and prompts that challenges writers to improve their abilities and skills. 

Share your creations with the community, extend others' responses, make your own prompts, and participate in writing events and competitions.

## Setup 

TBD

## CLI

write and wood was designed first and for most as a command line tool. Which means that all of its functions are *perfectely* accessible from a terminal. The reason behind this is that designs for terminal applications rarely change while GUIs almost always do. So to get this right we finished the CLI before we finished the GUI. 
This way we kept developer experience first and allow you to easily access and use our api!

### Commands

The CLI has several commands but has all of the basics. Modeled after other great tools like Git and cURL. We intend for ours to be much simpler and easier to use if that is possible. 

We call the write and wood CLI 'wraw', pronounced 'RAW'. All commands must use this prefix.

1. HELP

The first and best command. Everyone needs a helping. Help will give you information about what a command does. 

Technically help is not a command but a flag. We believe that having help as a flag falls in line with standard practice and is more intuitive. Which is why help can be accessed passing ``-h`` or ``--help`` into the CLI. Here is an example: 

```sh
wraw --help
```

You have probably seen this in Command line tools before if you have used them. ***Help and all flags must be suffixed otherwise the command will fail!***


