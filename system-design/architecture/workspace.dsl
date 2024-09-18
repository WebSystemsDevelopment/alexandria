workspace "Alexandria" "Library management system" {

    model {
        alexandria = softwareSystem "Alexandria" "Library management system" {

            database = container "Database" "Database" {
                db_adapter = component "Adapter" "Allows core to persist data" 
            }

            ui = container "User interface" "Web application" {
                rest_adapter = component "REST adapter" "Allows external web applications access"
            }

            core = container "Core" "Application" {
                registrar = component "Registrar" "Registers patrons"
                book_management = component "Book Management" "Handles book operations"
                summary = component "Summary" "Performs reports and summaries"

                db_port = component "Database port" "Port for database connection"
                rest_port = component "REST port" "Port for restful connection"

                book_management -> registrar "requests patron data"
                summary -> book_management "requests book data"
                summary -> registrar "requests patron data"

                book_management -> db_port "persists book data"
                registrar -> db_port "persists patron data"

                db_port -> db_adapter
                rest_port -> rest_adapter
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

        admin = person "Admin" "Adminsterst the library" {
            this -> alexandria "Wants reports and summaries"
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
