import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../../../environments/environment';
import { EventAttendee } from '../models/event-attendee';
import { User } from '../models/IUsers';
import { Event } from '../models/event';

@Injectable({
  providedIn: 'root'
})
export class EventAttendeeService {
  private apiUrl = `${environment.apiBaseUrl}/api/event-attendees`;

  constructor(private http: HttpClient) { }

  registerAttendee(attendee: Omit<EventAttendee, 'id' | 'registered_at'>): Observable<EventAttendee> {
    return this.http.post<EventAttendee>(`${this.apiUrl}`, attendee);
  }

  removeAttendee(eventId: number, userId: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/event/${eventId}/user/${userId}`);
  }

  getEventAttendees(eventId: number): Observable<User[]> {
    return this.http.get<User[]>(`${this.apiUrl}/event/${eventId}/attendees`);
  }

  getUserEvents(userId: number): Observable<Event[]> {
    return this.http.get<Event[]>(`${this.apiUrl}/user/${userId}/events`);
  }

  isUserRegistered(eventId: number, userId: number): Observable<boolean> {
    return this.http.get<boolean>(`${this.apiUrl}/check/${eventId}/${userId}`);
  }
}