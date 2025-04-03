import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, forkJoin, of } from 'rxjs';
import { mergeMap } from 'rxjs/operators';
import { environment } from '../../../../environments/environment';
import { EventAttendee } from '../models/event-attendee';
import { User } from '../models/IUsers';
import { Event } from '../models/event';
import { UsersService } from '../../../services/Users.Service';

@Injectable({
  providedIn: 'root'
})
export class EventAttendeeService {
  private apiUrl = `${environment.apiBaseUrl}/event-attendees`;

  constructor(private http: HttpClient, private userService: UsersService) { }

  registerAttendee(attendee: Omit<EventAttendee, 'id' | 'registered_at'>): Observable<EventAttendee> {
    return this.http.post<EventAttendee>(`${this.apiUrl}/register`, attendee);
  }

  removeAttendee(eventId: number, userId: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/events/${eventId}/users/${userId}`);
  }

  getEventAttendees(eventId: number): Observable<User[]> {
    return this.http.get<EventAttendee[]>(`${this.apiUrl}/events/${eventId}/attendees`).pipe(
      mergeMap((attendees: EventAttendee[]) => {
        if (attendees.length === 0) {
          return of([]);
        }
        return forkJoin(
          attendees.map((attendee: EventAttendee) => 
            this.userService.getUserById(attendee.user_id)
          )
        );
      })
    );
  }

  getUserEvents(userId: number): Observable<Event[]> {
    return this.http.get<Event[]>(`${this.apiUrl}/users/${userId}/events`);
  }

  isUserRegistered(eventId: number, userId: number): Observable<boolean> {
    return this.http.get<boolean>(`${this.apiUrl}/events/${eventId}/users/${userId}/check`);
  }
}