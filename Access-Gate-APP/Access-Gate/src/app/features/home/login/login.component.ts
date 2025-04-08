import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { UsersService } from '../../../services/Users.Service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
  standalone: false
})
export class LoginComponent {
  loginForm: FormGroup;
  loginError: string = '';
  isLoading: boolean = false;

  constructor(
    private fb: FormBuilder,
    private usersService: UsersService,
    private router: Router
  ) {
    this.loginForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]]
    });
  }

  onSubmit(): void {
    if (this.loginForm.valid) {
      this.isLoading = true;
      this.loginError = '';
      
      // Multiple console logs to track data flow
      console.log('🔍 Form Values:', this.loginForm.value);
      console.log('📝 Email:', this.loginForm.get('email')?.value);
      console.log('🔑 Password:', this.loginForm.get('password')?.value);

      this.usersService.login(this.loginForm.value).subscribe({
        next: (response) => {
          console.log('✅ Login Response:', response);
          if (response && response.token) {
            console.log('🎫 Token received:', response.token);
            this.usersService.setToken(response.token);
            this.router.navigate(['/admin']);
          } else {
            this.loginError = 'Invalid response from server';
          }
        },
        error: (error) => {
          if (error.status === 401) {
            this.loginError = 'Invalid email or password';
          } else {
            this.loginError = 'An error occurred during login. Please try again.';
          }
          console.error('Login error:', error);
        },
        complete: () => {
          this.isLoading = false;
        }
      });
    } else {
      this.loginError = 'Please fill in all required fields correctly';
    }
  }
}
