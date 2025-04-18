import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { UsersService } from '../../../services/Users.Service';
import { User } from '../../admin/models/IUsers';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css'],
  standalone: false
})
export class RegisterComponent implements OnInit {
  registerForm: FormGroup;
  registerError: string = '';

  constructor(
    private fb: FormBuilder,
    private usersService: UsersService,
    private router: Router
  ) {
    this.registerForm = this.fb.group({
      name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      confirmPassword: ['', Validators.required]
    }, { validator: this.passwordMatchValidator });
  }

  ngOnInit(): void {}

  passwordMatchValidator(g: FormGroup) {
    return g.get('password')?.value === g.get('confirmPassword')?.value
      ? null : { mismatch: true };
  }

  onSubmit(): void {
    if (this.registerForm.valid) {
      
      const userData: User = {
        id: 0,
        name: this.registerForm.value.name,
        email: this.registerForm.value.email,
        password_hash: this.registerForm.value.password,
        role: 'admin',
        created_at: new Date().toISOString(), // Ensure this field is correctly formatted
        // Add any other required fields here
      };
      
      console.log('📤 Sending registration data:', userData);
      
      this.usersService.RegisterAdminUser(userData).subscribe({
        next: (response) => {
          console.log('✅ Registration response:', response);
          this.usersService.login({
            email: userData.email,
            password: userData.password_hash
          }).subscribe({
            next: (loginResponse) => {
              this.usersService.setToken(loginResponse.token);
              this.router.navigate(['/admin']);
            },
            error: (error) => {
              this.registerError = 'Registration successful but login failed';
              this.router.navigate(['/login']);
            }
          });
        },
        error: (error) => {
          this.registerError = 'Registration failed. Please try again.';
          console.error('Registration error:', error);
        }
      });
    }
  }
}
