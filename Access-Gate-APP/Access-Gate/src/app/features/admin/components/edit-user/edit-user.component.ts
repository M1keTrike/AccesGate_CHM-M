import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';
import { UsersService } from '../../../../services/Users.Service';
import { User } from '../../models/IUsers';

@Component({
    selector: 'app-edit-user',
    templateUrl: './edit-user.component.html',
    styleUrl: './edit-user.component.css',
    standalone: false
})
export class EditUserComponent implements OnInit {
    userForm: FormGroup;
    isLoading = false;
    userId: number;
    originalCreatedBy: number = 0 ; // Store original created_by

    constructor(
        private fb: FormBuilder,
        private usersService: UsersService,
        private router: Router,
        private route: ActivatedRoute,
        private snackBar: MatSnackBar
    ) {
        this.userForm = this.fb.group({
            name: ['', [Validators.required, Validators.minLength(3)]],
            email: ['', [Validators.required, Validators.email]],
            password: ['', Validators.required],
            role: ['', Validators.required]
        });
        this.userId = 0;
        this.originalCreatedBy = 0; // Initialize originalCreatedBy
    }

    ngOnInit(): void {
        this.route.params.subscribe(params => {
            if (params['id']) {
                this.userId = +params['id'];
                this.loadUser();
            }
        });
    }

    private loadUser(): void {
        this.isLoading = true;
        this.usersService.getUserById(this.userId).subscribe({
            next: (user) => {
                this.userForm.patchValue({
                    name: user.name,
                    email: user.email,
                    password: user.password_hash,
                    role: user.role
                });
                this.originalCreatedBy = user.created_by ?? 0; // Use default value if undefined
                this.isLoading = false;
            },
            error: (error) => {
                console.error('Error al cargar usuario:', error);
                this.showError('Error al cargar usuario');
                this.isLoading = false;
                this.router.navigate(['/admin/view-users']);
            }
        });
    }

    onSubmit(): void {
        if (this.userForm.valid) {
            this.isLoading = true;
            const userData: User = {
                id: this.userId,
                name: this.userForm.value.name,
                email: this.userForm.value.email,
                password_hash: this.userForm.value.password,
                role: this.userForm.value.role,
                created_by: this.originalCreatedBy, // Use original created_by
                created_at: new Date().toISOString()
            };

            console.log('Submitting user data:', userData);

            this.usersService.updateUser(userData.id, userData).subscribe({
                next: () => {
                    this.showSuccess('Usuario actualizado exitosamente');
                    this.router.navigate(['/admin/view-users']);
                },
                error: (error) => {
                    console.error('Error al actualizar usuario:', error);
                    this.showError('Error al actualizar usuario');
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