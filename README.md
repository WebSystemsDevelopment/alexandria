# Alexandria

This project is a Library Management System built using Hexagonal Architecture.

## Tech stack
UI,
    React is a well stablished technology, well documented and easy to use.
    Elm is functional language.

Core,
    The alexandria system is meant to handle big amounts of data. To reduce unexpected bugs we require a static typed language with good performance. At the same time, given that the system will work in a library the system cannot have a big footprint. Languages such as C# or Java have a garbage collector which affects how much space it uses. Rust without garbage collector and it gives more freedom to fine tune the footprint of an app.

Database,
    Given the flow of data where patrons borrow and returns books ther is a lot of ownership. A relational database works well in this kind of environment.

