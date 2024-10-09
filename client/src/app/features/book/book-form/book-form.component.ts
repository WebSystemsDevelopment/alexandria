import { Component } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { ReactiveFormsModule } from '@angular/forms';
import { BookService } from '../../../core/services/book.service';

@Component({
  selector: 'app-book-form',
  standalone: true,
  imports: [
    HttpClientModule,
    CommonModule, 
    MatFormFieldModule, 
    MatInputModule, 
    MatButtonModule, 
    ReactiveFormsModule
  ],
  templateUrl: './book-form.component.html',
  styleUrl: './book-form.component.css',
  providers: [BookService]
})
export class BookFormComponent {
  bookForm: FormGroup;

  constructor(private fb: FormBuilder, private bookService: BookService) {
    this.bookForm = this.fb.group({
      title: ['', [Validators.required, Validators.minLength(3)]],
      author: ['', Validators.required],
      isbn: ['', [Validators.required, Validators.pattern(/^[0-9]{10,13}$/)]],
    });
  }

  onSubmit(): void {
    if (this.bookForm.valid) {
      console.log('Book data:', this.bookForm.value);
      this.bookService.addBook(this.bookForm.value).subscribe(response => {
        console.log('Book added successfully');
      });
    } else {
      console.log('Form is invalid');
    }
  }

  onCancel(): void {
    this.bookForm.reset();
  }
}
