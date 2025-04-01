import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from '../models/IUsers';
import { environment } from '../../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class UsersService {
  private apiUrl = `${environment.apiBaseUrl}/users`;

  constructor(private http: HttpClient) {}

  private getHeaders(): HttpHeaders {
    return new HttpHeaders({
      'Content-Type': 'application/json',
      'Authorization': 'Bearer token'
    });
  }

  getUsersByRole(role: string): Observable<User[]> {
    return this.http.get<User[]>(`${this.apiUrl}/role/${role}`, {
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

  deleteUser(id: number): Observable<User> {
    return this.http.delete<User>(`${this.apiUrl}/${id}`, {
      headers: this.getHeaders()
    });
  }
}