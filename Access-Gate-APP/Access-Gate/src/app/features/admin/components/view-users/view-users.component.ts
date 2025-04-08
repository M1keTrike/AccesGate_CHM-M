import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UsersService } from '../../../../services/Users.Service';
import { User } from '../../models/IUsers';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
    selector: 'app-view-users',
    templateUrl: './view-users.component.html',
    styleUrl: './view-users.component.css',
    standalone: false
})
export class ViewUsersComponent implements OnInit {
    users: User[] = [];
    isLoading = false;

    constructor(
        private usersService: UsersService,
        private router: Router,
        private snackBar: MatSnackBar
    ) {}

    ngOnInit(): void {
        this.loadUsers();
    }

    private getCurrentUserId(): number {
        const token = localStorage.getItem('Authorization');
        if (token) {
            const tokenPayload = JSON.parse(atob(token.split('.')[1]));
            return tokenPayload.user_id;
        }
        return 0;
    }

    loadUsers(): void {
        this.isLoading = true;
        const currentUserId = this.getCurrentUserId();

        this.usersService.getUsersByCreatedBy(currentUserId).subscribe({
            next: (users) => {
                this.users = users;
                this.isLoading = false;
                console.log('Usuarios cargados:', users);
            },
            error: (error) => {
                console.error('Error al cargar usuarios:', error);
                this.showError('Error al cargar usuarios');
                this.isLoading = false;
            }
        });
    }

    editUser(user: User): void {
        this.router.navigate(['/admin/edit-user', user.id]);
    }

    deleteUser(user: User): void {
        if (confirm(`¿Estás seguro de que deseas eliminar al usuario ${user.name}?`)) {
            this.isLoading = true;
            this.usersService.deleteUser(user.id).subscribe({
                next: () => {
                    this.users = this.users.filter(u => u.id !== user.id);
                    this.showSuccess('Usuario eliminado exitosamente');
                    this.isLoading = false;
                },
                error: (error) => {
                    console.error('Error al eliminar usuario:', error);
                    this.showError('Error al eliminar usuario');
                    this.isLoading = false;
                }
            });
        }
    }

    private showSuccess(message: string): void {
        this.snackBar.open(message, 'Cerrar', { 
            duration: 3000,
            panelClass: ['success-snackbar']
        });
    }

    private showError(message: string): void {
        this.snackBar.open(message, 'Cerrar', { 
            duration: 3000,
            panelClass: ['error-snackbar']
        });
    }
}
