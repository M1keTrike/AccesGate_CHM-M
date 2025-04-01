import { Injectable } from '@angular/core';
import { HttpClient,HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../../../environments/environment';
import { Event } from '../models/event';

@Injectable({
    providedIn: 'root'
})
export class EventService {
    private apiUrl = `${environment.apiBaseUrl}/events`;

    constructor(private http: HttpClient) { }
    private getHeaders(): HttpHeaders {
        return new HttpHeaders({
          'Content-Type': 'application/json',
          'authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzEyMDI4NjY3fQ.Kp6uAqcY0wvFBsWAeGZRFtKqQxSXIQUQAwkBsLgUxU0'
        });
      }

    createEvent(event: Omit<Event, 'id' | 'created_at'>): Observable<Event> {
        return this.http.post<Event>(this.apiUrl, event, {
            headers: this.getHeaders()
          });
        }
    

    getEventById(id: number): Observable<Event> {
        return this.http.get<Event>(`${this.apiUrl}/${id}`, {
            headers: this.getHeaders()
          });
        }

    getAllEvents(): Observable<Event[]> {
        return this.http.get<Event[]>(this.apiUrl, {
            headers: this.getHeaders()
          });
        }

    updateEvent(event: Event): Observable<Event> {
        return this.http.put<Event>(`${this.apiUrl}/${event.id}`, event, {
            headers: this.getHeaders()
          });
        }

    deleteEvent(id: number): Observable<void> {
        return this.http.delete<void>(`${this.apiUrl}/${id}`, {
            headers: this.getHeaders()
          });
        }
    getEventsByCreator(userId: number): Observable<Event[]> {
        return this.http.get<Event[]>(`${this.apiUrl}/creator/${userId}`, {
            headers: this.getHeaders()
          });
        }
}