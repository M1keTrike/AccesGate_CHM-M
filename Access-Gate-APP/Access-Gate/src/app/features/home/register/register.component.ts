import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { UsersService } from '../../../services/Users.Service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss'],
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
      fullName: ['', Validators.required],
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
      const userData = {
        ...this.registerForm.value,
        role: 'user',
        username: this.registerForm.value.email // Using email as username
      };
      delete userData.confirmPassword;

      this.usersService.createUser(userData).subscribe({
        next: (response) => {
          // After successful registration, log in the user
          this.usersService.login({
            username: userData.username,
            password: userData.password
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
