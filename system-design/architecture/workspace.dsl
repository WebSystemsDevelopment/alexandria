workspace "Alexandria" "Library management system" {

    model {
        alexandria = softwareSystem "Alexandria" "Library management system" {
            ui = container "User interface" "Web application"
            database = container "Database" "Database" {
                db_adapter = component "Adapter" "Allows core to persist data" 
            }

            core = container "Core" "Application" {
                registrar = component "Registrar" "Registers patrons"
                book_management = component "Book Management" "Handles book operations"
                summary = component "Summary" "Performs reports and summaries"

                book_management -> registrar "requests patron data"
                summary -> book_management "requests book data"
                summary -> registrar "requests patron data"

                book_management -> database "persists book data"
                registrar -> database "persists patron data"
            }
        }

        librarian = person "Librarian" "Book manager" {
            this -> alexandria "Lends and returns"
        }

        registrar_office = person "Registrar office" "Manages patron cards" {
            this -> alexandria "Register patrons"
        }

        patron = person "Patron" "Book interested" {
            this -> alexandria "Borrow and return books"
        }
    }

    views {
        styles {
            element "Person" {
                shape Person
            }
        }
    }
}
