import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../../../environments/environment';
import { Event } from '../models/Event';

@Injectable({
  providedIn: 'root'
})
export class EventService {
  private apiUrl = `${environment.apiBaseUrl}/events`;

  constructor(private http: HttpClient) { }

  private getHeaders() {
    const token = localStorage.getItem('Authorization');
    return {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': `${token}` // Removed 'Bearer ' prefix as it should be included in the stored token
      })
    };
  }

  getAllEvents(): Observable<Event[]> {
    return this.http.get<Event[]>(`${this.apiUrl}`, this.getHeaders());
  }

  getEventById(id: number): Observable<Event> {
    return this.http.get<Event>(`${this.apiUrl}/${id}`, this.getHeaders());
  }

  createEvent(event: Omit<Event, 'id' | 'created_at'>): Observable<Event> {
    return this.http.post<Event>(`${this.apiUrl}`, event, this.getHeaders());
  }

  updateEvent(id: number, event: Partial<Event>): Observable<Event> {
    return this.http.put<Event>(`${this.apiUrl}/${id}`, event, this.getHeaders());
  }

  deleteEvent(id: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`, this.getHeaders());
  }

  getEventsByCreator(userId: number): Observable<Event[]> {
    return this.http.get<Event[]>(`${this.apiUrl}/creator/${userId}`, this.getHeaders());
  }
}