import { Component, OnInit } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { MatTableModule } from '@angular/material/table';
import { MatIconModule } from '@angular/material/icon';
import { Book } from '../../../core/models/book.model';
import { BookService } from '../../../core/services/book.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-book-list',
  standalone: true,
  imports: [ 
    HttpClientModule,
    MatTableModule,
    MatIconModule
  ],
  templateUrl: './book-list.component.html',
  styleUrl: './book-list.component.css',
  providers: [BookService]
})
export class BookListComponent implements OnInit {
  displayedColumns: string[] = ['id', 'title', 'author', 'isbn', 'actions'];
  books: Book[] = [];

  constructor(private bookService: BookService, private router: Router) {}

  ngOnInit(): void {
    this.bookService.getBooks().subscribe(books => {
      this.books = books;
    });
  }

  onAddBook() {
    this.router.navigate(['/add-book']);
    console.log('Add Book');
  }

  onEditBook(bookId: number) {
    this.router.navigate([`/edit-book/${bookId}`]);
    console.log('Edit Book', bookId);
  }

  onDeleteBook(bookId: number) {
    console.log('Delete Book', bookId);
  }
}
