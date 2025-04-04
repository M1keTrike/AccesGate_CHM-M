import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from '../features/admin/models/IUsers';
import { environment } from '../../environments/environment';

interface LoginResponse {
  token: string;
  user: User;
}

interface LoginCredentials {
  email: string;
  password: string;
}

@Injectable({
  providedIn: 'root'
})
export class UsersService {
  private apiUrl = `${environment.apiBaseUrl}/users`;
  private token: string;

  constructor(private http: HttpClient) {
    try {
      this.token = localStorage.getItem('Authorization') || '';
    } catch (error) {
      console.warn('LocalStorage access error:', error);
      this.token = '';
    }
  }

  private getHeaders(): HttpHeaders {
    return new HttpHeaders({
      'Content-Type': 'application/json',
      'Authorization': `${this.token}`
    });
  }

  login(credentials: LoginCredentials): Observable<LoginResponse> {
    return this.http.post<LoginResponse>(`${this.apiUrl}/login`, credentials, {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    });
  }

  setToken(token: string): void {
    try {
      this.token = token;
      localStorage.setItem('Authorization', token);
    } catch (error) {
      console.warn('LocalStorage access error:', error);
    }
  }

  logout(): void {
    this.token = "";
    localStorage.removeItem('Authorization');
  }

  getUsersByRole(role: string): Observable<User[]> {
    return this.http.get<User[]>(`${this.apiUrl}/role/${role}`, {
      headers: this.getHeaders()
    });
  }

  getUsersByCreatedBy(createdBy: number): Observable<User[]> {
    return this.http.get<User[]>(`${this.apiUrl}/created-by/${createdBy}`, {
      headers: this.getHeaders()
    });
  }

  // Update other methods to use headers as well
  getAllUsers(): Observable<User[]> {
    return this.http.get<User[]>(this.apiUrl, {
      headers: this.getHeaders()
    });
  }

  getUserById(id: number): Observable<User> {
    return this.http.get<User>(`${this.apiUrl}/${id}`, {
      headers: this.getHeaders()
    });
  }

  createUser(data: User): Observable<User> {
    return this.http.post<User>(this.apiUrl, data, {
      headers: this.getHeaders()
    });
  }
  RegisterUser(data: User): Observable<User> {
    return this.http.post<User>(this.apiUrl, data, {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    });
  }

  deleteUser(id: number): Observable<User> {
    return this.http.delete<User>(`${this.apiUrl}/${id}`, {
      headers: this.getHeaders()
    });
  }

  updateUser(id: number, data: User): Observable<User> {
    return this.http.put<User>(`${this.apiUrl}/${id}`, data, {
      headers: this.getHeaders()
    });
  }
}