import { Routes } from '@angular/router';
import { BookListComponent } from './features/book/book-list/book-list.component';
import { BookFormComponent } from './features/book/book-form/book-form.component';

export const routes: Routes = [
    { path: 'books', component: BookListComponent },
    { path: 'add-book', component: BookFormComponent },
    { path: '', redirectTo: '/books', pathMatch: 'full' }
];
